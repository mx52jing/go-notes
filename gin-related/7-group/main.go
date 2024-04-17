package main

import "github.com/gin-gonic/gin"

type SuccessRes struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data any `json:"data"`
}

func main() {
	router := gin.Default()

	userRouter := router.Group("/user")
	{
		userRouter.GET("/info", func (ctx *gin.Context) {
			ctx.JSON(200, SuccessRes{0, "成功", gin.H{"userName": "张三", "age": 22} })
		})
		userRouter.GET("/id", func (ctx *gin.Context) {
			ctx.JSON(200, SuccessRes{0, "成功", 1 })
		})
		userV1 := userRouter.Group("/v1")
		{
			userV1.GET("/name", func (ctx *gin.Context) {
				ctx.JSON(200, SuccessRes{0, "成功", "张三" })
			})
		}
		{
			userV1.GET("/401", func (ctx *gin.Context) {
				ctx.JSON(401, SuccessRes{-1, "无权限", nil})
			})
		}
	}
	router.Run(":9100")
}