//
// Package bubblesort contains realisation of bubble sort algorithm.
//
// Bubble sort has worst-case and average complexity both О(n^2),
// where n is the number of items being sorted.
//
// Data structure              Array
// Best-case performance       O(n)
// Average performance         O(n^2)
// Worst-case performance      O(n^2)
// Worst-case space complexity O(1) auxiliary
//
package bubblesort

import (
	"github.com/ganelon13/go-algorithms/utils/randomizer"
	"sort"
	"testing"
)

func TestInts(t *testing.T) {
	data := randomizer.GenerateIntSlice(100, 999)
	Ints(data)
	if !sort.IsSorted(sort.IntSlice(data)) {
		t.Fatalf("Data is not sorted %v", data)
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
