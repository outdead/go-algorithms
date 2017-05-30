//
// Package shellsort provides sort operation for []int type
//
// Use Sedgewick variant (Insertion)
//
// Data structure              Array
// Best-case performance       O(n)
// Average performance         0(n^(7/6))
// Worst-case performance      O(n^(4/3))
// Worst-case space complexity O(1) auxiliary
//
package shellsort

// Ints sorts a slice of ints in increasing order.
func Ints(data []int) {
	h := 0
	for h = 1; h <= len(data)/9; h = 3*h + 1 {
	}

	for ; h > 0; h /= 3 {
		for i := h; i < len(data); i++ {
			for j := i; j >= h && data[j] < data[j-h]; j = j - h {
				data[j], data[j-h] = data[j-h], data[j]
			}
		}
	}
}
