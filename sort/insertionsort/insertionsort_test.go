//
// Package insertionsort provides sort operation for []int type
//
// Data structure	           Array
// Best-case performance       O(n) comparisons, O(1) swaps
// Average performance	       О(n^2) comparisons, swaps
// Worst-case performance      О(n^2) comparisons, swaps
// Worst-case space complexity O(1) auxiliary
//
package insertionsort

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
