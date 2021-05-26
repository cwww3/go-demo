package main

import (
	"errors"
	"fmt"
)

func main() {
	err := c()
	fmt.Println(err)
}

func c() (err error) {
	a, err := get()
	fmt.Println(a)
	return
}
func get() (int, error) {
	return 1, errors.New("111")
}
