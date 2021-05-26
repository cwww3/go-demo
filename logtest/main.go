package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	name := "test"
	f, err := os.Create(fmt.Sprintf("%v-%v.log", name, time.Now().Format("20060102")))
	if err != nil {
		fmt.Println("创建失败", err)
	}
	logger := log.New(f, "", log.LstdFlags)
	logger.Printf("%v的第一条日志", name)
	logger.Printf("%v的第二条日志", name)
	err = logger.Output(1, "输出")
	if err != nil {
		fmt.Println(err)
	}
}
