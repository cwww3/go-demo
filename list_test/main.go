package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func do() {
	resp, _ := http.Get("https://www.baidu.com")
	//_, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
}
func main() {
	fmt.Printf("此时goroutine个数= %d\n", runtime.NumGoroutine())
	//num := 6
	//for index := 0; index < num; index++ {
	//	do()
	//}
	do()
	fmt.Printf("此时goroutine个数= %d\n", runtime.NumGoroutine())
}
