//	One way Fibonacci sequence generator
package model

type fibonacciGenerator struct {
	x uint64
	y uint64
}

func NewFibonacciGenerator() *fibonacciGenerator {
	return &fibonacciGenerator{x: 0, y: 1}
}

func (gen *fibonacciGenerator) Next() uint64 {
	value := gen.x	// Store previous value
	gen.x, gen.y = gen.y, gen.x+gen.y
	return value
}

func (gen *fibonacciGenerator) Current() uint64 {
	return gen.x
}