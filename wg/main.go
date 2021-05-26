package main

import (
	"fmt"
	"os"
)

func main() {
	//sync.WaitGroup{}
	env := os.Getenv("env")
	fmt.Println(env)
	//ctx := context.Background()
	//cancelCtx, cancelFunc := context.WithCancel(ctx)
	//go func(ctx context.Context) {
	//	done := ctx.Done()
	//	for {
	//		select {
	//		case <-done:
	//
	//			fmt.Println("cancel")
	//			return
	//		default:
	//			time.Sleep(time.Second)
	//			fmt.Println("---")
	//		}
	//	}
	//}(cancelCtx)
	//time.Sleep(time.Second * 3)
	//cancelFunc()
}
