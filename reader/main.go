package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	// 数组 转 bytes.Reader
	//var data []byte
	//r1 := bytes.NewReader(data)

	// 数组转 bytes.Buffer
	//r2 := bytes.NewBuffer(data)

	// 字符串转bytes.Buffer
	var s string = "123"
	r3 := bytes.NewBufferString(s)

	// bytes.Reader 实现了io.Reader  bytes.Buffer 实现了io.Reader和Writer
	// 默认buf缓冲区大小4096 可以自己指定大小 bufio.NewReaderSize(r1,9999)

	// bytes.Buffer， bytes.Reader 转bufio.Reader
	//r4 := bufio.NewReader(r1)
	//r4 = bufio.NewReader(r2)
	r4 := bufio.NewReader(r3)
	value, err := ioutil.ReadAll(r4)
	fmt.Println(r4.Buffered())
	fmt.Println(value, err)
	v := new(string)
	// 等用户输入
	fmt.Println(fmt.Scanln(v))
	fmt.Println(*v)
}
