package main

import (
	"fmt"
)

var m map[point]int

func main() {
	m = make(map[point]int)
	k := 2
	n := 7
	r := dp(k, n)
	fmt.Println(r)
}

type point struct {
	x int
	y int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dp(k, n int) int {
	// base case
	if k == 1 { // 只有一个鸡蛋每一层都得试
		return n
	}
	if n == 0 { // 一层楼都没有就不用试了
		return 0
	}
	p := point{k, n}
	result, ok := m[p]
	if ok {
		return result
	}
	result = n
	// 状态 k个鸡蛋 n层楼
	// 选择  哪层楼扔鸡蛋
	// P.S 这里不需要知道是第几层楼  而只需要知道楼层的数量
	// 二分法
	i := (n + 1) / 2
	result = min(result, max(dp(k-1, i-i), dp(k, n-i))+1)
	m[p] = result
	return result

}
