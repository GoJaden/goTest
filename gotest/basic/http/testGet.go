package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	CheckVersion()
}

const url string = "https://itunes.apple.com/cn/lookup?id=1247810530"

func CheckVersion() {

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		//todo 日志记录
		fmt.Println(err)
	}
	//读取服务器返回的文件大小
	//aasize, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	ioutil.ReadAll()
	/*
		f ,err := os.Create("1.txt")
		if err != nil {
			panic(err)
		}
		_, err1 := io.Copy(f, resp.Body)
		if err1 != nil{
			fmt.Println(err)
		}

		f ,err2 := os.Open("1.txt")
		if err2 != nil{

		}
		buffer := make([]byte ,getFileSize("1.txt"))
		_,err3:=f.Read(buffer)
		if err3!=nil{
			fmt.Println(err3)
		}
		fmt.Println(string(buffer))
		//json.Unmarshal()*/

}

func getFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}
