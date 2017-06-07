//
// Package quicksort contains implementation of basic quicksort algorithm
// without improvements. Quicksort has a reputation as the fastest sort.
// Optimized variants of quicksort are common features of many languages
// and libraries. One often contrasts quicksort with merge sort, because
// both sorts have an average time of O(n log n)
//
// Data structure              Array
// Best-case performance       О(n log n)
// Average performance         О(n log n)
// Worst-case performance      О(n^2)
// Worst-case space complexity O(1) auxiliary
//
package quicksort

// Ints sorts a slice of ints in increasing order.
func Ints(data []int) {
	ints(data, 0, len(data)-1)
}

func ints(data []int, left int, right int) {
	if left >= right {
		return
	}

	pivotIndex := intsPivot(data, left, right)
	// TODO: check maxDepth is needed
	ints(data, left, pivotIndex-1)
	ints(data, pivotIndex+1, right)
}

func intsPivot(data []int, left int, right int) int {
	// middle is current pivot index
	middle := left + (right-left)/2

	// TODO: medianOfThree

	data[left], data[middle] = data[middle], data[left]
	first := left + 1
	last := right
	for first <= last {
		for first <= right && data[first] < data[left] {
			first++
		}
		for last >= left && data[left] < data[last] {
			last--
		}
		if first <= last {
			data[first], data[last] = data[last], data[first]
			first++
			last--
		}
	}
	data[left], data[last] = data[last], data[left]
	return last
}
