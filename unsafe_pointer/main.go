package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Println(str2slice("abcd"))
	fmt.Println(slice2str(str2slice("abcd")))
}

func str2slice(str string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	slice := reflect.SliceHeader{
		Data: strHeader.Data,
		Len:  strHeader.Len,
		Cap:  strHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}

func slice2str(data []byte) string {
	//sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	//str := reflect.StringHeader{
	//	Data: sliceHeader.Data,
	//	Len:  sliceHeader.Len,
	//}
	//return *(*string)(unsafe.Pointer(&str))

	return *(*string)(unsafe.Pointer(&data))
}
