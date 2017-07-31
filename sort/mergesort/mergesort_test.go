//
// Package mergesort contains 2 variants of merge sort methods -  topDown
// implementation and bottomUp implementation. BottomUp has a slightly higher
// constant productivity in relation to the topDown implementation.
//
// Data structure              Array
// Best-case performance       O(n log n) typical, O(n) natural variant
// Average performance         O(n log n)
// Worst-case performance      O(n log n)
// Worst-case space complexity O(n) auxiliary
//
package mergesort

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

		b.StopTimer()
		if !sort.IsSorted(sort.IntSlice(data)) {
			b.Fatal("Data is not sorted")
		}
		b.StartTimer()
	}
}

func benchmarkInts2(b *testing.B, size int) {
	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		data := randomizer.GenerateIntSlice(size, 999)
		b.StartTimer()

		Ints2(data)

		b.StopTimer()
		if !sort.IsSorted(sort.IntSlice(data)) {
			b.Fatal("Data is not sorted")
		}
		b.StartTimer()
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

func BenchmarkInts2_100(b *testing.B) {
	benchmarkInts2(b, 100)
}

func BenchmarkInts2_1000(b *testing.B) {
	benchmarkInts2(b, 1000)
}

func BenchmarkInts2_10000(b *testing.B) {
	benchmarkInts2(b, 10000)
}

func BenchmarkInts2_100000(b *testing.B) {
	benchmarkInts2(b, 100000)
}
