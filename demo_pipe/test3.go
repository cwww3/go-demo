package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	a, err := capture(func() {
		fmt.Print(strings.Repeat("1\n", 4<<10+1))
	})
	fmt.Println(a)
	fmt.Println(err)

}

func capture(f func()) (string, error) {
	rescueOut := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return "", fmt.Errorf("capture: %s", err.Error())
	}
	os.Stdout = w
	go func() {
		time.Sleep(3 * time.Second)
		f()
		w.Close()
	}()
	out, err := ioutil.ReadAll(r)

	if err != nil {
		log.Fatalf("capture: %s", err.Error())
	}
	os.Stdout = rescueOut
	return string(out), nil
}
