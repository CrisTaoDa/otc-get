package util

import (
	"fmt"
	"io"
	"os"
)

// 写入文件
func WriteFile(content string, filePath string) (bool, error) {
	var file *os.File
	var err error
	if checkFileIsExist(filePath) {
		// 打开文件
		file, err = os.OpenFile(filePath, os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
			return false, err
		}
	} else {
		file, err = os.Create(filePath)
	}
	defer file.Close()
	// 写入内容
	_, err = io.WriteString(file, content)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// 获取项目路径
func LocalPath() string {
	path, _ := os.Getwd()
	path = path + "\\..\\"
	return path
}
