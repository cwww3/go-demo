package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement2(nums, val))
}

func removeElement(nums []int, val int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		if nums[i] == val {
			nums[j], nums[i] = nums[i], nums[j]
			j--
		} else {
			i++
		}
	}
	return i
}

func removeElement2(nums []int, val int) int {
	// 定义元素在[left,right]中
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		// 相等只有快指针往前移动
		// 不相等都移动 且进行赋值
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
