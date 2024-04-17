package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getHeaderHandler(ctx *gin.Context) {
	// 首字母大小写不区分  单词与单词之间用 - 连接
	fmt.Printf("ctx.GetHeader(\"user-agent\"): %v\n", ctx.GetHeader("user-agent"))
	fmt.Printf("ctx.GetHeader(\"admin-token\"): %v\n", ctx.GetHeader("admin-token"))
	fmt.Printf("ctx.GetHeader(\"platform\"): %v\n", ctx.GetHeader("platform"))

	fmt.Printf("ctx.Request.Header.Get(\"platform\"): %v\n", ctx.Request.Header.Get("platform"))
	// Header 是一个普通的 map[string][]string
}

func main() {
	router := gin.Default()
	router.GET("/get_header", getHeaderHandler)
	router.Run(":9100")
}