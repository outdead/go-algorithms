//
// Package radixsort provides sort operation for []int type with LSD radix sort algorithm.
//
// Numbers are sorted by discharges. There are two options: least significant digit (LSD)
// and the most significant digit (MSD). With LSD sorting, low-order digits are first sorted,
// then older ones. With MSD sorting, everything is vice versa. With LSD sorting, the order
// is as follows: short keys go before long, keys of the same size are sorted alphabetically,
// this coincides with the normal representation of numbers 1, 2, 3, 4, 5, 6, 7, 8, 9, 10.
// With MSD sort an alphabetical order is obtained that is suitable for sorting strings
//
// Data structure              Array
// Best-case performance       O(n)
// Average performance         O(n * k)
// Worst-case performance      O(n * k)
// Worst-case space complexity O(n + k) auxiliary
//
package radixsort

import (
	"github.com/ganelon13/go-algorithms/utils/randomizer"
	"sort"
	"testing"
)

func TestInts(t *testing.T) {
	data := randomizer.GenerateIntSlice(100, 2147483647)
	data[1] = -11
	data[37] = -30000
	data[7] = -2147483648
	Ints(data)
	if !sort.IntsAreSorted(data) {
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
		if !sort.IntsAreSorted(data) {
			b.Fatalf("Data is not sorted %v", data)
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
