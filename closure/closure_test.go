package closure

import (
	"testing"
)

func BenchmarkClosureFib30(t *testing.B) {
	fib := NewClosure()
	fib.Fib(30)
}
