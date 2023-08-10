package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gopkg.in/fatih/set.v0"
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

		if msg.Type == 1 {
			zap.S().Info("这是一条私信:", msg.Content)
			tarNode, ok := clientMap[msg.TargetId]
			if !ok {
				zap.S().Info("不存在对应的node", msg.TargetId)
				return
			}

			tarNode.DataQueue <- data
			fmt.Println("发送成功：", string(data))
		}

	}
}
