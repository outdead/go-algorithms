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
package heapsort

import (
	"github.com/Ganelon13/go-algorithms/utils/randomizer"
	"sort"
	"testing"
)

func TestSortInts(t *testing.T) {
	slice := randomizer.GenerateIntSlice(100, 999)
	SortInts(slice)
	if !sort.IsSorted(sort.IntSlice(slice)) {
		t.Fatalf("Slice is not sorted %v", slice)
	}
}

func TestSortIntsWithHeap(t *testing.T) {
	slice := randomizer.GenerateIntSlice(100, 999)
	SortIntsWithHeap(slice)
	if !sort.IsSorted(sort.IntSlice(slice)) {
		t.Fatalf("Slice is not sorted %v", slice)
	}
}

func TestSortInterface(t *testing.T) {
	slice := randomizer.GenerateIntSlice(100, 999)
	SortInterface(sort.IntSlice(slice))
	if !sort.IsSorted(sort.IntSlice(slice)) {
		t.Fatalf("Slice is not sorted %v", slice)
	}
}

func benchmarkSortInts(b *testing.B, size int) {
	b.StopTimer()
	initial := randomizer.GenerateIntSlice(size, 999)
	slice := make([]int, len(initial))
	b.StartTimer()

	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		copy(slice, initial)
		b.StartTimer()

		SortInts(slice)
	}
}

func benchmarkSortIntsWithHeap(b *testing.B, size int) {
	b.StopTimer()
	initial := randomizer.GenerateIntSlice(size, 999)
	slice := make([]int, len(initial))
	b.StartTimer()

	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		copy(slice, initial)
		b.StartTimer()

		SortIntsWithHeap(slice)
	}
}

func benchmarkSortInterface(b *testing.B, size int) {
	b.StopTimer()
	initial := randomizer.GenerateIntSlice(size, 999)
	slice := make(sort.IntSlice, len(initial))
	b.StartTimer()

	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		copy(slice, sort.IntSlice(initial))
		b.StartTimer()

		SortInterface(sort.IntSlice(slice))
	}
}

func BenchmarkSortInts100(b *testing.B) {
	benchmarkSortInts(b, 100)
}

func BenchmarkSortInts100000(b *testing.B) {
	benchmarkSortInts(b, 100000)
}

func BenchmarkSortIntsWithHeap100(b *testing.B) {
	benchmarkSortIntsWithHeap(b, 100)
}

func BenchmarkSortIntsWithHeap100000(b *testing.B) {
	benchmarkSortIntsWithHeap(b, 100000)
}

func BenchmarkSortInterface100(b *testing.B) {
	benchmarkSortInterface(b, 100)
}

func BenchmarkSortInterface100000(b *testing.B) {
	benchmarkSortInterface(b, 100000)
}
