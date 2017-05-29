//
// Package insertionsort provides sort operation for []int type
//
// Data structure	           Array
// Best-case performance       O(n) comparisons, O(1) swaps
// Average performance	       О(n^2) comparisons, swaps
// Worst-case performance      О(n^2) comparisons, swaps
// Worst-case space complexity O(1) auxiliary
//
package insertionsort

// Ints sorts a slice of ints in increasing order.
func Ints(data []int) {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0 && data[j] < data[j-1]; j-- {
			data[j-1], data[j] = data[j], data[j-1]
		}
	}
}
