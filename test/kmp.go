package main

func main() {
	//s := "111"
	//aa := len(s)
	//const length = aa
	//table := [length][length]int{}
	//fmt.Println(table)
}

func strStr(haystack string, needle string) int {
	next := getArr(needle)
	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i++
			j = j - next[j]
		}
	}
	if j == len(needle) {
		return i - len(needle) + 1
	}
	return -1
}

func getArr(sub string) []int {
	next := make([]int, len(sub))
	next[0] = -1
	i := 1
	j := 0
	for i < len(sub) {
		cnt := 0
		for sub[i] == sub[j] && i < len(sub)-1 {
			i++
			j++
			cnt++
			next[i+1] = cnt
		}
		cnt = 0
		next[i] = 0
		i++
	}
	return next
}
