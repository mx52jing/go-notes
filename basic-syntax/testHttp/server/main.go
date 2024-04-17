package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)


func UserHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("req method", req.Method)
	fmt.Println("req url", req.URL)
	// 返回响应
	res.Write([]byte ("请求到了"))
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case http.MethodGet:
			html, err := os.ReadFile("home.html")
			if err != nil {
				fmt.Println("read home.html err", err)
				return
			}
			res.Write(html)
		case http.MethodPost:
			// 获取req.body
			bodyData, err := io.ReadAll(req.Body)
			if err != nil {
				fmt.Println("获取req.Body出错", err)
				return
			}
			fmt.Println(string(bodyData))
			res.Header().Set("admin-token", "*#HJEVJKG*^C@!$JN")
			res.Write([]byte(bodyData))
	}
}

func GetJsonHandler(res http.ResponseWriter, req *http.Request) {
	jsonData := make(map[string]string)
	jsonData["name"] = "张环"
	jsonData["like"] = "跟随元芳"
	respData, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("转换json数据失败", err)
		return
	}
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(respData)
}

func GetStructJson(res http.ResponseWriter, req *http.Request) {
	type UserInfo struct {
		UserName string
		Age int
		Like []string
	}
	userInfo := UserInfo{
		UserName: "张三",
		Age: 22,
		Like: []string {"篮球", "乒乓球", "羽毛球"},
	}
	respData, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println("转换json数据失败", err)
		return
	}	
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(respData)
}

func main() {
	// 绑定回调
	http.HandleFunc("/user", UserHandler)
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/get_json", GetJsonHandler)
	http.HandleFunc("/get_struct_json", GetStructJson)
	// 监听并启动服务
	fmt.Println("Server is Running at http://127.0.0.1:9100")
	http.ListenAndServe(":9100", nil)
}