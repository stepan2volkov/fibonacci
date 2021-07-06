package recursion

import (
	"testing"
)

func BenchmarkRecursionFib30(t *testing.B) {
	fib := NewRecursion()
	fib.Fib(30)
}
