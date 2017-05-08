package main

import (
	"testing"
	"sort"
	"runtime"
	"fmt"
)

const N = 10000000

var (
	initial = []int{1, 2, 5, 5, 12, 13, 13, 17, 21, 21, 24, 25, 25, 44, 45, 47, 56, 66, 75, 84, 85, 92, 97, 99, 99}
	hashMap map[int]int

	firstNeedle, firstWant = 1, 0
	centerNeedle, centerWant = 44, 13
	lastNeedle, lastWant = 99, 23
)

func init() {
	runtime.GOMAXPROCS(1)

	hashMap = make(map[int]int, len(initial))
	for key, value := range initial {
		hashMap[value] = key
	}
}

// *** tests

// Line Search
func testLineSearch(needle int, want int, t *testing.T) {
	index := LineSearch(initial, needle)
	if index != want {
		t.Fatalf("Want %v, but got %v", want, index)
	}
	fmt.Printf("testLineSearch ............. OK [needle value = %v, want index = %v]\n", needle, want)
}

func TestLineSearchFirst(t *testing.T) {
	testLineSearch(firstNeedle, firstWant, t)
}

func TestLineSearchCenter(t *testing.T) {
	testLineSearch(centerNeedle, centerWant, t)
}

func TestLineSearchLast(t *testing.T) {
	testLineSearch(lastNeedle, lastWant, t)
	fmt.Println()
}

// Line Predicate Search
func testLinePredicateSearch(needle int, want int, t *testing.T) {
	index := LineSearch(initial, needle)
	if index != want {
		t.Fatalf("Want %v, but got %v", want, index)
	}
	fmt.Printf("testLinePredicateSearch .... OK [needle value = %v, want index = %v]\n", needle, want)
}

func TestLinePredicateSearchFirst(t *testing.T) {
	testLinePredicateSearch(firstNeedle, firstWant, t)
}

func TestLinePredicateSearchCenter(t *testing.T) {
	testLinePredicateSearch(centerNeedle, centerWant, t)
}

func TestLinePredicateSearchLast(t *testing.T) {
	testLinePredicateSearch(lastNeedle, lastWant, t)
	fmt.Println()
}

// Binary Recursion Search
func testBinarySearch(needle int, want int, t *testing.T) {
	index := BinarySearch(initial, needle)
	if index != want {
		t.Fatalf("Want %v, but got %v", want, index)
	}
	fmt.Printf("testBinarySearch ........... OK [needle value = %v, want index = %v]\n", needle, want)
}

func TestBinarySearchFirst(t *testing.T) {
	testBinarySearch(firstNeedle, firstWant, t)
}

func TestBinarySearchCenter(t *testing.T) {
	testBinarySearch(centerNeedle, centerWant, t)
}

func TestBinarySearchLast(t *testing.T) {
	testBinarySearch(lastNeedle, lastWant, t)
	fmt.Println()
}

// Binary Search
func testBinaryRecursionSearch(needle int, want int, t *testing.T) {
	index := BinarySearch(initial, needle)
	if index != want {
		t.Fatalf("Want %v, but got %v", want, index)
	}
	fmt.Printf("testBinaryRecursionSearch .. OK [needle value = %v, want index = %v]\n", needle, want)
}

func TestBinaryRecursionSearchFirst(t *testing.T) {
	testBinaryRecursionSearch(firstNeedle, firstWant, t)
}

func TestBinaryRecursionSearchCenter(t *testing.T) {
	testBinaryRecursionSearch(centerNeedle, centerWant, t)
}

func TestBinaryRecursionSearchLast(t *testing.T) {
	testBinaryRecursionSearch(lastNeedle, lastWant, t)
	fmt.Println()
}

// Interpolation Search
func testInterpolationSearch(needle int, want int, t *testing.T) {
	index := LineSearch(initial, needle)
	if index != want {
		t.Fatalf("Want %v, but got %v", want, index)
	}
	fmt.Printf("testInterpolationSearch .... OK [needle value = %v, want index = %v]\n", needle, want)
}

func TestInterpolationeSearchFirst(t *testing.T) {
	testInterpolationSearch(firstNeedle, firstWant, t)
}

func TestInterpolationSearchCenter(t *testing.T) {
	testInterpolationSearch(centerNeedle, centerWant, t)
}

func TestInterpolationSearchLast(t *testing.T) {
	testInterpolationSearch(lastNeedle, lastWant, t)
	fmt.Println()
}

