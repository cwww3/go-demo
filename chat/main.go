package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var upgrade = websocket.Upgrader{
	// 允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// 得到websocket连接
		connect, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			fmt.Printf("upgrade err=%v\n", err)
			w.Write([]byte("获取不到连接"))
		}
		myConnect := InitWebSocketConnect(connect)
		fmt.Printf("init connect \n")
		go myConnect.LoopRead()
		go myConnect.LoopWrite()

		for {
			data, err := myConnect.ReadMsg()
			if err != nil {
				fmt.Printf("ReadMsg fail err=%v\n", err)
				myConnect.Close("ReadMsg")
				break
			}
			err = myConnect.WriteMsg(data)
			if err != nil {
				fmt.Printf("WriteMsg fail err=%v\n", err)
				myConnect.Close("WriteMsg")
				break
			}
		}
	})
	http.ListenAndServe("0.0.0.0:8080", nil)
}

type MyConnect struct {
	connect *websocket.Conn
	in      chan []byte
	out     chan []byte
	isOn    bool
	mux     *sync.Mutex
	close   chan bool
}

func InitWebSocketConnect(conn *websocket.Conn) (myConnect *MyConnect) {
	return &MyConnect{
		connect: conn,
		out:     make(chan []byte, 10),
		in:      make(chan []byte, 10),
		mux:     new(sync.Mutex),
		isOn:    true,
		close:   make(chan bool, 1),
	}
}

func (con *MyConnect) ReadMsg() (data []byte, err error) {
	select {
	case data = <-con.in:
	case <-con.close:
		err = errors.New("read err")
	}
	return
}
func (con *MyConnect) WriteMsg(data []byte) (err error) {
	select {
	case con.out <- data:
	case <-con.close:
		err = errors.New("write err")
	}
	return
}
func (con *MyConnect) Close(name string) {
	fmt.Printf("关----------%s\n", name)
	//线程安全的，可重入的
	con.connect.Close()
	fmt.Printf("socket close\n")
	// 不能关闭已经关闭的chan
	con.mux.Lock()
	defer con.mux.Unlock()
	if con.isOn {
		close(con.close)
		fmt.Printf("chan close\n")
		con.isOn = false
	}
}

func (con *MyConnect) LoopRead() {
	defer con.Close("LoopRead")
	for {
		_, data, err := con.connect.ReadMessage()
		if err != nil {
			fmt.Printf("websocket read err=%v\n", err)
			break
		}
		select {
		case con.in <- data:
		case <-con.close:
			goto END
		}
	}
END:
}
func (con *MyConnect) LoopWrite() {
	defer con.Close("LoopWrite")
	for {
		select {
		case data := <-con.out:
			err := con.connect.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Printf("websocket write err=%v\n", err)
				break
			}
		case <-con.close:
			goto END
		}

	}
END:
}
