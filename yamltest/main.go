package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"reflect"
)

//var c1 = `
//key:
// sub_key:
//  v0:
//   sub_value1: 1
//   sub_value2: 1
//  v1: 1
//  v2: 1
//`

var c1 = `
key: 
 sub_key1: 
  value0: 1
  value1: 1
  value2: 1
 sub_key2: 
  value0: 1
  value1: 1
 sub_key3: 1
`

//var c2 = `
//key:
// sub_key:
//  v0:
//   sub_value1: 1
//   sub_value2: 2
//`
var c2 = `
key: 
 sub_key1: 
  value0: 1
 sub_key2: 
  value0: 1
  value1: 
   name: ccc
   age: 18
 sub_key3: 2
`

//var s = `
//{
//	"name" : "cww"
//}
//`

type config struct {
	Key struct {
		SubKey1 struct {
			Value0 *int64 `yaml:"value0"`
			Value1 *int64 `yaml:"value1"`
			Value2 *int64 `yaml:"value2"`
		} `yaml:"sub_key1"`
		SubKey2 struct {
			Value0 *int64 `yaml:"value0"`
			Value1 *struct {
				Name *string `yaml:"name"`
				Age  *int64  `yaml:"age"`
			} `yaml:"value1"`
		} `yaml:"sub_key2"`
		SubKey3 *int64 `yaml:"sub_key3"`
	} `yaml:"key"`
}

func main() {
	c := new(config)
	yaml.NewDecoder(bytes.NewBufferString(c1)).Decode(c)
	fmt.Printf("%+12v\n", c)
	cc := new(config)
	yaml.NewDecoder(bytes.NewBufferString(c2)).Decode(cc)
	fmt.Printf("%+12v\n", cc)

	reset(c, cc)
	fmt.Printf("%+12v\n", c)
}

func reset(origin, append interface{}) {
	originValue := reflect.ValueOf(origin).Elem()
	appendValue := reflect.ValueOf(append).Elem()
	//fmt.Println(originValue.CanSet(), originValue.CanAddr(), appendValue.NumField())
	for i := 0; i < appendValue.NumField(); i++ {
		//originFiledValue := originValue.Field(i)
		appendFiledValue := appendValue.Field(i)
		//fmt.Println("零值", appendFiledValue.IsZero())
		if appendFiledValue.IsZero() { // 是零值没必要比了
			continue
		}
		// map和struct 才有子节点
		if appendFiledValue.Kind() == reflect.Map || appendFiledValue.Kind() == reflect.Struct {
			// 递归
			originI := originValue.Field(i).Addr().Interface()
			appendI := appendValue.Field(i).Addr().Interface()
			reset(originI, appendI)
		} else {
			// 进行替换
			originValue.Field(i).Set(appendFiledValue)
		}
	}
}
