package cache

import (
	"testing"
)

func BenchmarkCacheFib30(t *testing.B) {
	fib := NewCache()
	fib.Fib(30)
}
