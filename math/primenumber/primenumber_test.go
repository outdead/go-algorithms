//
// Package primenumber implements sieve of Eratosthenes and sieve of Atkin algorithms.
//
package primenumber

import (
	"github.com/ganelon13/go-algorithms/utils/reflects"
	"testing"
)

var result []int

func testIsPrime(t *testing.T, number int, want bool) {
	if isPrime := IsPrime(number); isPrime != want {
		t.Fatalf("IsPrime(%d) returns %v, want %v", number, isPrime, want)
	}
}

func TestIsPrime(t *testing.T) {
	data := []struct {
		Number int
		Want   bool
	}{
		{-1, false},
		{0, false}, {1, false}, {2, true}, {3, true}, {4, false},
		{5, true}, {6, false}, {7, true}, {8, false}, {9, false},
		{137, true}, {719, true}, {199999, true}, {200000, false},
		{2147483646, false}, {2147483647, true},
		//{3203431780337, true},
	}

	for _, v := range data {
		testIsPrime(t, v.Number, v.Want)
	}
}

func testGetSequence(t *testing.T, f func(int) []int, checkEach bool, limit int, wantLength int) {
	sequence := f(limit)

	if length := len(sequence); length != wantLength {
		t.Fatalf("%s, get len(sequence)=%d, want %d", reflects.GetFuncName(f), length, wantLength)
	}

	if checkEach {
		for i := 0; i < len(sequence)-1; i++ {
			if !IsPrime(sequence[i]) {
				t.Fatalf("%s, sequence[%d]=%d not is prime number", reflects.GetFuncName(f), i, sequence[i])
			}
		}
	}
}

func TestGetSequenceWithDivider(t *testing.T) {
	testGetSequence(t, GetSequenceWithTrialDivision, true, 10, 4)
	testGetSequence(t, GetSequenceWithTrialDivision, true, 10000, 1229)
	testGetSequence(t, GetSequenceWithTrialDivision, true, 100000, 9592)
	testGetSequence(t, GetSequenceWithTrialDivision, true, 104729, 10000)
	testGetSequence(t, GetSequenceWithTrialDivision, true, 1000000, 78498)
	//testGetSequence(t, GetSequenceWithDividers, false, 2147483647, 105097565) // too slow
}

func TestGetSequenceWithSieveOfEratosthenes(t *testing.T) {
	testGetSequence(t, GetSequenceWithSieveOfEratosthenes, true, 10, 4)
	testGetSequence(t, GetSequenceWithSieveOfEratosthenes, true, 10000, 1229)
	testGetSequence(t, GetSequenceWithSieveOfEratosthenes, true, 100000, 9592)
	testGetSequence(t, GetSequenceWithSieveOfEratosthenes, true, 104729, 10000)
	testGetSequence(t, GetSequenceWithSieveOfEratosthenes, true, 1000000, 78498)
	//testGetSequence(t, GetSequenceWithSieveOfEratosthenes, false, 2147483647, 105097565) // 47.505s
}

func TestGetSequenceWithSieveOfAtkin(t *testing.T) {
	testGetSequence(t, GetSequenceWithSieveOfAtkin, true, 10, 4)
	testGetSequence(t, GetSequenceWithSieveOfAtkin, true, 10000, 1229)
	testGetSequence(t, GetSequenceWithSieveOfAtkin, true, 100000, 9592)
	testGetSequence(t, GetSequenceWithSieveOfAtkin, true, 104729, 10000)
	testGetSequence(t, GetSequenceWithSieveOfAtkin, true, 1000000, 78498)
	//testGetSequence(t, GetSequenceWithSieveOfAtkin, false, 2147483647, 105097565) // 80.932s
}

func benchmarkIsPrime(b *testing.B, f func(int) bool, number int) {
	for i := 1; i <= b.N; i++ {
		_ = f(number)
	}
}

func benchmarkGetSequence(b *testing.B, f func(int) []int, limit int) {
	var sequence []int
	for i := 1; i <= b.N; i++ {
		sequence = f(limit)
	}
	result = sequence
}

func BenchmarkIsPrime2147483646(b *testing.B) {
	benchmarkIsPrime(b, IsPrime, 2147483646)
}
func BenchmarkIsPrime2147483647(b *testing.B) {
	benchmarkIsPrime(b, IsPrime, 2147483647)
}

func BenchmarkGetSequenceWithTrialDivision1000000(b *testing.B) {
	benchmarkGetSequence(b, GetSequenceWithTrialDivision, 1000000)
}
func BenchmarkGetSequenceWithTrialDivision10000000(b *testing.B) {
	benchmarkGetSequence(b, GetSequenceWithTrialDivision, 10000000)
}
func BenchmarkGetSequenceWithTrialDivision100000000(b *testing.B) {
	// Too slow
	b.N = 0
	b.ResetTimer()
}
func BenchmarkGetSequenceWithTrialDivision1000000000(b *testing.B) {
	// Too slow
	b.N = 0
	b.ResetTimer()
}

func BenchmarkGetSequenceWithSieveOfEratosthenes1000000(b *testing.B) {
	benchmarkGetSequence(b, GetSequenceWithSieveOfEratosthenes, 1000000)
}
func BenchmarkGetSequenceWithSieveOfEratosthenes10000000(b *testing.B) {
	benchmarkGetSequence(b, GetSequenceWithSieveOfEratosthenes, 10000000)
}
func BenchmarkGetSequenceWithSieveOfEratosthenes100000000(b *testing.B) {
	benchmarkGetSequence(b, GetSequenceWithSieveOfEratosthenes, 100000000)
}
func BenchmarkGetSequenceWithSieveOfEratosthenes1000000000(b *testing.B) {
	benchmarkGetSequence(b, GetSequenceWithSieveOfEratosthenes, 1000000000)
}

func BenchmarkGetSequenceWithSieveOfAtkin1000000(b *testing.B) {
	benchmarkGetSequence(b, GetSequenceWithSieveOfAtkin, 1000000)
}
func BenchmarkGetSequenceWithSieveOfAtkin10000000(b *testing.B) {
	benchmarkGetSequence(b, GetSequenceWithSieveOfAtkin, 10000000)
}
func BenchmarkGetSequenceWithSieveOfAtkin100000000(b *testing.B) {
	benchmarkGetSequence(b, GetSequenceWithSieveOfAtkin, 100000000)
}
func BenchmarkGetSequenceWithSieveOfAtkin1000000000(b *testing.B) {
	benchmarkGetSequence(b, GetSequenceWithSieveOfAtkin, 1000000000)
}
