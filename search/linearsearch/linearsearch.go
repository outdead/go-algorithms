//
// Package linearsearch implements linear (sequential) search.
//
// Data structure              Array
// Best-case performance       O(1)
// Average performance         0(n)
// Worst-case performance      O(n)
// Worst-case space complexity O(1) iterative
//
package linearsearch

// Ints searches for needle in data not sorted slice of ints and returns the index
// as specified by Search. The return value is the index to insert needle if needle is
// not present (it could be len(data)).
func Ints(data []int, needle int) int {
	for key, value := range data {
		if value == needle {
			return key
		}
	}
	return -1
}
