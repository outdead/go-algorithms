package main

import (
	"runtime"
	"testing"
)

const N = 1000000

var (
	needle = 61
	n = 137
)

func init() {
	runtime.GOMAXPROCS(1)
}

func BenchmarkIsPrimeDivider(b *testing.B) {
	b.N = N

	for i := 1; i <= b.N; i++ {
		_ = IsPrimeDivider(needle)
	}
}

func BenchmarkIsPrimeEratosthenes(b *testing.B) {
	b.N = N

	for i := 1; i <= b.N; i++ {
		_ = IsPrimeEratosthenes(needle)
	}
}

func BenchmarkGeneratePrimeDivider(b *testing.B) {
	b.N = N

	for i := 1; i <= b.N; i++ {
		_ = GeneratePrimeDivider(n)
	}
}

func BenchmarkGeneratePrimeEratosthenes(b *testing.B) {
	b.N = N

	for i := 1; i <= b.N; i++ {
		_ = GeneratePrimeEratosthenes(n)
	}
}

func BenchmarkGeneratePrimeAtkin(b *testing.B) {
	b.N = N

	for i := 1; i <= b.N; i++ {
		_ = GeneratePrimeAtkin(n)
	}
}
