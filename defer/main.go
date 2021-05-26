package main

import (
	"fmt"
	"time"
)

// Ctrl + c 结束 不会触发defer

func main() {
	testInterrupt()
}

func testInterrupt() {
	defer fmt.Println("中断也能打印")
	time.Sleep(time.Second * 100)
}
