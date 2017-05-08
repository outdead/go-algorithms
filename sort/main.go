package main

import (
	"fmt"
	"sort"
	"runtime"
	"strings"
	"encoding/binary"
	"bytes"
)

func main() {
	// 25
	var initial = []int{25, 13, 99, 84, 99, 5, 21, 17, 12, 97, 45, 44, 1, 92, 25, 13, 5, 85, 2, 47, 56, 21, 66, 24, 75}

	// 100
	//var initial = []int{
	//	273, 217, 481, 821, 502, 254, 4, 233, 674, 181, 733, 563, 571, 447, 772, 125, 29, 511, 354, 656, 216, 621, 229, 995, 95,
	//	25, 13, 99, 84, 99, 5, 21, 25, 12, 99, 45, 44, 1, 92, 25, 13, 5, 85, 2, 47, 56, 21, 666, 24, 75,
	//	679, 930, 500, 68, 486, 550, 707, 592, 142, 312, 131, 495, 468, 88, 160, 155, 274, 470, 262, 490, 54, 357, 937, 849, 919,
	//	732, 499, 525, 504, 412, 712, 119, 90, 142, 634, 534, 594, 188, 474, 854, 630, 359, 464, 525, 20, 401, 286, 328, 266, 545,
	//}

	array := make([]int, len(initial))
	fmt.Printf("Initial array ...... %v\n\n", initial)

	copy(array, initial)
	sort.Ints(array)
	fmt.Println("Native sort Ints ...", array)

	copy(array, initial)
	sort.Sort(sort.IntSlice(array))
	fmt.Println("Native Quick sort ..", array)

	copy(array, initial)
	BubbleSort(array)
	fmt.Println("Bubble sort: .......", array)

	copy(array, initial)
	InsertionSort(array)
	fmt.Println("Insertion sort .....", array)

	copy(array, initial)
	ShellSort(array)
	fmt.Println("Shell sort: ........", array)

	copy(array, initial)
	SelectionSort(array)
	fmt.Println("Selection sort .....", array)

	copy(array, initial)
	MergeSort(array)
	fmt.Println("Merge sort .........", array)

	copy(array, initial)
	MergeBubbleSort(array)
	fmt.Println("Merge-Bubble sort ..", array)

	copy(array, initial)
	MergeQuickSort(array)
	fmt.Println("Merge-Quick sort ...", array)

	copy(array, initial)
	QuickSort(array)
	fmt.Println("Quick sort .........", array)

	copy(array, initial)
	CountingSort(array, 0, 1000)
	fmt.Println("Counting sort ......", array)

	array2 := []int32{25, 13, 99, 84, 99, 5, 21, 17, 12, 99, 45, 44, 1, 92, 25, 13, 5, 85, 2, 47, 56, 21, 66, 24, 75}
	RadixSort(array2, 4, 10)
	fmt.Println("Radix sort .........", array2)
}

func BubbleSort(array []int) {
	var sorted bool
	for !sorted {
		sorted = true
		for i := 0; i < len(array) - 1; i++ {
			if array[i] > array[i + 1] {
				array[i], array[i + 1] = array[i + 1], array[i]
				sorted = false
			}
		}
	}
}

func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j - 1]; j-- {
			array[j - 1], array[j] = array[j], array[j - 1]
		}
	}
}

// Sedgewick variant - Insertion
func ShellSort(array []int) {
	h := 0
	for h = 1; h <= len(array) / 9; h = 3 * h + 1 {}

	for ; h > 0; h /= 3 {
		for i := h; i < len(array); i++ {
			for j := i; j >= h && array[j] < array[j - h]; j = j - h {
				array[j], array[j - h] = array[j - h], array[j]
			}
		}
	}
}

func SelectionSort(array []int) {
	for i := 0; i < len(array) - 1; i++ {
		minValue := array[i]
		minKey := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < minValue {
				minValue = array[j]
				minKey = j
			}
		}
		array[i], array[minKey] = minValue, array[i]
	}
}

func MergeSort(array []int) {
	if len(array) < 2 {
		return
	}

	mid := len(array) / 2

	MergeSort(array[:mid])
	MergeSort(array[mid:])
	if array[mid - 1] <= array[mid] {
		return
	}

	merge(array, mid)
}

func MergeBubbleSort(array []int) {
	if len(array) < 2 {
		return
	}

	mid := len(array) / 2

	if mid <= 3 {
		BubbleSort(array)
		return
	}

	MergeBubbleSort(array[:mid])
	MergeBubbleSort(array[mid:])
	if array[mid - 1] <= array[mid] {
		return
	}

	merge(array, mid)
}

func MergeQuickSort(array []int) {
	if len(array) < 2 {
		return
	}

	mid := len(array) / 2

	if mid <= 5 {
		QuickSort(array)
		return
	}

	MergeQuickSort(array[:mid])
	MergeQuickSort(array[mid:])
	if array[mid - 1] <= array[mid] {
		return
	}

	merge(array, mid)
}

