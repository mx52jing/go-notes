package main

import (
	"bytes"
	"fmt"
	"github.com/bytedance/sonic"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var requestUrl = "http://172.16.40.168:6527/api/v1"

func viewResponseInfo(response *http.Response) {
	fmt.Printf("response Proto：%s\n", response.Proto)
	fmt.Printf("response Status：%s\n", response.Status)
	fmt.Printf("response StatusCode：%d\n", response.StatusCode)
	fmt.Println("response header")
	// 遍历查看一下response Header 的内容
	for key, value := range response.Header {
		fmt.Printf("response Header中【%s】的值为 %v\n", key, value)
	}
	fmt.Println("response body")
	io.Copy(os.Stdout, response.Body) //两个io数据流的拷贝
	os.Stdout.WriteString("\n")
}

/**
GET 请求
*/
// 不带参数的get 请求
func getNoParams() {
	response, err := http.Get(requestUrl + "/user/list")
	if err != nil {
		log.Printf("http Get 请求失败：%s", err)
		return
	}
	// 注意：一定要调用response.Body.Close()，否则会协程泄漏（同时引发内存泄漏）
	defer response.Body.Close()
	viewResponseInfo(response)
}

// 带参数的get 请求，参数直接拼接在url后面
func getWithParams() {
	response, err := http.Get(requestUrl + "/user/list?name=zhang&memory=256")
	if err != nil {
		log.Printf("http Get 请求失败：%s", err)
		return
	}
	// 注意：一定要调用response.Body.Close()，否则会协程泄漏（同时引发内存泄漏）
	defer response.Body.Close()
	viewResponseInfo(response)
}

/*
*
HEAD 请求
HEAD 请求类似于GET，不过HEAD方法只能获取到http response 报文头部，取不到response.body HEAD请求通常用来验证一个url是否存活
*/
func headRequest() {
	response, err := http.Head(requestUrl + "/user/list")
	if err != nil {
		log.Printf("http Head 请求失败：%s", err)
		return
	}
	// 注意：一定要调用response.Body.Close()，否则会协程泄漏（同时引发内存泄漏）
	defer response.Body.Close()
	//状态码为200就说明url存活
	viewResponseInfo(response)
}

/*
*
POST
*/

// POST请求 无请求体 只有query
func postNoBodyRequest() {
	response, err := http.Post(requestUrl+"/user/add?name=li&age=2", "application/json", nil)
	if err != nil {
		log.Printf("http Post 请求失败：%s", err)
	}
	defer response.Body.Close()
	fmt.Println("response body")
	io.Copy(os.Stdout, response.Body) //两个io数据流的拷贝
	os.Stdout.WriteString("\n")
}

type UserRequest struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}
type UserResponse struct {
	UserName string `json:"userName"`
	Age      int    `json:"age"`
}

type ResponseData struct {
	Code int
	Msg  string
	Data UserResponse
}

// POST请求 有请求体 请求体类型为【字符串】 只有query
func postRequestWithBody() {
	requestData, _ := sonic.Marshal(&UserRequest{
		UserId: "dfjkgh893y4bhseg",
		Token:  "sdfmsdjbfshfvjkgf",
	})
	//request body是byte切片
	response, err := http.Post(requestUrl+"/user/add?name=li&age=2", "application/json", bytes.NewReader(requestData))
	if err != nil {
		log.Printf("http Post 请求失败：%s", err)
	}
	defer response.Body.Close()
	fmt.Printf("response Status：%s\n", response.Status)
	fmt.Printf("response StatusCode：%d\n", response.StatusCode)
	// 将response body流全部读取到byte切片中
	responseByte, err := io.ReadAll(response.Body) //两个io数据流的拷贝
	if err != nil {
		log.Printf("读取 reponse body 流失败：%s", err)
	}
	var responseData ResponseData
	sonic.Unmarshal(responseByte, &responseData)
	fmt.Printf("请求到的响应式为：%v\n", responseData)
}

// post请求，请求参数放在form表单里
// 接口传的值 通常是一个切片
func formRequest() {
	formRequestData := url.Values{
		"name": []string{"狄仁杰"},
		"age":  []string{"22"},
	}
	response, err := http.PostForm(requestUrl+"/user/edit", formRequestData)
	if err != nil {
		log.Printf("http Form 请求失败：%s", err)
		return
	}
	viewResponseInfo(response)
}

// 通过http.Client发送请求是一种万能的方式，可以涵盖以上的所有方式，并且可以设计Header和Cookie(以上方法不能设置)
func createRequestByClient() {
	userData := &UserRequest{
		UserId: "987586",
		Token:  "98475_(ABC)+#@KJ",
	}
	userByte, err := sonic.Marshal(userData)
	if err != nil {
		log.Printf("json序列化失败：%s", err)
		return
	}
	requestByte := bytes.NewReader(userByte)
	// 构建request
	req, err := http.NewRequest("POST", requestUrl+"/user/add", requestByte)
	// 设置一些请求头
	req.Header.Add("Authorization", "Bearer dev")
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Printf("构建请求失败：%s", err)
		return
	}
	// 构造http client 这里可以配置一些请求设置
	httpClient := &http.Client{
		Timeout: 6000 * time.Second,
	}
	//response, err := http.DefaultClient.Do()
	response, err := httpClient.Do(req)
	if err != nil {
		log.Printf("发送请求失败：%s", err)
		return
	}
	viewResponseInfo(response)
}

func main() {
	//getNoParams()
	//getWithParams()
	//headRequest()
	//postNoBodyRequest()
	//postRequestWithBody()
	//formRequest()
	createRequestByClient()
}
