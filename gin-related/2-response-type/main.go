package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理字符串
func stringHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello world")
} 

// 处理json
func jsonHandler(ctx *gin.Context) {
	// type User struct {
	// 	UserName string `json:"user_name"`
	// 	Password string `json:"-"`
	// 	Age int `json:"age"`
	// }
	// user := User{ "张三", "123456", 22}
	// ctx.JSON(http.StatusOK, user)

	// user1 := map[string]string{
	// 	"userName": "张三",
	// 	"age": "22",
	// 	"like": "乒乓球",
	// }
	// ctx.JSON(http.StatusOK, user1)
	ctx.JSON(http.StatusOK, gin.H{
		"user_name": "流失",
		"age": 22,
		"like": []string{"乒乓球", "羽毛球"},
	})
} 

// 响应xml
func xmlHandler(ctx *gin.Context) {
	// ctx.XML(http.StatusOK, "hello world")
	type User struct {
		UserName string `xml:"user_name"`
		Password string `xml:"-"`
		Age int `xml:"age"`
	}
	user := User{ "张三", "123456", 22}
	ctx.XML(http.StatusOK, user)
	// ctx.XML(http.StatusOK, gin.H{
	// 	"user_name": "流失",
	// 	"age": 22,
	// 	"like": []string{"乒乓球", "羽毛球"},
	// })
} 

// 响应yaml
func yamlHandler(ctx *gin.Context) {
	type User struct {
		UserName string `yaml:"user_name"`
		Password string `yaml:"-"`
		Age int `yaml:"age"`
	}
	user := User{ "张三", "123456", 22}
	ctx.YAML(http.StatusOK, user)
	// ctx.YAML(http.StatusOK, gin.H{
	// 	"user_name": "流失",
	// 	"age": 22,
	// 	"like": []string{"乒乓球", "羽毛球"},
	// })
} 

// 响应html
func htmlHandler(ctx *gin.Context) {
	type User struct {
		UserName string
		Age int
		Like []string
	}
	user := User{
		"张三",
		22,
		[]string {"乒乓球", "羽毛球"},
	}
	ctx.HTML(http.StatusOK, "user.html", user)
} 

func fileHanlder(ctx *gin.Context) {
	// ctx.File 从文件中提供数据
	// ctx.File("assets/img/ysy.jpeg")
	// ctx.FileAttachment  从文件中提供数据，并且可以对返回的文件重命名
	// ctx.FileAttachment("assets/img/ysy.jpeg", "a.jepg")
	var fs http.FileSystem = http.Dir("./")
	ctx.FileFromFS("assets/img/ysy.jpeg", fs)
}

func redirectHandler(ctx *gin.Context) {
	ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	// ctx.Redirect(http.StatusMovedPermanently, "/get_html")
}

func main() {
	// 初始化路由
	router := gin.Default()

	// 响应字符串
	router.GET("/get_string", stringHandler)

	// 响应JSON
	router.GET("/get_json", jsonHandler)

	// 响应XML
	router.GET("/get_xml", xmlHandler)

	// 响应YAML
	router.GET("/get_yaml", yamlHandler)

	// 使用glob规则匹配并加载所有的html模板
	// func (*gin.Engine).LoadHTMLGlob(pattern string)
	router.LoadHTMLGlob("templates/*")
	// 加载多个html文件
	// func (*gin.Engine).LoadHTMLFiles(files ...string)
	// router.LoadHTMLFiles("templates/user.html")

	// 响应HTML
	router.GET("/get_html", htmlHandler)

	//响应文件
	router.GET("/get_file", fileHanlder)

	// 提供静态文件服务
	// 比如图片路径为 assets/img/ysy.jpeg assets/img/avatar.jpeg
	// 设置静态资源实际访问路径和实际存放文件的路径
	// /s 是指实际访问的路径，assets是指服务器放置静态文件的路径
	// http://ip:9100/s/img/ysy.jpeg
	router.Static("/s", "assets")

	//和 router.Static 一样，不过可以自定义文件系统目录 http.FileSystem，gin 默认使用的是 gin.Dir()
	router.StaticFS("/ss", http.Dir("assets"))

	// 设置单个静态文件访问路径别名
	router.StaticFile("avatar.jpeg", "assets/img/avatar.jpeg")
	router.StaticFile("y.jpeg", "assets/img/ysy.jpeg")

	//和 router.StaticFile 一样，不过可以自定义文件系统目录 http.FileSystem，gin 默认使用的是 gin.Dir()
	router.StaticFileFS("sfs.jpeg", "img/ysy.jpeg", http.Dir("assets"))

	
	router.GET("/get_redirect", redirectHandler)
	// 启动http服务并监听
	router.Run(":9100")
}