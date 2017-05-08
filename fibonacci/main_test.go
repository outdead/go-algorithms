package main

import (
	"testing"
	"runtime"
	"fmt"
)

const N = 1000000

var benchmarkNeedle = 10

func init() {
	runtime.GOMAXPROCS(1)
}

// *** tests

// Fibonacci Recursion
func testFibonacciRecursion(needle int, want int, t *testing.T) {
	f := FibonacciRecursion(needle)
	if f != want {
		t.Fatalf("Want %v, but got %v", want, f)
	}
	fmt.Printf("testFibonacciRecursion ........... OK [needle index = %v want number = %v]\n", needle, f)
}

func TestFibonacciRecursion1(t *testing.T) {
	testFibonacciRecursion(0, 0, t)
}

func TestFibonacciRecursion2(t *testing.T) {
	testFibonacciRecursion(1, 1, t)
}

func TestFibonacciRecursion3(t *testing.T) {
	testFibonacciRecursion(2, 1, t)
}

func TestFibonacciRecursion4(t *testing.T) {
	testFibonacciRecursion(10, 55, t)
}

func TestFibonacciRecursion5(t *testing.T) {
	testFibonacciRecursion(-10, -1, t)
	fmt.Println()
}

// Fibonacci Recursion Remember
func testFibonacciRecursionRemember(needle int, want int, t *testing.T) {
	var c []int
	if (needle >= 0) {
		c = make([]int, needle + 1)
	}

	f := FibonacciRecursionRemember(needle, c)
	if f != want {
		t.Fatalf("Want %v, but got %v", want, f)
	}
	fmt.Printf("testFibonacciRecursionRemember ... OK [needle index = %v want number = %v] %v\n", needle, f, c)
}

func TestFibonacciRecursionRemember1(t *testing.T) {
	testFibonacciRecursionRemember(0, 0, t)
}

func TestFibonacciRecursionRemember2(t *testing.T) {
	testFibonacciRecursionRemember(1, 1, t)
}

func TestFibonacciRecursionRemember3(t *testing.T) {
	testFibonacciRecursionRemember(2, 1, t)
}

func TestFibonacciRecursionRemember4(t *testing.T) {
	testFibonacciRecursionRemember(10, 55, t)
}

func TestFibonacciRecursionRemember5(t *testing.T) {
	testFibonacciRecursionRemember(-10, -1, t)
	fmt.Println()
}

// Fibonacci Closure. Go way
func testFibonacciClosure(needle int, want int, t *testing.T) {
	f := FibonacciClosure(needle)
	if f != want {
		t.Fatalf("Want %v, but got %v", want, f)
	}
	fmt.Printf("testFibonacciClosure ............. OK [needle index = %v want number = %v]\n", needle, f)
}

func TestFibonacciRecursionClosure1(t *testing.T) {
	testFibonacciClosure(0, 0, t)
}

func TestFibonacciRecursionClosure2(t *testing.T) {
	testFibonacciClosure(1, 1, t)
}

func TestFibonacciRecursionClosure3(t *testing.T) {
	testFibonacciClosure(2, 1, t)
}

func TestFibonacciRecursionClosure4(t *testing.T) {
	testFibonacciClosure(10, 55, t)
}

func TestFibonacciRecursionClosure5(t *testing.T) {
	testFibonacciClosure(-10, -1, t)
	fmt.Println()
}

// *** benchmarks

// Fibonacci Recursion
func BenchmarkFibonacciRecursion(b *testing.B) {
	b.N = N

	for i := 1; i <= b.N; i++ {
		_ = FibonacciRecursion(benchmarkNeedle)
	}
}

// Fibonacci Recursion with remember
func BenchmarkFibonacciRecursionRemember(b *testing.B) {
	b.N = N
	var c = make([]int, benchmarkNeedle + 1)

	for i := 1; i <= b.N; i++ {
		_ = FibonacciRecursionRemember(benchmarkNeedle, c)
	}
}

// Fibonacci Closure
func BenchmarkFibonacciClosure(b *testing.B) {
	b.N = N

	for i := 1; i <= b.N; i++ {
		_ = FibonacciClosure(benchmarkNeedle)
	}
}
