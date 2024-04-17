package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	UserName string `json:"user_name" binding:"required"`
	Age      uint8  `json:"age" binding:"required,gte=18"`
	Gender   string `json:"gender" binding:"required,oneof=male female"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// @Summary	添加新用户
// @Produce	json
// @Param		user	body		User		true	"用户信息数据"
// @Success	200		{object}	Response	"接口响应成功"
// @Failure	500		{object}	Response	"接口响应失败"
// @Router		/user/add [POST]
func AddUser(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("参数错误：%s", err)
		ctx.JSON(http.StatusOK, &Response{
			Code: -1,
			Msg:  "参数错误，请检查后再提交",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, &Response{
		Code: 0,
		Msg:  "success",
		Data: user,
	})
}
