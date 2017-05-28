//
// Package bubblesort contains realisation of bubble sort algorithm.
//
// Bubble sort has worst-case and average complexity both О(n^2),
// where n is the number of items being sorted.
//
// Data structure              Array
// Best-case performance       O(n)
// Average performance         O(n^2)
// Worst-case performance      O(n^2)
// Worst-case space complexity O(1) auxiliary
//
package bubblesort

// Ints sorts a slice of ints in increasing order.
func Ints(a []int) {
	var sorted bool
	for !sorted {
		sorted = true
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				sorted = false
			}
		}
	}
}
