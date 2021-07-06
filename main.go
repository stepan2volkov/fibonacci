package main

import (
	"fmt"

	"github.com/qencept/fibonacci/cache"
	"github.com/qencept/fibonacci/closure"
	"github.com/qencept/fibonacci/random"
	"github.com/qencept/fibonacci/recursion"
	"github.com/qencept/fibonacci/web"
)

func main() {
	fibImplementations := []web.Fiber{recursion.NewRecursion(), cache.NewCache(), closure.NewClosure()}
	fiber := random.NewRandomStruct(fibImplementations)
	fmt.Println("Starting Fibonacci (Random)")
	web.Serve(fiber)
}