// *** benchmarks

// SearchInts from sort package
func benchmarkNativeSearchInts(needle int, b *testing.B) {
	b.N = N

	for i := 1; i <= b.N; i++ {
		_ = sort.SearchInts(initial, needle)
	}
}

func BenchmarkNativeSearchIntsFirst(b *testing.B) {
	benchmarkNativeSearchInts(firstNeedle, b)
}

func BenchmarkNativeSearchIntsCenter(b *testing.B) {
	benchmarkNativeSearchInts(centerNeedle, b)
}

func BenchmarkNativeSearchIntsLast(b *testing.B) {
	benchmarkNativeSearchInts(lastNeedle, b)
}

// Line Search
func benchmarkLineSearch(needle int, b *testing.B) {
	b.N = N

	for i := 1; i <= b.N; i++ {
		_ = LineSearch(initial, needle)
	}
}

func BenchmarkLineSearchFirst(b *testing.B) {
	benchmarkLineSearch(firstNeedle, b)
}

func BenchmarkLineSearchCenter(b *testing.B) {
	benchmarkLineSearch(centerNeedle, b)
}

func BenchmarkLineSearchLast(b *testing.B) {
	benchmarkLineSearch(lastNeedle, b)
}

// Line Search with predicate function. Go way. Wy it is more slowly then simple Line Search?
func benchmarkLinePredicateSearch(needle int, b *testing.B) {
	b.N = N

	for i := 1; i <= b.N; i++ {
		_ = LinePredicateSearch(initial, needle)
	}
}

func BenchmarkLinePredicateSearchFirst(b *testing.B) {
	benchmarkLinePredicateSearch(firstNeedle, b)
}

func BenchmarkLinePredicateSearchCenter(b *testing.B) {
	benchmarkLinePredicateSearch(centerNeedle, b)
}

func BenchmarkLinePredicateSearchLast(b *testing.B) {
	benchmarkLinePredicateSearch(lastNeedle, b)
}

// Binary Search
func benchmarkBinarySearch(needle int, b *testing.B) {
	b.N = N
	for i := 1; i <= b.N; i++ {
		_ = BinarySearch(initial, needle)
	}
}

func BenchmarkBinarySearchFirst(b *testing.B) {
	benchmarkBinarySearch(firstNeedle, b)
}

func BenchmarkBinarySearchCenter(b *testing.B) {
	benchmarkBinarySearch(centerNeedle, b)
}

func BenchmarkBinarySearchLast(b *testing.B) {
	benchmarkBinarySearch(lastNeedle, b)
}

// Binary Recursion Search
func benchmarkBinaryRecursionSearch(needle int, b *testing.B) {
	b.N = N
	for i := 1; i <= b.N; i++ {
		_ = BinarySearch(initial, needle)
	}
}

func BenchmarkBinaryRecursionSearchFirst(b *testing.B) {
	benchmarkBinaryRecursionSearch(firstNeedle, b)
}

func BenchmarkBinaryRecursionSearchCenter(b *testing.B) {
	benchmarkBinaryRecursionSearch(centerNeedle, b)
}

func BenchmarkBinaryRecursionSearchLast(b *testing.B) {
	benchmarkBinaryRecursionSearch(lastNeedle, b)
}

// Interpolation Search
func benchmarkInterpolationSearch(needle int, b *testing.B) {
	b.N = N
	for i := 1; i <= b.N; i++ {
		_ = InterpolationSearch(initial, needle)
	}
}

func BenchmarkInterpolationSearchFirst(b *testing.B) {
	benchmarkInterpolationSearch(firstNeedle, b)
}

func BenchmarkInterpolationSearchCenter(b *testing.B) {
	benchmarkInterpolationSearch(centerNeedle, b)
}

func BenchmarkInterpolationSearchLast(b *testing.B) {
	benchmarkInterpolationSearch(lastNeedle, b)
}

// Map Search
func benchmarkMapSearch(needle int, b *testing.B) {
	b.N = N

	for i := 1; i <= b.N; i++ {
		_, _ = hashMap[needle];
	}
}

func BenchmarkMapSearchFirst(b *testing.B) {
	benchmarkMapSearch(firstNeedle, b)
}

func BenchmarkMapSearchCenter(b *testing.B) {
	benchmarkMapSearch(centerNeedle, b)
}

func BenchmarkMapSearchLast(b *testing.B) {
	benchmarkMapSearch(lastNeedle, b)
}
