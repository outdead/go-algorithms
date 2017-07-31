//
// Package fibonacci implements recursive, closure and rounding algorithms
//
package fibonacci

import (
	"github.com/ganelon13/go-algorithms/utils/maths"
	"math"
)

// FillSequence creates fibonacci sequence and returns last generated
// fibonacci number. Recursion method was implemented with remembering
// previous iterations.
//
// Data structure              Array
// Best-case performance       O(n)
// Average performance         O(n)
// Worst-case performance      O(n)
// Worst-case space complexity O(n) auxiliary
//
func FillSequence(index int, sequence *[]int) int {
	if index < 0 || index > 92 {
		return -1
	}

	if len(*sequence)-1 < index {
		*sequence = append(*sequence, make([]int, index-len(*sequence)+1)...)
	}

	if index == 0 {
		(*sequence)[0] = 0
		return 0
	} else if index == 1 {
		(*sequence)[1] = 1
		return 1
	}

	if (*sequence)[index] > 0 {
		return (*sequence)[index]
	}

	(*sequence)[index] = FillSequence(index-1, sequence) + FillSequence(index-2, sequence)
	return (*sequence)[index]
}

// GetNumber calculates fibonacci number with index index
//
// Data structure              Array
// Best-case performance       O(log(n))
// Average performance         O(log(n))
// Worst-case performance      O(log(n))
// Worst-case space complexity O(1) auxiliary
//
func GetNumber(index int) int {
	if index < 0 || index > 92 {
		return -1
	}

	fib := func() func() int {
		a, b := 0, 1
		return func() int {
			a, b = b, a+b
			return a
		}
	}

	f := fib()
	var number int
	for i := 0; i < index; i++ {
		number = f()
	}

	return number
}

// ComputeApproximateNumber calculates fibonacci number with index index
// After ComputeApproximateNumber(78) this implementation is not exact.
//
// Data structure              Array
// Best-case performance       O(1)
// Average performance         O(1)
// Worst-case performance      O(1)
// Worst-case space complexity O(1)
//
func ComputeApproximateNumber(index int) int {
	if index < 0 || index > 78 {
		return -1
	}

	sqrt5 := math.Sqrt(float64(5))
	var phi float64 = (sqrt5 + 1) / 2
	return maths.Round(math.Pow(phi, float64(index)) / sqrt5)
}
