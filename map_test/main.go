package main

import "fmt"

var m = make(map[string]int)

func main() {
	m["ok"]++
	fmt.Println(m["ok"])
	m["ok"]++
	fmt.Println(m["ok"])
	m["ok"]++
	fmt.Println(m["ok"])
}
