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
	return route
}
