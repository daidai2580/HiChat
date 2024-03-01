package main

import (
	_ "HiChat/docs"
	"HiChat/initialize"
	"HiChat/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		127.0.0.1:8000
//	@BasePath	/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	initialize.InitConfig()
	//初始化数据库
	initialize.InitDB()
	initialize.InitRedis()
	//初始化日志
	initialize.InitLogger()

	route := router.Route()
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := route.Run(":8000")
	if err != nil {
		return
	}
}
