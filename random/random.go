package random

import (
	"math/rand"
	"time"

	"github.com/qencept/fibonacci/web"
)

type RandomStruct struct {
	fibers []web.Fiber
}

func NewRandomStruct(fibers []web.Fiber) *RandomStruct {
	rand.Seed(time.Now().UnixNano())
	return &RandomStruct{fibers}
}

func (rs *RandomStruct) Fib(n int) int {
	fiber := rs.fibers[rand.Intn(len(rs.fibers))]
	return fiber.Fib(n)
}

func (rs *RandomStruct) Name() string {
	fiber := rs.fibers[rand.Intn(len(rs.fibers))]
	return fiber.Name()
}
