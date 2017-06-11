//
// Package interpolationsearch implements interpolation search
//
// Data structure              Array
// Best-case performance       O(1)
// Average performance         O(log log n)
// Worst-case performance      O(n)
// Worst-case space complexity O(1) iterative
//
package interpolationsearch

// Ints searches for needle in data not sorted slice of ints and returns the index
// as specified by Search. The return value is the index to insert needle if needle is
// not present (it could be len(data)).
// The slice must be sorted in ascending order.
func Ints(data []int, needle int) int {
	left := 0
	right := len(data) - 1
	mid := 0

	if right == 0 || needle < data[0] || needle > data[right] {
		return -1 // Out of range
	}

	for data[left] < needle && data[right] > needle {
		mid = left + ((needle-data[left])*(right-left))/(data[right]-data[left])

		if data[mid] < needle {
			left = mid + 1
		} else if data[mid] > needle {
			right = mid - 1
		} else {
			return mid
		}
	}

	if data[left] == needle {
		return left
	} else if data[right] == needle {
		return right
	}

	return -1 // Not found
}
