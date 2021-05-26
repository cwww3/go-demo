package main

import "fmt"

type b int

func (a) name() {
	fmt.Println("1")
}
func (b) name() {
	fmt.Println("1")
}

func main() {
	var v int64 = 1111

	fmt.Println(a(v))
}
