package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func queryHanlder(ctx *gin.Context) {
	user_name := ctx.Query("userName")
	age := ctx.Query("age")
	// 可以设置默认值
	like := ctx.DefaultQuery("like", "篮球")
	// 判断参数是否存在
	password, isPasswordExist := ctx.GetQuery("password")
	// ctx.Query 张三 88
	// fmt.Println("ctx.Query", user_name, age)
	fmt.Println("ctx.Query", user_name, age, like)
	ctx.JSON(http.StatusOK, gin.H{
		"user_name": user_name,
		"age": age,
		"like": like,
		"password": password,
		"isPasswordExist": isPasswordExist,
	})
}

type User struct {
	UserName string `form:"user_name" json:"user_name"`
	Age int `form:"age" json:"age"`
	Address string `form:"address" json:"address" binding:"required"`
}

func queryBindHanlder(ctx *gin.Context) {
	var user User
	// err := ctx.BindQuery(&user)
	err := ctx.ShouldBindQuery(&user)
	if err != nil {
		fmt.Println("bindQuery Error", err)
	}
	ctx.JSON(http.StatusOK, user)
}


func multiQueryHanlder(ctx *gin.Context) {
	// usernames := ctx.QueryArray("userName")
	values, ok := ctx.GetQueryArray("userName")
	// ctx.JSON(http.StatusOK, usernames)
	ctx.JSON(http.StatusOK, gin.H{ "usernames": values, "ok": ok })
}

func mapQueryHanlder(ctx *gin.Context) {
	infoMap := ctx.QueryMap("info")
	infoMap, isInfoMapExist := ctx.GetQueryMap("info")
	// ctx.JSON(http.StatusOK, infoMap)
	ctx.JSON(http.StatusOK, gin.H{ "infoMap": infoMap, "isInfoMapExist": isInfoMapExist })
}


func paramHanlder(ctx *gin.Context) {
	userName := ctx.Param("userName")
	age := ctx.Param("age")
	ctx.JSON(http.StatusOK, gin.H{"userName": userName, "age": age})
}

func postFormHanlder(ctx *gin.Context) {
	userName := ctx.PostForm("userName")
	password := ctx.PostForm("password")
	address := ctx.PostFormArray("address")
	addressMap := ctx.PostFormMap("addressMap")
	like := ctx.DefaultPostForm("like", "羽毛球")
	ctx.JSON(http.StatusOK, gin.H{
		"userName": userName,
		"password": password,
		"address": address,
		"addressMap": addressMap,
		"like": like,
	})
}

func postJsonHanlder(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("ctx.ShouldBindJSON", err)
	}
	ctx.JSON(http.StatusOK, user)
}

func fileHanlder(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		fmt.Println("ctx.MultipartForm err", err)
	}
	fmt.Println("form", form)
	files := form.File
	fmt.Println("form.file", files)
	for _, fileArray := range files {
		for _, v := range fileArray {
			// ctx.SaveUploadedFile(v, "./"+ v.Filename)
			fmt.Println(v)
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/get_query", queryHanlder)
	router.GET("/get_bind_query", queryBindHanlder)
	router.GET("/get_multi_query", multiQueryHanlder)
	router.GET("/get_map_query", mapQueryHanlder)
	router.GET("/get_param/:userName/:age", paramHanlder)
	router.POST("/post_form", postFormHanlder)
	router.POST("/post_json", postJsonHanlder)
	router.POST("/file", fileHanlder)
	router.Run(":9100")
}