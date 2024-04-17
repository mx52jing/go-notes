package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"net/http"
)

// 创建https服务中间件
func secureMiddleware() gin.HandlerFunc {
	secureProcessor := secure.New(secure.Options{
		SSLRedirect: true,
	})
	return func(ctx *gin.Context) {
		err := secureProcessor.Process(ctx.Writer, ctx.Request)
		if err != nil {
			fmt.Printf("start up https server fail：%v\n", err)
			ctx.Abort()
			return
		}
		status := ctx.Writer.Status()
		if status > 300 && status < 399 {
			ctx.Abort()
		}
	}
}

func main() {
	engine := gin.Default()
	engine.Use(secureMiddleware())
	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello I am https")
	})
	engine.RunTLS(":9527", "config/cert.pem", "config/key.pem")
}
