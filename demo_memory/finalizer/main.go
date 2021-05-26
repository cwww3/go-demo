package main

import (
	"fmt"
	"runtime"
)

func main() {
	type T struct {
		v [1 << 20]int
		//t *T
	}

	var finalizer = func(t *T) {
		fmt.Println("finalizer called")
	}

	var x T

	// 此SetFinalizer函数调用将使x逃逸到堆上。
	runtime.SetFinalizer(&x, finalizer)

	// 下面这行将形成一个包含x和y的循环引用值组。
	// 这有可能造成x和y不可回收。
	//x.t, y.t = &y, &x // y也逃逸到了堆上。
}
