package slice

var s0 []int

func Solution1(s1 []int) {
	s := append([]int(nil), s1[:30]...)
	s0 = s
}
func Solution2(s1 []int) {
	s := make([]int, 30)
	copy(s, s1)
	s0 = s
}

func createSliceWithLength(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, 1)
	}
	return s
}
