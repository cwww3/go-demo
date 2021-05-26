package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name string
}

var m = make(map[People]int)
var n = make(map[People]int)

func main() {
	var a float64 = 1
	var b float64 = 1
	fmt.Println(reflect.DeepEqual(a, b))
	ss := People{Name: "1"}
	yy := People{Name: "1"}
	m[ss] = 1
	n[yy] = 1
	fmt.Println(reflect.DeepEqual(m, n))
	//reflect.Indirect()
}
