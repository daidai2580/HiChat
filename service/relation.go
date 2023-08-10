package service

import (
	"HiChat/common"
	"HiChat/dao"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type user struct {
	Name     string
	Avatar   string
	Gender   string
	Phone    string
	Email    string
	Identity string
}

func FriendList(c *gin.Context) {
	id, _ := strconv.Atoi(c.Request.FormValue("userId"))
	users, err := dao.FriendList(uint(id))
	if err != nil {
		zap.S().Info("获取好友列表失败", err)
		c.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "好友为空",
		})
		return
	}
	infos := make([]user, 0)

	for _, v := range *users {
		info := user{
			Name:     v.Name,
			Avatar:   v.Avatar,
			Gender:   v.Gender,
			Phone:    v.Phone,
			Email:    v.Email,
			Identity: v.Identity,
		}
		infos = append(infos, info)
	}
	common.RespOKList(c.Writer, infos, len(infos))
}
