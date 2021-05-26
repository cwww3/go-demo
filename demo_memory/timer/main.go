package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

/**
time.After oom 验证demo
*/
func main() {
	ch := make(chan string, 100)

	go func() {
		for {
			ch <- "asong"
		}
	}()
	go func() {
		// 开启pprof，监听请求
		ip := "127.0.0.1:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
		}
	}()
	t := time.NewTimer(time.Minute * 3)
	defer t.Stop()
	for {
		t.Reset(time.Minute * 3)
		select {
		case <-ch:
		case <-t.C:
		}
	}
}
