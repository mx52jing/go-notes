package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

const requestUrl = "https://xxxxxxxxxx?mobileType=-1&mobileArea=-1&isLength=1&startTime=2024-06-10+00:00:00&endTime=2024-06-26+23:59:59&ecid=108683&msgtype=0&page=1&limit=100"

// 接口请求会来的数据结构
type ResponseData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data ResData
}

type ResData struct {
	Records []Record
}

type Record struct {
	UserID     string  `json:"userid"`
	Phone      string  `json:"phone"`
	MobileType int     `json:"mobileType"`
	MobileArea int     `json:"mobileArea"`
	IsLength   int     `json:"isLength"`
	ErrorCode  string  `json:"errorCode"`
	ErrorDesc  *string `json:"errorDesc"` // 使用指针以允许 null 值
	CustID     string  `json:"custid"`
	PtMsgID    string  `json:"ptmsgid"`
	SendTime   string  `json:"sendTime"`
	RecvTime   string  `json:"recvTime"`
	Message    string  `json:"message"`
	PkTotal    int     `json:"pktotal"`
	Strip      *string `json:"strip"` // 使用指针以允许 null 值
	WordNum    int     `json:"wordNum"`
	TotalWords int     `json:"totalwords"`
	MsgSrcIP   string  `json:"msgsrcip"`
}

func fetchSmsSendData() {
	req, err := http.NewRequest("GET", requestUrl, nil)
	// 设置需要的请求头
	req.Header.Add("Authorization", "Authorization")
	req.Header.Add("Cookie", "Cookie")
	if err != nil {
		log.Printf("构建请求失败：%s", err)
		return
	}
	res, err := http.DefaultClient.Do(req)
	log.Println("请求中.............")
	if err != nil {
		log.Printf("请求失败=> %s", err)
		return
	}
	log.Println("请求完成")
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("读取res.body失败=> %s", err)
		return
	}
	// 解析响应 并获得想要的数据
	var responseData ResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Printf("序列化res.body失败=> %s", err)
		return
	}
	records := responseData.Data.Records
	if len(records) == 0 {
		log.Printf("<<<<<<<<<<<<<没有符合的数据>>>>>>>>>>>>>")
		return
	}
	result := make([]Record, 0)
	for _, record := range records {
		if record.ErrorCode == "M2:0010" {
			result = append(result, record)
		}
	}
	log.Printf("<<<<<<<<<<<<<文件筛选完毕，准备写入文件>>>>>>>>>>>>>")
	// 将筛选出来的数组写入 json文件中
	writeDataToFile(result)
}

func writeDataToFile(data []Record) {
	// MarshalIndent 第二个参数表示 新的一行以什么开头 第三个参数表示开头后面跟着的缩进
	jsonByteData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Printf("data 转换byte失败%s\n", err)
		return
	}
	file, err := os.Create("records.json")
	if err != nil {
		log.Printf("创建records.json文件失败=> %s\n", err)
		return
	}
	defer file.Close()
	// 使用 bufio.Writer 提高写入效率
	writer := bufio.NewWriter(file)
	// 将 JSON 数据写入文件
	_, err = writer.Write(jsonByteData)
	if err != nil {
		log.Printf("文件写入失败%s\n", err)
		return
	}
	// 刷新 确保所有缓冲数据都写入文件
	writer.Flush()
	log.Printf("<<<<<<<<<<<<<文件写入完成>>>>>>>>>>>>>")
}

func main() {
	fetchSmsSendData()
}
