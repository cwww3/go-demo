package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	s := sha1.New()
	s.Write([]byte("wlOU1c3CIpJiLXqWvEFTDQ=="))
	s.Write([]byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	data := s.Sum(nil)

	fmt.Println(base64.StdEncoding.EncodeToString(data))
}
