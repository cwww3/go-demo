package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file,_ := os.Open("./file/my-avatar.png")
	defer file.Close()
	newFile,_ := os.Create("./file/avatar.png")
	defer newFile.Close()
	fmt.Println(file.Seek(0,io.SeekCurrent))
	fmt.Println(newFile.Seek(0,io.SeekCurrent))
	io.Copy(newFile,file)
	fmt.Println("after")
	fmt.Println(file.Seek(0,io.SeekCurrent))
	fmt.Println(newFile.Seek(0,io.SeekCurrent))
	fmt.Println("again")
	newFile2,_ := os.Create("./file/avatar2.png")
	defer newFile2.Close()
	fmt.Println(io.Copy(newFile2,file))
	fmt.Println("reset")
	fmt.Println(file.Seek(0,io.SeekStart))
	fmt.Println(newFile.Seek(0,io.SeekStart))
}
