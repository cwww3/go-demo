package main

import (
	"fmt"
	"project/generate/painkiller"
)

//go:generate stinger -type=PILL
func main() {
	res := painkiller.Acetaminophen
	fmt.Println(res.String())
}
