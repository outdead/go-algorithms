//
// Package radixsort provides sort operation for []int type with LSD radix sort algorithm.
//
// Numbers are sorted by discharges. There are two options: least significant digit (LSD)
// and the most significant digit (MSD). With LSD sorting, low-order digits are first sorted,
// then older ones. With MSD sorting, everything is vice versa. With LSD sorting, the order
// is as follows: short keys go before long, keys of the same size are sorted alphabetically,
// this coincides with the normal representation of numbers 1, 2, 3, 4, 5, 6, 7, 8, 9, 10.
// With MSD sort an alphabetical order is obtained that is suitable for sorting strings
//
// Data structure              Array
// Best-case performance       O(n)
// Average performance         O(n * k)
// Worst-case performance      O(n * k)
// Worst-case space complexity O(n + k) auxiliary
//
package radixsort

const (
	bitSize uint = 32 * (1 + ^uint(0)>>63)
	minInt  int  = -1 << (bitSize - 1)
)

// Ints sorts a slice of ints in increasing order.
func Ints(data []int) {
	if len(data) < 2 {
		return
	}

	buff := make([]int, len(data))
	offsets := make([]int, 256)

	for keyOffset := uint(0); keyOffset < bitSize; keyOffset += 8 {
		counts := make([]int, 256)
		keyMask := int(255 << keyOffset)

		sorted := true
		previousValue := minInt
		for _, value := range data {
			key := uint8((value & keyMask) >> keyOffset)
			counts[key]++
			if sorted {
				sorted = value >= previousValue
				previousValue = value
			}
		}

		if sorted {
			if (keyOffset/8)%2 == 1 {
				copy(buff, data)
			}
			return
		}

		if keyOffset != bitSize-8 {
			offsets[0] = 0
			for i := 1; i < len(offsets); i++ {
				offsets[i] = offsets[i-1] + counts[i-1]
			}
		} else {
			negativesCount := 0
			for i := 128; i < 256; i++ {
				negativesCount += counts[i]
			}

			offsets[0], offsets[128] = negativesCount, 0
			for i := 1; i < 128; i++ {
				offsets[i] = offsets[i-1] + counts[i-1]         // Positive values
				offsets[i+128] = offsets[i+127] + counts[i+127] // Negative values
			}
		}

		for _, value := range data {
			key := uint8((value & keyMask) >> keyOffset) // Get a byte of each value in the radix
			buff[offsets[key]] = value
			offsets[key]++
		}
		buff, data = data, buff
	}
}
