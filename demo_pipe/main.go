package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 创建命令
	cmd := exec.Command("echo", "-n", "My first command")
	// 获得一个输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("get stdout pipe fail err=%v\n", err)
	}
	// 执行命令
	err = cmd.Start()
	if err != nil {
		fmt.Printf("start command fail err=%v\n", err)
	}
	// 读取管道信息
	data := make([]byte, 20)
	n, err := stdout.Read(data)
	if err != nil {
		fmt.Printf("read from pipe err=%v\n", err)
	}
	fmt.Println(string(data[:n]))

}
