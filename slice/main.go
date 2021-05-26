package main

import (
	"fmt"
	"reflect"
)

func main() {
	//s := []byte("123456789")
	//a := s[1:1:1]
	//fmt.Println(len(a), cap(a))

	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[0:3]
	s3 := s1[0:3]
	fmt.Println(reflect.DeepEqual(s2, s3))
	s4 := s1[0:4]
	fmt.Println(reflect.DeepEqual(s2, s4))
}
