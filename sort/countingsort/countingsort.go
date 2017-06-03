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
	"errors"
	"fmt"
)

const (
	MinInt = 0
	MaxInt = 65535 // Will use 524288 Bytes
)

var LimitsError = errors.New("Incorrect limits")

// Ints sorts a slice of ints in increasing order.
func Ints(data []int, minValue int, maxValue int) (err error) {
	if len(data) == 0 {
		return
	}

	if minValue < MinInt || maxValue > MaxInt {
		err = LimitsError
		return
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			return
		}
	}()

	counts := make([]int, maxValue-minValue+1)
	for _, value := range data {
		counts[value-minValue]++
	}

	i := 0
	for value, count := range counts {
		for ; count > 0; count-- {
			data[i] = value + minValue
			i++
		}
	}
	return
}
