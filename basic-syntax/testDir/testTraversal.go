package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"sync"
)

var (
	ignoreFileNameReg = regexp.MustCompile(`(\.DS_Store)`) // 忽略的文件名称正则
	count int
	mLock sync.Mutex
)

func traversalDir(dirPath string) []string {
	allFiles := make([]string, 0)
	filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fileInfo, fileInfoErr := d.Info()
		if fileInfoErr != nil {
			return fileInfoErr
		}
		fileName := fileInfo.Name()
		matchResult := ignoreFileNameReg.FindAllSubmatchIndex([]byte(fileName), -1)
		// 判断当前path是一个文件，并且文件名称不为 .DS_Store
		if fileInfo.Mode().IsRegular() && len(matchResult) == 0 {
			// fmt.Println(path, "是文件", fileName, ignoreFileNameReg.FindAllSubmatchIndex([]byte(fileName), -1))
			allFiles = append(allFiles, path)
		}
		return nil
	})
	return allFiles
}

func processFile(file string) {
	// fmt.Println(file, "file")
	mLock.Lock()
	defer mLock.Unlock()
	count += 1
}

func main() {
	dirPath := "/Users/mx/Desktop/mx/util-related/static-blog"
	allFiles := traversalDir(dirPath)
	filesCount := len(allFiles)
	if filesCount == 0 {
		return
	}
	fmt.Println("filesCount", filesCount)
	var wg sync.WaitGroup
	wg.Add(filesCount)
	for _, file := range allFiles {
		go func(curFile string) {
			defer wg.Done()
			processFile(curFile)
		}(file)
	}
	wg.Wait()
	fmt.Println("<<<count>>>", count)
}