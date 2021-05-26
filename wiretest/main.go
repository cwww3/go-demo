package main

import (
	"fmt"
	"github.com/Fs02/wire"
)

type A struct {
	//Value string `wire:"value"`
}
type B struct {
	A     A      `wire:"a"`
	Value string `wire:"value"`
}

func main() {
	c := wire.New()
	a := A{}
	//b := B{a, "b"}
	c.Connect(&a)
	//c.Connect(&b)
	c.Apply()
	var aa A
	wire.Resolve(&aa)
	fmt.Println(aa)
}
