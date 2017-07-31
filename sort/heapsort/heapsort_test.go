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
	"github.com/ganelon13/go-algorithms/utils/randomizer"
	"sort"
	"testing"
)

func TestIntsWithHeap(t *testing.T) {
	data := randomizer.GenerateIntSlice(100, 999)
	IntsWithHeap(data)
	if !sort.IsSorted(sort.IntSlice(data)) {
		t.Fatalf("Data is not sorted %v", data)
	}
}

func TestInts(t *testing.T) {
	data := randomizer.GenerateIntSlice(100, 999)
	Ints(data)
	if !sort.IsSorted(sort.IntSlice(data)) {
		t.Fatalf("Data is not sorted %v", data)
	}
}

func TestSort(t *testing.T) {
	data := randomizer.GenerateIntSlice(100, 999)
	Sort(sort.IntSlice(data))
	if !sort.IsSorted(sort.IntSlice(data)) {
		t.Fatalf("Data is not sorted %v", data)
	}
}

func benchmarkIntsWithHeap(b *testing.B, size int) {
	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		data := randomizer.GenerateIntSlice(size, 999)
		b.StartTimer()

		IntsWithHeap(data)
	}
}

func benchmarkInts(b *testing.B, size int) {
	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		data := randomizer.GenerateIntSlice(size, 999)
		b.StartTimer()

		Ints(data)
	}
}

func benchmarkSort(b *testing.B, size int) {
	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		data := randomizer.GenerateIntSlice(size, 999)
		b.StartTimer()

		Sort(sort.IntSlice(data))
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
