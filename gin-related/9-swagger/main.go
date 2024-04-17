package main

import (
	"gin-related/9-swagger/api"
	_ "gin-related/9-swagger/docs" // 这个docs文件模块必须要填引入
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	// 下面这个必须要有
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/user/add", api.AddUser)
	router.Run(":6100")
}
