package main

import "fmt"

func main() {
	needle := 10

	fmt.Printf("FibonacciRecursion .............. %v\n", FibonacciRecursion(needle))

	var c []int
	if (needle >= 0) {
		c = make([]int, needle + 1)
	}
	fmt.Printf("FibonacciRecursionWithRemember .. %v %v\n", FibonacciRecursionRemember(needle, c), c)

	fmt.Printf("FibonacciClosure ................ %v\n", FibonacciClosure(needle))
}

// https://youtu.be/eqWzZGNO_XM?list=PLrCZzMib1e9pDxHYzmEzMmnMMUK-dz0_7
func FibonacciRecursion(n int) int {
	switch {
	case n < 0:
		return -1
	case n == 0:
		return 0
	case n == 1:
		return 1
	}

	return FibonacciRecursion(n - 1) + FibonacciRecursion(n - 2)
}

// https://youtu.be/eqWzZGNO_XM?list=PLrCZzMib1e9pDxHYzmEzMmnMMUK-dz0_7
func FibonacciRecursionRemember(n int, c []int) int {
	switch {
	case n < 0:
		return -1
	case n == 0:
		return 0
	case n == 1:
		c[1] = 1
		return 1
	}

	if c[n] > 0 {
		return c[n]
	}

	c[n] = FibonacciRecursionRemember(n - 1, c) + FibonacciRecursionRemember(n - 2, c)

	return c[n]
}

// https://golang.org/doc/play/fib.go
func FibonacciClosure(n int) int {
	if n < 0 {
		return -1
	}

	f := fibonacciClosure()
	var res int

	for i := 0; i < n; i++ {
		res = f()
	}

	return res
}

func fibonacciClosure() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}
