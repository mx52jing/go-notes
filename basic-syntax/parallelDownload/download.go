package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

func downloadFile(remoteUrl, filename string) {
	resp, err := http.Get(remoteUrl)
	if err != nil {
		fmt.Println("获取文件失败", err)
		return
	}
	defer resp.Body.Close()
	localFilePath, err := generateFilePath(filename)
	if err != nil {
		fmt.Println("获取保存路径失败", err)
		return
	}
	writeFile, err := os.OpenFile(localFilePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("os.OpenFile err", err)
		return
	}
	defer writeFile.Close()
	// reader := bufio.NewReader(resp.Body)
	// writer := bufio.NewWriter(writeFile)
	// for {
	// 	byte, readByteErr := reader.ReadByte()
	// 	if readByteErr == io.EOF {
	// 		fmt.Println("======读取完毕======")
	// 		break
	// 	}
	// 	writer.WriteByte(byte)
	// }
	// writer.Flush()
	// 使用io.Copy效率更高
	_, copyErr := io.Copy(writeFile, resp.Body)
	if copyErr != nil {
		fmt.Println("Copy File error", copyErr)
	}
}

func generateFilePath(filename string) (path string, err error) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前执行程序目录失败", err)
		return "", err
	}
	dirPath := filepath.Join(dir, "assets")
	if _, statErr := os.Stat(dirPath); os.IsNotExist(statErr) {
		os.MkdirAll(dirPath, os.ModePerm)
	}
	path = filepath.Join(dirPath, filename)
	return
}

func main() {
	var wg sync.WaitGroup
	fileUrls := []string{
		"https://s1.best-wallpaper.net/wallpaper/m/1701/Brown-dog-look-at-you_m.webp",
		"https://www.xdsucai.cn/wp-content/uploads/2021/02/1613135056-588f17728b565e7.png",
		"https://s1.best-wallpaper.net/wallpaper/m/2307/Mountains-green-cliff-village-trees-China_m.webp",
		"https://www.w3school.com.cn/example/html5/mov_bbb.mp4",
		"https://picnew3.photophoto.cn/20081229/ziranfengjingtupian-13317683_1.jpg",
		"https://s1.best-wallpaper.net/wallpaper/m/1802/Road-fence-wood-house-sunset_m.webp",
		"https://media.w3.org/2010/05/sintel/trailer.mp4",
	}
	wg.Add(len(fileUrls))
	for idx, val := range fileUrls {
		ext := filepath.Ext(val)
		fileName := fmt.Sprintf("%d%s", idx+1, ext)
		go func(url, fileName string) {
			defer wg.Done()
			downloadFile(url, fileName)
		}(val, fileName)
	}
	wg.Wait()
}
