//
// Package insertionsort provides sort operation for []int type
//
// Data structure	           Array
// Best-case performance       O(n) comparisons, O(1) swaps
// Average performance	       О(n^2) comparisons, swaps
// Worst-case performance      О(n^2) comparisons, swaps
// Worst-case space complexity О(n) total, O(1) auxiliary
//
package insertionsort

// Ints sorts a slice of ints in increasing order.
func Ints(array []int) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j-1], array[j] = array[j], array[j-1]
		}
	}
}
