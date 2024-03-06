package service

import (
	"HiChat/dao"
	"HiChat/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// @Tags 新闻
// @Summary	新闻列表
// @Produce	json
// @Success	200			{object}	string	"成功"
// @Failure	400			{object}	string	"请求错误"
// @Failure	500			{object}	string	"内部错误"
// @Router		/news/list [get]
func NewsList(ctx *gin.Context) {
	list, err := dao.GetNewsList()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "获取新闻列表失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, list)
}

// @Tags 新闻
// @Summary	新闻列表2
// @Produce	json
// @Success	200			{object}	string	"成功"
// @Failure	400			{object}	string	"请求错误"
// @Failure	500			{object}	string	"内部错误"
// @Router		/news/listAndUser [get]
func GetNewsAndUser(ctx *gin.Context) {
	list, err := dao.GetNewsAndUser()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "获取新闻列表失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, list)
}

// @Tags 新闻
// @Summary	新增新闻列表
// @Produce	json
// @Param		news			body		string	true	"内容"
// @Success	200			{object}	string	"成功"
// @Failure	400			{object}	string	"请求错误"
// @Failure	500			{object}	string	"内部错误"
// @Router		/news/add [post]
func AddNews(ctx *gin.Context) {
	var news = models.News{}
	err := ctx.ShouldBind(&news)
	if err != nil {
		zap.S().Info("解析json失败")
		fmt.Println("错误信息：", err)
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "解析json失败！",
			"data":    news,
		})
		return
	}

	dao.AddNews(news)
	ctx.JSON(200, gin.H{
		"code":    0,
		"message": "新增新闻成功",
		"data":    news,
	})
}
