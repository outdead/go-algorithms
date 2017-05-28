//
// Package heapsort contains sorting methods based on a binary heap.
//
// Heapsort is a comparison-based sorting algorithm. Heapsort can be
// thought of as an improved selection sort: like that algorithm, it
// divides its input into a sorted and an unsorted region, and it
// iteratively shrinks the unsorted region by extracting the
// largest/smallest element and moving that to the sorted region.
// The improvement consists of the use of a heap data structure rather
// than a linear-time search to find the maximum
//
// Data structure              Array
// Best-case performance       O(n log n)
// Average performance         O(n log n)
// Worst-case performance      O(n log n)
// Worst-case space complexity O(1) auxiliary
//
package heapsort

import (
	"github.com/Ganelon13/go-algorithms/utils/randomizer"
	"sort"
	"testing"
)

func TestIntsWithHeap(t *testing.T) {
	slice := randomizer.GenerateIntSlice(100, 999)
	IntsWithHeap(slice)
	if !sort.IsSorted(sort.IntSlice(slice)) {
		t.Fatalf("Slice is not sorted %v", slice)
	}
}

func TestInts(t *testing.T) {
	slice := randomizer.GenerateIntSlice(100, 999)
	Ints(slice)
	if !sort.IsSorted(sort.IntSlice(slice)) {
		t.Fatalf("Slice is not sorted %v", slice)
	}
}

func TestSort(t *testing.T) {
	slice := randomizer.GenerateIntSlice(100, 999)
	Sort(sort.IntSlice(slice))
	if !sort.IsSorted(sort.IntSlice(slice)) {
		t.Fatalf("Slice is not sorted %v", slice)
	}
}

func benchmarkIntsWithHeap(b *testing.B, size int) {
	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		slice := randomizer.GenerateIntSlice(size, 999)
		b.StartTimer()

		IntsWithHeap(slice)
	}
}

func benchmarkInts(b *testing.B, size int) {
	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		slice := randomizer.GenerateIntSlice(size, 999)
		b.StartTimer()

		Ints(slice)
	}
}

func benchmarkSort(b *testing.B, size int) {
	b.StopTimer()
	initial := randomizer.GenerateIntSlice(size, 999)
	slice := make(sort.IntSlice, len(initial))
	b.StartTimer()

	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		copy(slice, sort.IntSlice(initial))
		b.StartTimer()

		Sort(sort.IntSlice(slice))
	}
}

func BenchmarkIntsWithHeap100(b *testing.B) {
	benchmarkIntsWithHeap(b, 100)
}

func BenchmarkIntsWithHeap1000(b *testing.B) {
	benchmarkIntsWithHeap(b, 1000)
}

func BenchmarkIntsWithHeap10000(b *testing.B) {
	benchmarkIntsWithHeap(b, 10000)
}

func BenchmarkIntsWithHeap100000(b *testing.B) {
	benchmarkIntsWithHeap(b, 100000)
}

func BenchmarkInts100(b *testing.B) {
	benchmarkInts(b, 100)
}

func BenchmarkInts1000(b *testing.B) {
	benchmarkInts(b, 1000)
}

func BenchmarkInts10000(b *testing.B) {
	benchmarkInts(b, 10000)
}

func BenchmarkInts100000(b *testing.B) {
	benchmarkInts(b, 100000)
}

func BenchmarkSort100(b *testing.B) {
	benchmarkSort(b, 100)
}

func BenchmarkSort1000(b *testing.B) {
	benchmarkSort(b, 1000)
}

func BenchmarkSort10000(b *testing.B) {
	benchmarkSort(b, 10000)
}

func BenchmarkSort100000(b *testing.B) {
	benchmarkSort(b, 100000)
}
