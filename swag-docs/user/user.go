package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	UserName string `json:"user_name" binding:"required"`
	Age      uint8  `json:"age" binding:"required"`
	Gender   string `json:"gender" binding:"required,oneof=male female"`
	Birthday string `json:"birthday"`
}

// AddUser @Summary	添加新用户
// @Produce	json
// @Param		user	body		User		true	"用户信息"
// @Success	200		string	string	"添加用户成功"
// @Failure	500		string	string	"添加用户失败"
// @Router		/user/add [POST]
func AddUser(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数错误",
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": user,
	})
}
