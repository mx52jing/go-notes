package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

// 上传单个文件
func singleFileHanlder(ctx *gin.Context) {
	name := ctx.PostForm("name")
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("ctx.FormFile err",err)
	}
	savePath := path.Join("uploads", file.Filename)
	fmt.Println(name, file.Filename, file.Size, "============")
	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		ctx.JSON(http.StatusOK, gin.H{ 
			"code": -1, 
			"msg": "文件上传失败", 
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{ 
		"code": 0, 
		"msg": "文件上传成功", 
		"data": gin.H{ "filepath": savePath },
	})
}

// 上传多个文件
func multiFileHanlder(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		fmt.Println("ctx.MultipartForm err =>>>", err)
	}
	files := form.File["file"]
	var paths []string
	for _, file := range files {
		savePath := path.Join("uploads", file.Filename)	
		if err := ctx.SaveUploadedFile(file, savePath); err != nil {
			ctx.JSON(http.StatusOK, gin.H{ 
				"code": -1, 
				"msg": "文件上传失败", 
				"data": nil,
			})
			return
		}
		paths = append(paths, savePath)
	}
	ctx.JSON(http.StatusOK, gin.H{ 
		"code": 0, 
		"msg": "文件上传成功", 
		"data": paths,
	})
}

func main()  {
	router := gin.Default()
	//MaxMultipartMemory 参数用于设置内存缓冲区的最大容量，也就是上传的数据超过该容量时，将会把文件流写入磁盘中，而不是一直保留在内存中。这是为了防止攻击者发送大量数据导致服务器内存耗尽，从而引发拒绝服务攻击（DoS）。
	// 8 << 20 表示将数字 8 左移 20 位，即将数字 8 转换为字节为单位的容量值，得到的结果是 8 MB
	router.MaxMultipartMemory = 8 << 20 // 8M
	router.Static("/static", "assets")
	router.Static("/tpl", "templates")
	router.POST("/single_file", singleFileHanlder)
	router.POST("/multi_file", multiFileHanlder)

	router.Run(":9100")
}