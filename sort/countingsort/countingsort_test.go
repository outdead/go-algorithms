//
// Package countingsort provides sort operation for []int type
//
// Counting sort is a sorting algorithm that uses the range of the numbers
// of the array being sorted to count the matching elements. The use of
// counting sort is advisable only if the numbers have a range of possible
// values that is small enough, for example, a million natural numbers less
// than 1000.
//
// Data structure              Array
// Best-case performance       O(n + k)
// Average performance         O(n + k)
// Worst-case performance      O(n + k)
// Worst-case space complexity O(k) auxiliary
//
package countingsort

import (
	"github.com/ganelon13/go-algorithms/utils/randomizer"
	"sort"
	"testing"
)

func TestInts(t *testing.T) {
	data := randomizer.GenerateIntSlice(100, MaxInt)
	err := Ints(data, 0, MaxInt)
	if err != nil {
		t.Fatalf("Sorting interrupted with error: %v", err)
	}
	if !sort.IsSorted(sort.IntSlice(data)) {
		t.Fatalf("Data is not sorted %v", data)
	}
}

func benchmarkInts(b *testing.B, size int) {
	for i := 1; i <= b.N; i++ {
		b.StopTimer()
		data := randomizer.GenerateIntSlice(size, 999)
		b.StartTimer()

		Ints(data, 0, 999)
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
