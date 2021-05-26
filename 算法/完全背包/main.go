package main

import "fmt"

func main() {
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
	// 正序
	for j := weight[0]; j <= w; j++ {
		dp[0][j] = j / weight[0] * value[0]
	}

	for i := 1; i < len(weight); i++ {
		for j := 0; j <= w; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	fmt.Println(dp[len(weight)-1][w])
	fmt.Println(dp)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func er() {
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
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				temp := j
				for temp >= weight[i] {
					dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
					temp -= weight[i]
				}
			}
		}
	}
	fmt.Println(dp[len(weight)-1][w])
	fmt.Println(dp)
}

func yi() {
	var w = 4
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}

	dp := make([]int, w+1)

	//初始化 dp[0]=0

	for i := 0; i < len(weight); i++ {
		for j := 0; j <= w; j++ {
			if j >= weight[i] {
				dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			}
		}
	}
	fmt.Println(dp[w])
}
