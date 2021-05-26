package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	//// 生成秘钥对
	//privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	//if err != nil {
	//	panic(err)
	//}
	//// 公钥
	//publicKey := privateKey.PublicKey
	//msg := []byte("重要信息")
	//sha := sha256.New()
	//// 加密信息
	//encryptedMsg, err := rsa.EncryptOAEP(sha, rand.Reader, &publicKey, msg, nil)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(encryptedMsg)
	//// 解密信息
	//decryptedMsg, err := privateKey.Decrypt(nil, encryptedMsg, &rsa.OAEPOptions{Hash: crypto.SHA256})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(decryptedMsg))
	ch := make(chan string, 100)

	go func() {
		for {
			ch <- "asong"
		}
	}()
	go func() {
		// 开启pprof，监听请求
		ip := "127.0.0.1:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
		}
	}()
	// 创建一个
	t := time.NewTimer(time.Minute * 3)
	//s := time.NewTicker(time.Minute * 3)
	defer t.Stop()
	time.Time{}
	//defer s.Stop()
	for {
		t.Reset(time.Minute * 3)
		select {
		case <-ch:
			//case <-s.C:
		case <-t.C:
		}
	}
}
