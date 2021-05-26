package main

func main() {

}

func search(nums []int, target int) int {
	// 定义 target在区间[left,right]里
	left := 0
	right := len(nums) - 1
	// 根据定义 使用<=
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func search2(nums []int, target int) int {
	// 定义 target在区间[left,right)里
	left := 0
	right := len(nums)
	// 根据定义 使用<
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			// 有边界取不到 所以用=
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}
