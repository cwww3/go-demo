package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLazySingleton(t *testing.T) {
	assert.Equal(t, GetLazySingleton(), GetLazySingleton())
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazySingleton() != GetLazySingleton() {
				b.Errorf("test fail")
			}
		}
	})
}
