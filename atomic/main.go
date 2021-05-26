package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

var value int64 = 0
var ch = make(chan int64)

func main() {
	//var i64 int64
	//fmt.Println("=====old i64 value=====")
	//fmt.Println(i64)
	//atomic.AddInt64(&i64, -3)
	//fmt.Println("=====new i64 value=====")
	//fmt.Println(i64)
	//atomic.LoadInt32
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go addValue(1, ctx)
	go addValue(2, ctx)

	for {
		select {
		case <-ch:
			cancel()
			fmt.Println(value)
			return
		}
	}

}

//不断地尝试原子地更新value的值,直到操作成功为止
func addValue(delta int64, ctx context.Context) {
	//在被操作值被频繁变更的情况下,CAS操作并不那么容易成功
	//so 不得不利用for循环以进行多次尝试
	for {
		select {
		case <-ctx.Done():
			return
		default:
			v := value
			if atomic.CompareAndSwapInt64(&value, v, v+delta) {
				//在函数的结果值为true时,退出循环
				time.Sleep(time.Second)
				ch <- delta
				return
			} else {
				fmt.Println("false", delta)
			}
		}
		//操作失败的缘由总会是value的旧值已不与v的值相等了.
		//CAS操作虽然不会让某个Goroutine阻塞在某条语句上,但是仍可能会使流产的执行暂时停一下,不过时间大都极其短暂.
	}
}
