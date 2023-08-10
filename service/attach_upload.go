package service

import (
	"HiChat/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Image(c *gin.Context) {
	w := c.Writer
	req := c.Request

	file, header, err := req.FormFile("file")
	if err != nil {
		common.RespFail(w, err.Error())
		return
	}
	suffix := ".png"
	ofilName := header.Filename
	tem := strings.Split(ofilName, ".")
	if len(tem) > 1 {
		suffix = "." + tem[len(tem)-1]
	}
	name := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	temp, err := os.Create("./asset/upload/" + name)
	if err != nil {
		common.RespFail(w, err.Error())
		return
	}

	_, err = io.Copy(temp, file)
	if err != nil {
		common.RespFail(w, err.Error())
	}
	url := "./asset/upload/" + name
	common.RespOK(w, url, "发送成功")
}
