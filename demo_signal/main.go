package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	sigRecv1 := make(chan os.Signal, 1)
	sigRecv2 := make(chan os.Signal, 1)
	// 自行处理SIGINT和SIGQUIT信号
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	sigs2 := []os.Signal{syscall.SIGINT}
	signal.Notify(sigRecv1, sigs1...)
	signal.Notify(sigRecv2, sigs2...)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for sig := range sigRecv1 {
			fmt.Printf("no1 Recieved a signal: %s\n", sig)
		}
	}()
	go func() {
		defer wg.Done()
		for sig := range sigRecv2 {
			fmt.Printf("no2 Recieved a signal: %s\n", sig)
		}
	}()
	// 恢复sigRecv1中接收到的信号的默认操作
	//time.Sleep(3 * time.Second)
	//signal.Stop(sigRecv1)
	//close(sigRecv1)
	wg.Wait()

}
