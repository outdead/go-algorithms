//
// Package binarysearch implements binary search
//
// Data structure              Array
// Best-case performance       O(1)
// Average performance         O(log n)
// Worst-case performance      O(log n)
// Worst-case space complexity O(1)
//
package binarysearch

import (
	"github.com/ganelon13/go-algorithms/utils/randomizer"
	"sort"
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

func TestIntsRecursion(t *testing.T) {
	data := randomizer.GenerateSortedIntSlice(1000)

	wantIndex := 137
	wantValue := data[wantIndex]

	searchIndex := IntsRecursion(data, wantValue)

	if searchIndex < 0 || searchIndex > len(data) {
		t.Fatalf("searchIndex is %d, want %d", searchIndex, wantIndex)
	}

	if data[searchIndex] != data[wantIndex] {
		t.Fatalf("searchIndex is %d, want %d", searchIndex, wantIndex)
	}
}

func TestIntsRecursion2(t *testing.T) {
	data := randomizer.GenerateSortedIntSlice(1000)

	wantIndex := -1
	searchValue := data[999] + 1

	searchIndex := IntsRecursion(data, searchValue)

	if searchIndex != wantIndex {
		t.Errorf("searchIndex is %d, want %d", searchIndex, wantIndex)
	}
}

func TestIntsRecursion3(t *testing.T) {
	data := []int{1, 2}

	wantIndex := 1
	wantValue := 2

	searchIndex := IntsRecursion(data, wantValue)

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

func benchmarkIntsRecursion(b *testing.B, size int, index int) {
	b.StopTimer()
	data := randomizer.GenerateSortedIntSlice(size)
	needle := data[index-1]
	b.StartTimer()

	for i := 1; i <= b.N; i++ {
		IntsRecursion(data, needle)
		if i == 999999999 {
			b.N = i
			return
		}
	}
}

func benchmarkSearchInts_sortPackage(b *testing.B, size int, index int) {
	b.StopTimer()
	data := randomizer.GenerateSortedIntSlice(size)
	needle := data[index-1]
	b.StartTimer()

	for i := 1; i <= b.N; i++ {
		sort.SearchInts(data, needle)
		if i == 999999999 {
			b.N = i
			return
		}
	}
}

// IntsFirst

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

// IntsMiddle

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

// IntsLast

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

// IntsRecursionFirst

func BenchmarkIntsRecursionFirst100(b *testing.B) {
	benchmarkIntsRecursion(b, 100, 1)
}

func BenchmarkIntsRecursionFirst1000(b *testing.B) {
	benchmarkIntsRecursion(b, 1000, 1)
}

func BenchmarkIntsRecursionFirst10000(b *testing.B) {
	benchmarkIntsRecursion(b, 10000, 1)
}

func BenchmarkIntsRecursionFirst100000(b *testing.B) {
	benchmarkIntsRecursion(b, 100000, 1)
}

// IntsRecursionMiddle

func BenchmarkIntsRecursionMiddle100(b *testing.B) {
	benchmarkIntsRecursion(b, 100, 50)
}

func BenchmarkIntsRecursionMiddle1000(b *testing.B) {
	benchmarkIntsRecursion(b, 1000, 500)
}

func BenchmarkIntsRecursionMiddle10000(b *testing.B) {
	benchmarkIntsRecursion(b, 10000, 5000)
}

func BenchmarkIntsRecursionMiddle100000(b *testing.B) {
	benchmarkIntsRecursion(b, 100000, 50000)
}

// IntsRecursionLast

func BenchmarkIntsRecursionLast100(b *testing.B) {
	benchmarkIntsRecursion(b, 100, 100)
}

func BenchmarkIntsRecursionLast1000(b *testing.B) {
	benchmarkIntsRecursion(b, 1000, 1000)
}

func BenchmarkIntsRecursionLast10000(b *testing.B) {
	benchmarkIntsRecursion(b, 10000, 10000)
}

func BenchmarkIntsRecursionLast100000(b *testing.B) {
	benchmarkIntsRecursion(b, 100000, 100000)
}

// SearchInts_sortPackageFirst

func BenchmarkSearchInts_sortPackageFirst100(b *testing.B) {
	benchmarkSearchInts_sortPackage(b, 100, 1)
}

func BenchmarkSearchInts_sortPackageFirst1000(b *testing.B) {
	benchmarkSearchInts_sortPackage(b, 1000, 1)
}

func BenchmarkSearchInts_sortPackageFirst10000(b *testing.B) {
	benchmarkSearchInts_sortPackage(b, 10000, 1)
}

func BenchmarkSearchInts_sortPackageFirst100000(b *testing.B) {
	benchmarkSearchInts_sortPackage(b, 100000, 1)
}

// SearchInts_sortPackageMiddle

func BenchmarkSearchInts_sortPackageMiddle100(b *testing.B) {
	benchmarkSearchInts_sortPackage(b, 100, 50)
}

func BenchmarkSearchInts_sortPackageMiddle1000(b *testing.B) {
	benchmarkSearchInts_sortPackage(b, 1000, 500)
}

func BenchmarkSearchInts_sortPackageMiddle10000(b *testing.B) {
	benchmarkSearchInts_sortPackage(b, 10000, 5000)
}

func BenchmarkSearchInts_sortPackageMiddle100000(b *testing.B) {
	benchmarkSearchInts_sortPackage(b, 100000, 50000)
}

// SearchInts_sortPackageLast

func BenchmarkSearchInts_sortPackageLast100(b *testing.B) {
	benchmarkSearchInts_sortPackage(b, 100, 100)
}

func BenchmarkSearchInts_sortPackageLast1000(b *testing.B) {
	benchmarkSearchInts_sortPackage(b, 1000, 1000)
}

func BenchmarkSearchInts_sortPackageLast10000(b *testing.B) {
	benchmarkSearchInts_sortPackage(b, 10000, 10000)
}

func BenchmarkSearchInts_sortPackageLast100000(b *testing.B) {
	benchmarkSearchInts_sortPackage(b, 100000, 100000)
}
