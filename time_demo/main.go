package main

import (
	"fmt"
	"time"
)

var timezoneOffset int

const weekSecond = 7 * 86400

var m map[string]int64

func main() {
	fmt.Println(time.Now().Format("2006-01-02"))
	t := time.Unix(WeekAgo(0), 0)
	fmt.Println(t.Format("2006-01-02"))

}
func init() {
	_, timezoneOffset = time.Now().Zone()
}
func WeekAgo(week int) int64 {
	now := time.Now()
	offset := int64(timezoneOffset) - 4*86400
	return (now.Unix()+offset)/weekSecond*weekSecond - offset - int64(week)*weekSecond

}
