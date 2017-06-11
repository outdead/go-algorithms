//
// Package binarysearch implements binary search
//
// Data structure              Array
// Best-case performance       O(1)
// Average performance         O(log n)
// Worst-case performance      O(log n)
// Worst-case space complexity O(1)
//
package binarysearch

// Ints searches for needle in data not sorted slice of ints and returns the index
// as specified by Search. The return value is the index to insert needle if needle is
// not present (it could be len(data)).
// The slice must be sorted in ascending order.
func Ints(data []int, needle int) int {
	left := 0
	right := len(data)

	if right == 0 || needle < data[0] || needle > data[right-1] {
		return -1 // Out of range
	}

	for left < right {
		mid := left + (right-left)/2

		if needle <= data[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if data[right] == needle {
		return right
	}

	return -1 // Not found
}

// IntsRecursion searches for needle in data not sorted slice of ints and returns the index
// as specified by Search. The return value is the index to insert needle if needle is
// not present (it could be len(data)).
// The slice must be sorted in ascending order.
func IntsRecursion(data []int, needle int) int {
	left := 0
	right := len(data)
	if right == 0 || needle < data[left] || needle > data[right-1] {
		return -1 // Out of range
	}
	return intsRecursion(data, needle, left, right-1)
}

func intsRecursion(data []int, needle int, left int, right int) int {
	if left >= right {
		if data[left] == needle {
			return left
		}
		return -1 // Not found
	}

	mid := left + (right-left)/2

	if v := data[mid]; v > needle {
		return intsRecursion(data, needle, left, mid)
	} else if v < needle {
		return intsRecursion(data, needle, mid+1, right)
	}

	return mid
}
