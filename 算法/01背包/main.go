package main

import "fmt"

func main() {
	yi01()
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func er01() {
	var w = 4
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}

	dp := make([][]int, len(weight))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, w+1)
	}

	//初始化
	for i := 0; i < len(weight); i++ {
		dp[i][0] = 0
	}
	for j := w; j >= weight[0]; j-- {
		dp[0][j] = dp[0][j-weight[0]] + value[0]
	}

	for i := 1; i < len(weight); i++ {
		for j := 0; j <= w; j++ {
			if weight[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	fmt.Println(dp[len(weight)-1][w])
}

func yi01() {
	var w = 4
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}

	dp := make([]int, w+1)

	//初始化 dp[0]=0

	for i := 0; i < len(weight); i++ {
		for j := w; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	fmt.Println(dp)
}

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 如果是奇数 可定不会相等
	if sum%2 == 1 {
		return false
	}

	w := sum / 2
	dp := make([]int, w+1)

	for i := 0; i < len(nums); i++ {
		for j := w; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	if dp[w] == w {
		return true
	}
	return false
}

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, num := range stones {
		sum += num
	}

	w := sum / 2
	dp := make([]int, w+1)

	for i := 0; i < len(stones); i++ {
		for j := w; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	// dp[w] <= w
	return sum - dp[w] - dp[w]
}
