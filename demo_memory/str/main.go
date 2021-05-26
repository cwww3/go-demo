package str

import (
	"strings"
)

var s0 string

func Solution1(s string) {
	var b strings.Builder
	// 设置足够大的cap,避免写过程中扩容 造成数据的不必要复制
	b.Grow(50)
	b.WriteString(s[:50])
	s0 = b.String()
}

func Solution2(s string) {
	s0 = strings.Repeat(s[:50], 1)
}

func createStringWithLengthOnHeap(n int) string {
	s := strings.Repeat("a", n)
	return s
}
