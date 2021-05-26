package main

import "fmt"

func main() {
	str := "中国"
	for i := 0; i < len(str); i++ {
		fmt.Println(i, string(str[i]))
	}
	fmt.Println("--------")
	for i, v := range str {
		fmt.Println(i, string(v))
	}

	fmt.Println("=====控制输出的宽度和精度====")
	fmt.Printf("|%5d|%5d|\n", 12, 345)
	fmt.Println("=====输出宽度,同时指定浮点数====")
	fmt.Printf("|%5.2f|%5.2f|\n", 1.2, 3.45)
	fmt.Println("=====左对齐====")
	fmt.Printf("|%-5.2f|%-5.2f|\n", 1.2, 3.45)
}
