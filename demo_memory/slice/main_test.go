package slice

import "testing"

func BenchmarkSolution1(b *testing.B) {
	s := createSliceWithLength(1 << 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Solution1(s)
	}
}

func BenchmarkSolution2(b *testing.B) {
	s := createSliceWithLength(1 << 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Solution2(s)
	}
}

func TestSolution2(t *testing.T) {
	s := createSliceWithLength(1 << 10)
	Solution2(s)
}
