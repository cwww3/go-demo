package main

import (
	"fmt"
	"time"
)

func main() {

	timeA := time.Now()
	//fmt.Println(timeA.UnixNano() > 1e9)
	timeB := time.Unix(1, 1)

	fmt.Println(timeA) // 2018-10-26 16:37:02.216165074 +0800 CST m=+0.000363156
	fmt.Println(timeB) // 2018-10-26 16:37:02.216165074 +0800 CST

	format := "20060102150405"
	// t1 没有写 time.Now() 是为了避免秒以下单位的时间的影响
	// 除此之外和写 time.Now() 是一样的
	t1 := time.Date(2017, time.November, 30, 0, 0, 0, 0, time.Local)

	// t1 使用本地时区进行 format，结果是 "20171130000000"
	// 由进行 parse，由于没有指定时区，结果是 UTC 时间 2017/11/30 00:00:00
	t2, _ := time.Parse(format, t1.Format(format))
	// t1 使用本地时区进行 format，结果是 "20171130000000"
	// t2 使用 UTC 时间进行 format，结果是 "20171130000000"
	// 所以输出 true
	println("1-1 ", t1.Format(format) == t2.Format(format))

	// 很显然不相等，既不是指同一个时间点，时区信息也不一样，所以输出 false
	println("1-2 ", t1 == t2)

	// 显然不相等，t1 和 t2 不是指同一个时间点，所以输出 false
	println("1-3 ", t1.Equal(t2))

	// t1 使用本地时区进行 format，结果是 "20171130000000"
	// 由进行 parse，指定了本地时区，结果是本地时间 2017/11/30 00:00:00
	t2, _ = time.ParseInLocation(format, t1.Format(format), time.Local)

	// 显然相等，输出 true
	println("2-1 ", t1.Format(format) == t2.Format(format))
	// 既指同一个时间点，时区信息也一样，输出 true
	println("2-2 ", t1 == t2)
	// 显然相等，输出 true
	println("2-3 ", t1.Equal(t2))

	// 原本 t2 与 t1 完全相等，现在将 t2 改为 UTC 时间
	t2 = t2.UTC()
	// t1 使用本地时区进行 format，结果是 "20171130000000"
	// t2 使用 UTC 时间进行 format，结果是 "20171129160000"
	// 所以输出 false
	println("3-1 ", t1.Format(format) == t2.Format(format))

	// t1 和 t2 表示了相同的时间点，但各自时区信息不同，所以输出 false
	println("3-2 ", t1 == t2)

	// 由于 t1 和 t2 表示了相同的时间点，所以输出 true
	println("3-3 ", t1.Equal(t2))
}
