package str

import "testing"

func BenchmarkSolution1(b *testing.B) {
	s := createStringWithLengthOnHeap(1 << 20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Solution2(s)
	}
}

func BenchmarkSolution2(b *testing.B) {
	s := createStringWithLengthOnHeap(1 << 20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Solution2(s)
	}
}
