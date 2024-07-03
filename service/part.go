package service

import (
	"HiChat/dao"
	"HiChat/global"
	"HiChat/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @Tags 物料
// @Summary	创建物料信息
// @Produce	json
// @Param		body body models.Part true "请求body"
// @Success	200			{object}	string	"成功"
// @Failure	400			{object}	string	"请求错误"
// @Failure	500			{object}	string	"内部错误"
// @Router		/part/create [post]
func CreatePart(ctx *gin.Context) {
	b, _ := ctx.GetRawData()
	// 定义map或结构体
	var part models.Part
	// 反序列化
	_ = json.Unmarshal(b, &part)

	m, err := dao.CreatePart(part)
	if err != nil {
		ctx.JSON(200,
			gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
		return
	}
	ctx.JSON(http.StatusOK, m)

}

// @Tags 物料
// @Summary	根据物料oid查找物料
// @Produce	json
// @Param		oid query  string true "物料oid"
// @Success	200			{object}	string	"成功"
// @Failure	400			{object}	string	"请求错误"
// @Failure	500			{object}	string	"内部错误"
// @Router		/part/findByOid [get]
func FinPartByOid(ctx *gin.Context) {
	oid := ctx.Query("oid")
	lock, _ := global.RedisDB.Get(ctx, oid).Result()
	if lock != "" {
		ctx.JSON(200, gin.H{"code": -1, "msg": "baocuoneir"})
		return
	}
	set := global.RedisDB.Set(ctx, oid, "string", time.Second*60)
	print(set)
	part, err := dao.FindPartByOid(oid)
	if err != nil {
		ctx.JSON(200, gin.H{"code": -1, "msg": "baocuoneir"})
		return
	}
	ctx.JSON(http.StatusOK, part)
}
