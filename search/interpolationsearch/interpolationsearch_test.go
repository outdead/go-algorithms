//
// Package interpolationsearch implements interpolation search
//
// Data structure              Array
// Best-case performance       O(1)
// Average performance         O(log log n)
// Worst-case performance      O(n)
// Worst-case space complexity O(1) iterative
//
package interpolationsearch

import (
	"github.com/ganelon13/go-algorithms/utils/randomizer"
	"testing"
)

func TestInts(t *testing.T) {
	data := randomizer.GenerateSortedIntSlice(1000)

	wantIndex := 137
	wantValue := data[wantIndex]

	searchIndex := Ints(data, wantValue)

	if searchIndex < 0 || searchIndex > len(data) {
		t.Fatalf("searchIndex is %d, want %d", searchIndex, wantIndex)
	}

	if data[searchIndex] != data[wantIndex] {
		t.Fatalf("searchIndex is %d, want %d", searchIndex, wantIndex)
	}
}

func TestInts2(t *testing.T) {
	data := randomizer.GenerateSortedIntSlice(1000)

	wantIndex := -1
	wantValue := data[999] + 1

	searchIndex := Ints(data, wantValue)

	if searchIndex != wantIndex {
		t.Errorf("searchIndex is %d, want %d", searchIndex, wantIndex)
	}
}

func TestInts3(t *testing.T) {
	data := []int{1, 2}

	wantIndex := 1
	wantValue := 2

	searchIndex := Ints(data, wantValue)

	if searchIndex != wantIndex {
		t.Errorf("searchIndex is %d, want %d", searchIndex, wantIndex)
	}
}

func benchmarkInts(b *testing.B, size int, index int) {
	b.StopTimer()
	data := randomizer.GenerateSortedIntSlice(size)
	needle := data[index-1]
	b.StartTimer()

	for i := 1; i <= b.N; i++ {
		Ints(data, needle)
		if i == 999999999 {
			b.N = i
			return
		}
	}
}

// First

func BenchmarkIntsFirst100(b *testing.B) {
	benchmarkInts(b, 100, 1)
}

func BenchmarkIntsFirst1000(b *testing.B) {
	benchmarkInts(b, 1000, 1)
}

func BenchmarkIntsFirst10000(b *testing.B) {
	benchmarkInts(b, 10000, 1)
}

func BenchmarkIntsFirst100000(b *testing.B) {
	benchmarkInts(b, 100000, 1)
}

// Middle

func BenchmarkIntsMiddle100(b *testing.B) {
	benchmarkInts(b, 100, 50)
}

func BenchmarkIntsMiddle1000(b *testing.B) {
	benchmarkInts(b, 1000, 500)
}

func BenchmarkIntsMiddle10000(b *testing.B) {
	benchmarkInts(b, 10000, 5000)
}

func BenchmarkIntsMiddle100000(b *testing.B) {
	benchmarkInts(b, 100000, 50000)
}

// Last

func BenchmarkIntsLast100(b *testing.B) {
	benchmarkInts(b, 100, 100)
}

func BenchmarkIntsLast1000(b *testing.B) {
	benchmarkInts(b, 1000, 1000)
}

func BenchmarkIntsLast10000(b *testing.B) {
	benchmarkInts(b, 10000, 10000)
}

func BenchmarkIntsLast100000(b *testing.B) {
	benchmarkInts(b, 100000, 100000)
}
