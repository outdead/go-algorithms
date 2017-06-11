//
// Package maths contains helpers to the mathematical calculations
//
package maths

import "testing"

func testRound(t *testing.T, val float64, want int) {
	round := Round(val)
	if round != want {
		t.Fatalf("Round(%v) returns %d, want %d", val, round, want)
	}
}

func TestRound(t *testing.T) {
	testRound(t, -4.27, -4)
	testRound(t, -0.67, -1)
	testRound(t, 0.67, 1)
	testRound(t, 1.21, 1)
	testRound(t, 1.51, 2)
	testRound(t, 7.49999999, 7)
	testRound(t, 7.50000001, 8)
	testRound(t, 2880067194370816.500001, 2880067194370817)

	// Next test will fail:
	// testRound(t, 28800671943708162, 28800671943708162)
	/*
		28800671943708162 --> 28800671943708160
		28800671943708162  ~  2.880067194370816e+16
		2.880067194370816e+16 + 0.5 ~ 2.880067194370816e+16
	*/
}
