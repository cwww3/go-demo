package main

import (
	"fmt"
	"os"
)

func main() {
	// 获得读写管道
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Printf("get read and write pipe fail err=%v\n", err)
		return
	}
	content := "just for test"

	// 并发执行不会阻塞
	// 但是测试时发现 基于go1.13 1.14 顺序执行也没有出现阻塞
	_, err = w.Write([]byte(content))
	if err != nil {
		fmt.Printf("write fail err=%v\n", err)
		return
	}
	//go func() {
	data := make([]byte, 100)
	_, err = r.Read(data)
	if err != nil {
		fmt.Printf("write fail err=%v\n", err)
		return
	}
	//}()

	fmt.Println(string(data))
	//time.Sleep(time.Second)
}
