//
// Package fibonacci implements recursive, closure and rounding algorithms
//
package fibonacci

import (
	"math/rand"
	"testing"
)

func testFillSequence(t *testing.T, index int, wantNumber int) {
	var sequence []int
	number := FillSequence(index, &sequence)
	if number != wantNumber {
		t.Fatalf("FillSequence(%d) returns %v, want %v", index, number, wantNumber)
	}
}

func testGetNumber(t *testing.T, index int, want int) {
	number := GetNumber(index)
	if number != want {
		t.Fatalf("GetNumber(%d) returns %v, want %v", index, number, want)
	}
}

func testComputeApproximateNumber(t *testing.T, index int, want int) {
	number := ComputeApproximateNumber(index)
	if number != want {
		t.Fatalf("ComputeApproximateNumber(%d) returns %v, want %v", index, number, want)
	}
}

func TestFillSequence(t *testing.T) {
	testFillSequence(t, -1, -1)
	testFillSequence(t, 0, 0)
	testFillSequence(t, 1, 1)
	testFillSequence(t, 2, 1)
	testFillSequence(t, 10, 55)
	testFillSequence(t, 25, 75025)
	testFillSequence(t, 50, 12586269025)
	testFillSequence(t, 78, 8944394323791464)
	testFillSequence(t, 92, 7540113804746346429)
	testFillSequence(t, 93, -1)
}

func TestGetNumber(t *testing.T) {
	testGetNumber(t, -1, -1)
	testGetNumber(t, 0, 0)
	testGetNumber(t, 1, 1)
	testGetNumber(t, 2, 1)
	testGetNumber(t, 10, 55)
	testGetNumber(t, 25, 75025)
	testGetNumber(t, 50, 12586269025)
	testGetNumber(t, 78, 8944394323791464)
	testGetNumber(t, 92, 7540113804746346429)
	testGetNumber(t, 93, -1)
}

func TestComputeApproximateNumber(t *testing.T) {
	testComputeApproximateNumber(t, -1, -1)
	testComputeApproximateNumber(t, 0, 0)
	testComputeApproximateNumber(t, 1, 1)
	testComputeApproximateNumber(t, 2, 1)
	testComputeApproximateNumber(t, 10, 55)
	testComputeApproximateNumber(t, 25, 75025)
	testComputeApproximateNumber(t, 50, 12586269025)
	testComputeApproximateNumber(t, 78, 8944394323791464)
	testComputeApproximateNumber(t, 79, -1)
}

func benchmarkFillSequence(b *testing.B, size int, want int) {
	var number int
	for i := 1; i <= b.N; i++ {
		var sequence []int
		number = FillSequence(size, &sequence)
		if number != want {
			b.Fatalf("FillSequence returns %d, want %d", number, want)
		}
	}
}

func benchmarkGetNumber(b *testing.B, size int, want int) {
	var number int
	for i := 1; i <= b.N; i++ {
		number = GetNumber(size)
		if number != want {
			b.Fatalf("GetNumber returns %d, want %d", number, want)
		}
	}
}

func benchmarkGetNumberFromSequence(b *testing.B, size int) {
	var controlSequence []int
	FillSequence(size, &controlSequence)

	var number int
	var sequence []int
	for i := 1; i <= b.N; i++ {
		index := rand.Intn(size)
		number = FillSequence(index, &sequence)
		if number != controlSequence[index] {
			b.Fatalf("GetNumber returns %d, want %d", number, controlSequence[index])
		}
	}
}

func benchmarkComputeApproximateNumber(b *testing.B, size int, want int) {
	var number int
	for i := 1; i <= b.N; i++ {
		number = ComputeApproximateNumber(size)
		if number != want {
			b.Fatalf("ComputeApproximateNumber returns %d, want %d", number, want)
		}
	}
}

func BenchmarkFillSequence10(b *testing.B) {
	benchmarkFillSequence(b, 10, 55)
}

func BenchmarkFillSequence25(b *testing.B) {
	benchmarkFillSequence(b, 25, 75025)
}

func BenchmarkFillSequence50(b *testing.B) {
	benchmarkFillSequence(b, 50, 12586269025)
}

func BenchmarkFillSequence_GetRand50(b *testing.B) {
	benchmarkGetNumberFromSequence(b, 50)
}

func BenchmarkGetNumber10(b *testing.B) {
	benchmarkGetNumber(b, 10, 55)
}

func BenchmarkGetNumber25(b *testing.B) {
	benchmarkGetNumber(b, 25, 75025)
}

func BenchmarkGetNumber50(b *testing.B) {
	benchmarkGetNumber(b, 50, 12586269025)
}

func BenchmarkComputeApproximateNumber10(b *testing.B) {
	benchmarkComputeApproximateNumber(b, 10, 55)
}

func BenchmarkComputeApproximateNumber25(b *testing.B) {
	benchmarkComputeApproximateNumber(b, 25, 75025)
}

func BenchmarkComputeApproximateNumber50(b *testing.B) {
	benchmarkComputeApproximateNumber(b, 50, 12586269025)
}
