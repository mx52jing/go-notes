package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "swag-docs/docs"
	"swag-docs/student"
	"swag-docs/user"
)

func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/user/add", user.AddUser)
	router.GET("/get/:id", student.GetUser)         // get方法，restful风格的url
	router.POST("/update_user", student.UpdateUser) // post方法
	router.Run(":9600")
}
