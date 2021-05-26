package main

import "fmt"

func main() {
	var w = 10
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	nums := []int{2, 3, 2}

	dp := make([]int, w+1)

	//初始化 dp[0]=0

	for i := 0; i < len(weight); i++ {
		for j := w; j >= weight[i]; j-- {
			// 以上为01背包，然后加⼀个遍历个数
			for k := 1; k <= nums[i] && (j-k*weight[i]) >= 0; k++ { //遍历个数
				dp[j] = max(dp[j], dp[j-k*weight[i]]+k*value[i])
			}
		}
		fmt.Println(dp)
	}
	//fmt.Println(dp)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
