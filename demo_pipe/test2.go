package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "go")
	var outputBuf1 bytes.Buffer
	// 设置输出
	cmd1.Stdout = &outputBuf1
	err := cmd1.Start()
	if err != nil {
		fmt.Printf("start command1  err=%v\n", err)
		return
	}
	// 等待命令执行完毕
	err = cmd1.Wait()
	if err != nil {
		fmt.Printf("wait command1 err=%v\n", err)
		return
	}
	// 将命令1的输出作为命令2的输入
	cmd2.Stdin = &outputBuf1
	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2
	err = cmd2.Start()
	if err != nil {
		fmt.Printf("start command2  err=%v\n", err)
		return
	}
	err = cmd2.Wait()
	if err != nil {
		fmt.Printf("wait command2 err=%v\n", err)
		return
	}
	fmt.Println(outputBuf2.String())
}
