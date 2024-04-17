package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func user1(ctx *gin.Context) {
	fmt.Println("进入user1")
	ctx.Next()
	fmt.Println("退出user1")
}

func user2(ctx *gin.Context) {
	fmt.Println("进入user2")
	ctx.Abort()
	fmt.Println("退出user2")
}

func user(ctx *gin.Context) {
	fmt.Println("进入user")
	ctx.Next()
	fmt.Println("退出user")
}

func globalMiddleware(ctx *gin.Context) {
	fmt.Println("进入globalMiddleware")
	ctx.Next()
	fmt.Println("退出globalMiddleware")	
}

func userHanlder(ctx *gin.Context) {
	fmt.Println("userHanlder")
}
func articleHanlder(ctx *gin.Context) {
	fmt.Println("articleHanlder")
}

func v1Middleware(ctx *gin.Context) {
	fmt.Println("进入v1Middleware")
	ctx.Next()
	fmt.Println("退出v1Middleware")	
}

func tagsHanlder(ctx *gin.Context) {
	fmt.Println("tagsHanlder")
}


func main() {
	router := gin.Default()

	router.Use(globalMiddleware)
	// 单个路由中间件
	// router.GET("/user", user1, user2, user)
	router.GET("/user", userHanlder)
	router.GET("/article", articleHanlder)

	v1Gorup := router.Group("/v1")
	v1Gorup.Use(v1Middleware)
	{
		v1Gorup.GET("/tags", tagsHanlder)
	}
	v2Gorup := router.Group("/v2")
	{
		v2Gorup.GET("/tags", tagsHanlder)
	}
	router.Run(":9100")
}