func merge(array []int, mid int) {
	// merge step, with the copy-half optimization
	var step = make([]int, mid + 1)
	copy(step, array[:mid])

	left, right := 0, mid
	for i := 0; ; i++ {
		if step[left] <= array[right] {
			array[i] = step[left]
			left++
			if left == mid {
				break
			}
		} else {
			array[i] = array[right]
			right++
			if right == len(array) {
				copy(array[i + 1:], step[left:mid])
				break
			}
		}
	}
}

// @source: https://rosettacode.org/wiki/Sorting_algorithms/Quicksort#Go
func QuickSort(a []int) {
	var pex func(int, int)
	pex = func(lower, upper int) {
		for {
			switch upper - lower {
			case -1, 0: // 0 or 1 item in segment.  nothing to do here!
				return
			case 1: // 2 items in segment
				// < operator respects strict weak order
				if a[upper] < a[lower] {
					// a quick exchange and we're done.
					a[upper], a[lower] = a[lower], a[upper]
				}
				return
			// Hoare suggests optimized sort-3 or sort-4 algorithms here,
			// but does not provide an algorithm.
			}

			// Hoare stresses picking a bound in a way to avoid worst case
			// behavior, but offers no suggestions other than picking a
			// random element.  A function call to get a random number is
			// relatively expensive, so the method used here is to simply
			// choose the middle element.  This at least avoids worst case
			// behavior for the obvious common case of an already sorted list.
			bx := (upper + lower) / 2
			b := a[bx]  // b = Hoare's "bound" (aka "pivot")
			lp := lower // lp = Hoare's "lower pointer"
			up := upper // up = Hoare's "upper pointer"
			outer:
			for {
				// use < operator to respect strict weak order
				for lp < upper && !(b < a[lp]) {
					lp++
				}
				for {
					if lp > up {
						// "pointers crossed!"
						break outer
					}
					// < operator for strict weak order
					if a[up] < b {
						break // inner
					}
					up--
				}
				// exchange
				a[lp], a[up] = a[up], a[lp]
				lp++
				up--
			}
			// segment boundary is between up and lp, but lp-up might be
			// 1 or 2, so just call segment boundary between lp-1 and lp.
			if bx < lp {
				// bound was in lower segment
				if bx < lp - 1 {
					// exchange bx with lp-1
					a[bx], a[lp - 1] = a[lp - 1], b
				}
				up = lp - 2
			} else {
				// bound was in upper segment
				if bx > lp {
					// exchange
					a[bx], a[lp] = a[lp], b
				}
				up = lp - 1
				lp++
			}
			// "postpone the larger of the two segments" = recurse on
			// the smaller segment, then iterate on the remaining one.
			if up - lower < upper - lp {
				pex(lower, up)
				lower = lp
			} else {
				pex(lp, upper)
				upper = up
			}
		}
	}
	pex(0, len(a) - 1)
}

// @source: https://rosettacode.org/wiki/Sorting_algorithms/Counting_sort#Go
func CountingSort(a []int, aMin, aMax int) {
	defer func() {
		if x := recover(); x != nil {
			// one error we'll handle and print a little nicer message
			if _, ok := x.(runtime.Error); ok &&
				strings.HasSuffix(x.(error).Error(), "index out of range") {
				fmt.Printf("data value out of range (%d..%d)\n", aMin, aMax)
				return
			}
			// anything else, we re-panic
			panic(x)
		}
	}()

	count := make([]int, aMax - aMin + 1)
	for _, x := range a {
		count[x - aMin]++
	}
	z := 0
	// optimization over task pseudocode:   variable c is used instead of
	// count[i-min].  This saves some unneccessary calculations.
	for i, c := range count {
		for ; c > 0; c-- {
			a[z] = i + aMin
			z++
		}
	}
}

// @source: https://rosettacode.org/wiki/Sorting_algorithms/Radix_sort#Go
func RadixSort(array []int32, wordLen int, high uint8) {
	var highBit int32 = -1 << high

	buf := bytes.NewBuffer(nil)
	ds := make([][]byte, len(array))
	for i, x := range array {
		binary.Write(buf, binary.LittleEndian, x ^ highBit)
		b := make([]byte, wordLen)
		buf.Read(b)
		ds[i] = b
	}
	bins := make([][][]byte, 256)
	for i := 0; i < wordLen; i++ {
		for _, b := range ds {
			bins[b[i]] = append(bins[b[i]], b)
		}
		j := 0
		for k, bs := range bins {
			copy(ds[j:], bs)
			j += len(bs)
			bins[k] = bs[:0]
		}
	}

	var w int32
	for i, b := range ds {
		buf.Write(b)
		binary.Read(buf, binary.LittleEndian, &w)
		array[i] = w ^ highBit
	}
}
