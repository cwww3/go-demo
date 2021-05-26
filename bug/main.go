package main

import (
	"fmt"
	"gitee.com/hzsuoyi/unite-go/uni"
	"gitee.com/hzsuoyi/unite-go/unite"
	"sync"
)

type demo struct {
	Member interface{}
}

func main() {
	var wg = sync.WaitGroup{}
	wg.Add(1)
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("PANIC ", err)
	//		debug.PrintStack()
	//	}
	//}()
	//a := demo{Member: "11"}
	//ss := a.Member.(string)
	//fmt.Println(ss)
	//fmt.Println(ok)
	//var m demo
	//item := &demo{
	//	Member: struct {
	//		Name string `json:"name"`
	//	}{
	//		//Name: "cwww",
	//	},
	//}
	ss := get()
	go func() {
		a := (ss).(*uni.Client).SendSubscribeMessage(unite.WxaTemplate{})
		fmt.Println(a)
	}()
	wg.Wait()
	//fmt.Println(a)
}
func get() unite.CommonClient {
	return nil
}
