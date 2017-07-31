//
// Package quicksort contains implementation of basic quicksort algorithm
// without improvements. Quicksort has a reputation as the fastest sort.
// Optimized variants of quicksort are common features of many languages
// and libraries. One often contrasts quicksort with merge sort, because
// both sorts have an average time of O(n log n)
//
// Data structure              Array
// Best-case performance       О(n log n)
// Average performance         О(n log n)
// Worst-case performance      О(n^2)
// Worst-case space complexity O(1) auxiliary
//
package quicksort

import (
	"github.com/ganelon13/go-algorithms/utils/randomizer"
	"sort"
	"testing"
)

func TestInts(t *testing.T) {
	data := randomizer.GenerateIntSlice(100, 2147483647)
	data[13] = -2147483648
	Ints(data)
	if !sort.IntsAreSorted(data) {
		t.Fatalf("Data is not sorted %v", data)
	}
}

func benchmarkInts_sortPackage(b *testing.B, size int) {
	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		data := randomizer.GenerateIntSlice(size, 999)
		b.StartTimer()

		sort.Ints(data)

		b.StopTimer()
		if !sort.IntsAreSorted(data) {
			b.Fatalf("Data is not sorted %v", data)
		}
		b.StartTimer()
	}
}

func benchmarkInts(b *testing.B, size int) {
	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		data := randomizer.GenerateIntSlice(size, 999)
		//data := randomizer.GenerateSortedIntSlice(size)
		// I did not include it in the benchmarks, but you can do it yourself.
		// `left + (right-left)/2` as pivotIndex makes this case faster
		// BenchmarkInts100-4                        500000              2780 ns/op
		// BenchmarkInts1000-4                        50000             23321 ns/op
		// BenchmarkInts10000-4                       10000            300517 ns/op
		// BenchmarkInts100000-4                        500           3502202 ns/op
		b.StartTimer()

		Ints(data)

		b.StopTimer()
		if !sort.IntsAreSorted(data) {
			b.Fatalf("Data is not sorted %v", data)
		}
		b.StartTimer()
	}
}

func BenchmarkInts_sortPackage100(b *testing.B) {
	benchmarkInts_sortPackage(b, 100)
}

func BenchmarkInts_sortPackage1000(b *testing.B) {
	benchmarkInts_sortPackage(b, 1000)
}

func BenchmarkInts_sortPackage10000(b *testing.B) {
	benchmarkInts_sortPackage(b, 10000)
}

func BenchmarkInts_sortPackage100000(b *testing.B) {
	benchmarkInts_sortPackage(b, 100000)
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
