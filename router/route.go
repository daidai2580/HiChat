package router

import (
	"HiChat/middlewear"
	"HiChat/service"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	route := gin.Default()
	route.Static("/assets", "./assets")
	route.StaticFile("/upload1691654470816168831", "./assets/upload1691654470816168831.txt")
	v1 := route.Group("v1")

	user := v1.Group("user")
	part := v1.Group("part")
	{
		part.POST("/create", service.CreatePart)
		part.GET("/findByOid", service.FinPartByOid)
	}

	{
		user.GET("/list", middlewear.JWY(), service.List)
		user.POST("/login_pw", service.LoginByNameAndPasseWord)
		user.POST("/new", service.NewUser)
		user.DELETE("/delete", middlewear.JWY(), service.DeleteUser)
		user.POST("/update", middlewear.JWY(), service.UpdateUser)
		user.GET("/sendMsg", service.SendUserMsg)

	}

	relation := v1.Group("relation").Use(middlewear.JWY())
	{
		relation.POST("/list", service.FriendList)
		relation.POST("/add", service.AddFriendByName)
		relation.POST("/new_group", service.NewGroup)
		relation.POST("/group_list", service.GroupList)
		relation.POST("/join_group", service.JoinGroup)
	}

	upload := v1.Group("upload")
	{
		upload.POST("/image", service.Image)
	}
	news := v1.Group("news")
	{
		news.GET("/list", service.NewsList)
		news.GET("/listAndUser", service.GetNewsAndUser)
		news.GET("/listAndUser2", service.GetNewsAndUser2)
		news.POST("/add", service.AddNews)
	}
	v1.POST("/user/redisMsg", service.RedisMsg)

	return route
}
