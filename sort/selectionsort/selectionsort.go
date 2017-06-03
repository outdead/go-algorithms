//
// Package selectionsort provides sort operation for []int type
//
// Data structure              Array
// Best-case performance       О(n^2)
// Average performance         О(n^2)
// Worst-case performance	   О(n^2)
// Worst-case space complexity O(1) auxiliary
//
package selectionsort

// Ints sorts a slice of ints in increasing order.
func Ints(data []int) {
	for i := 0; i < len(data)-1; i++ {
		minKey, minValue := i, data[i]
		for j := i + 1; j < len(data); j++ {
			if data[j] < minValue {
				minKey, minValue = j, data[j]
			}
		}
		data[i], data[minKey] = minValue, data[i]
	}
}
