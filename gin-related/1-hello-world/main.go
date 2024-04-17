package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func UserHandle(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello world2")
} 

func main() {
	// 初始化路由
	router := gin.Default()
	router.GET("/user", UserHandle)
	// 启动http服务并监听
	router.Run(":9100")
}