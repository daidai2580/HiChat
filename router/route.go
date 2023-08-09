package router

import (
	"HiChat/service"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	route := gin.Default()

	v1 := route.Group("v1")

	user := v1.Group("user")
	{
		user.GET("/list", service.List)
		user.POST("/login_pw", service.LoginByNameAndPasseWord)
		user.POST("/new", service.NewUser)
		user.DELETE("/delete", service.DeleteUser)
		user.POST("/update", service.UpdateUser)
	}
	return route
}
