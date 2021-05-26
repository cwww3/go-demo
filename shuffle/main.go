package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := make([]int, 100)
	for i := 0; i < 100; i++ {
		nums[i] = i + 1
	}
	fmt.Println(shuffle(nums))
}

func shuffle(nums []int) []int {
	cnt := len(nums) - 1
	// 100个数交换了99次
	for cnt > 0 {
		//随机出来的值在要交换的值得前面
		randInt := rand.Intn(cnt)
		nums[cnt], nums[randInt] = nums[randInt], nums[cnt]
		cnt--
	}
	return nums
}
