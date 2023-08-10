package router

import (
	"HiChat/middlewear"
	"HiChat/service"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	route := gin.Default()

	v1 := route.Group("v1")

	user := v1.Group("user")

	{
		user.GET("/list", middlewear.JWY(), service.List)
		user.POST("/login_pw", service.LoginByNameAndPasseWord)
		user.POST("/new", service.NewUser)
		user.DELETE("/delete", middlewear.JWY(), service.DeleteUser)
		user.POST("/update", middlewear.JWY(), service.UpdateUser)
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

	return route
}
