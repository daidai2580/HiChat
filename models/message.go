package models

import (
	"HiChat/global"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gopkg.in/fatih/set.v0"
	"net"
	"net/http"
	"strconv"
	"sync"
)

type Message struct {
	Model
	FromId   int64  `json:"userId"`
	TargetId int64  `json:"targetId"`
	Type     int    //聊天类型：群聊 私聊 广播
	Media    int    //信息类型：文字 图片 音频
	Content  string //内容
	Pic      string `json:"url"` //图片url
	Url      string //文件相关
	Desc     string
	Amount   int //其它数据
}

func (m *Message) MessageTable() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	Addr      string
	DataQueue chan []byte
	GroupSets set.Interface
}

var clientMap = make(map[int64]*Node, 0)
var upSendChan chan []byte = make(chan []byte, 1204)
var rwLock sync.RWMutex

// Chat    需要 ：发送者ID ，接受者ID ，消息类型，发送的内容，发送类型
func Chat(w http.ResponseWriter, r *http.Request) {
	//获取参数信息发送者ID
	query := r.URL.Query()
	id := query.Get("userId")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		zap.S().Info("类型转换失败", err)
		return
	}
	//升级为socket
	var isvalida = true
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	rwLock.Lock()
	clientMap[userId] = node
	rwLock.Unlock()

	go sendProc(node)

	go recProc(node)
}

// 从node中获取信息并写入websocket中
func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				zap.S().Info("写入消息失败", err)
				return
			}
			fmt.Println("数据发送socket成功")
		}

	}
}

// 从websocket中奖消息体拿出，然后近些解析，在进行消息类型判断，最后奖消息发送值目的用户的node种
func recProc(node *Node) {
	for {
		//获取信息
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			zap.S().Info("读取消息失败", err)
			return
		}

		//这里是简单实现的一种方法
		msg := Message{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			zap.S().Info("json解析失败", err)
			return
		}
		brodMsg(data)
	}
}

func brodMsg(msg []byte) {
	upSendChan <- msg
}

func sendMsg(id int64, msg []byte) {
	rwLock.Lock()
	node, ok := clientMap[id]
	rwLock.Unlock()

	if !ok {
		zap.S().Info("UserID没有对应的node")
		return
	}
	zap.S().Info("targetId:", id, "node:", node)
	if ok {
		node.DataQueue <- msg
	}
}

func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		zap.S().Info("消息解析失败：", err)
		return
	}
	fmt.Println("解析数据:", msg, "msg.FormId", msg.FromId, "targetId:", msg.TargetId, "type:", msg.Type)
	//判断消息类型
	switch msg.Type {
	case 1:
		sendMsgAndSave(msg.TargetId, data)
	case 2:
		sendGroup(uint(msg.FromId), uint(msg.TargetId), data)

	}
}

func sendGroup(fromId uint, targetId uint, data []byte) (int, error) {
	users, err := FindUsers(targetId)
	if err != nil {
		return -1, err
	}
	for _, userId := range *users {
		if fromId != userId {
			sendMsgAndSave(int64(userId), data)
		}
	}
	return 0, nil
}

func sendMsgAndSave(userId int64, msg []byte) {
	rwLock.RLock()
	node, ok := clientMap[userId]
	rwLock.RUnlock()
	jsonMsg := Message{}
	json.Unmarshal(msg, &jsonMsg)
	ctx := context.Background()
	targetIdstr := strconv.Itoa(int(userId))
	userIdStr := strconv.Itoa(int(jsonMsg.FromId))

	if ok {
		//如果当前用户在线，将洗洗转发到用户的websocker链接中然后进行存储
		node.DataQueue <- msg
	}
	var key string
	if userId > jsonMsg.FromId {
		key = "msg_" + userIdStr + "_" + targetIdstr
	} else {
		key = "msg_" + targetIdstr + "_" + userIdStr
	}
	res, err := global.RedisDB.ZRevRange(ctx, key, 0, 1).Result()

	if err != nil {
		fmt.Println(err)
		return
	}
	score := float64(cap(res)) + 1
	result, err := global.RedisDB.ZAdd(ctx, key, &redis.Z{score, msg}).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func UpdRecProc() {
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})
	if err != nil {
		zap.S().Info("监听udp端口时报", err)
		return
	}
	defer udpConn.Close()
	for {
		var buf [1024]byte
		n, err := udpConn.Read(buf[0:])
		if err != nil {
			zap.S().Info("读取Udp数据失败", err)
			return
		}
		dispatch(buf[0:n])
	}
}

func UdpSendProc() {
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
		Zone: "",
	})
	if err != nil {
		zap.S().Info("拨号UDP端口失败", err)
		return
	}
	defer udpConn.Close()
	for {
		select {
		case data := <-upSendChan:
			_, err := udpConn.Write(data)
			if err != nil {
				zap.S().Info("写入udp消息失败", err)
				return
			}
		}
	}
}

func RedisMsg(userIdA int64, userIdB int64, start int64, end int64, isRev bool) []string {
	ctx := context.Background()
	userIdStr := strconv.Itoa(int(userIdA))
	targetIdStr := strconv.Itoa(int(userIdB))

	var key string
	if userIdA > userIdB {
		key = "msg_" + targetIdStr + "_" + userIdStr
	} else {
		key = "msg_" + userIdStr + "_" + targetIdStr
	}
	var rels []string
	var err error

	if isRev {
		rels, err = global.RedisDB.ZRange(ctx, key, start, end).Result()

	} else {
		rels, err = global.RedisDB.ZRevRange(ctx, key, start, end).Result()
	}
	if err != nil {
		fmt.Println(err)
	}
	return rels

}

func init() {
	go UdpSendProc()
	go UpdRecProc()
}
