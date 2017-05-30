//
// Package mergesort contains 2 variants of merge sort methods -  topDown
// implementation and bottomUp implementation. BottomUp has a slightly higher
// constant productivity in relation to the topDown implementation.
//
// Data structure              Array
// Best-case performance       O(n log n) typical, O(n) natural variant
// Average performance         O(n log n)
// Worst-case performance      O(n log n)
// Worst-case space complexity O(n) auxiliary
//
package mergesort

// Ints sorts a slice of ints in increasing order. Use bottomUp implementation
func Ints(data []int) {
	n := len(data)
	buff := make([]int, n)

	bottomUp(data, buff, n)
}

// Ints2 sorts a slice of ints in increasing order. Use topDown implementation
func Ints2(data []int) {
	n := len(data)
	buff := append([]int{}, data...)

	topDown(data, buff, 0, n)
}

// The top-down implementation is the implementation that uses recursion.
// It starts at the top of the tree and proceeds downwards, each time asking the
// same question (what do I need to do to sort this array?) and answering it
// (split it into two, make recursive calls, merge results), until you get to
// the bottom of the tree
func topDown(data []int, buff []int, left int, right int) {
	if right-left < 2 {
		return
	}

	middle := left + (right-left)/2
	topDown(buff, data, left, middle)
	topDown(buff, data, middle, right)

	merge(data, buff, left, middle, right)
}

// The bottom-up implementation doesn't use recursion. It directly starts at the
// bottom of the tree and proceeds upwards by iterating over the pieces and merging them.
func bottomUp(data []int, buff []int, n int) {
	for width := 1; width < n; width = 2 * width {
		for i := 0; i < n; i = i + 2*width {
			merge(data, buff, i, min(i+width, n), min(i+2*width, n))
		}
		// A more efficient implementation would swap the roles of data and buff.
		for i := 0; i < n; i++ {
			data[i] = buff[i]
		}
	}
}

func merge(data []int, buff []int, left int, middle int, right int) {
	i := left
	j := middle

	for k := left; k < right; k++ {
		if i < middle && (j >= right || buff[i] <= buff[j]) {
			data[k] = buff[i]
			i = i + 1
		} else {
			data[k] = buff[j]
			j = j + 1
		}
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
