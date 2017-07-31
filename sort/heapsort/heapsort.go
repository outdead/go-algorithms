//
// Package heapsort contains sorting methods based on a binary heap.
//
// Heapsort is a comparison-based sorting algorithm. Heapsort can be
// thought of as an improved selection sort: like that algorithm, it
// divides its input into a sorted and an unsorted region, and it
// iteratively shrinks the unsorted region by extracting the
// largest/smallest element and moving that to the sorted region.
// The improvement consists of the use of a heap data structure rather
// than a linear-time search to find the maximum
//
// Data structure              Array
// Best-case performance       O(n log n)
// Average performance         O(n log n)
// Worst-case performance      O(n log n)
// Worst-case space complexity O(1) auxiliary
//
package heapsort

import (
	"github.com/ganelon13/go-algorithms/structures/binaryheap"
	"sort"
)

// IntsWithHeap sorts a slice of ints in increasing order.
// Method Use the implemented binary heap
func IntsWithHeap(data []int) {
	h := binaryheap.Init(data, binaryheap.MIN)
	for i := 0; h.Len() > 0; i++ {
		data[i], _ = h.Pop()
	}
}

// Ints sorts a slice of ints in increasing order.
// Method independently implements a binary heap.
func Ints(data []int) {
	for start := (len(data) - 2) / 2; start >= 0; start-- {
		siftDownInts(data, start, len(data)-1)
	}
	for end := len(data) - 1; end > 0; end-- {
		data[0], data[end] = data[end], data[0]
		siftDownInts(data, 0, end-1)
	}
}

// Sort sorting any type that implements sort.Interface.
// Method independently implements a binary heap.
func Sort(data sort.Interface) {
	for start := (data.Len() - 2) / 2; start >= 0; start-- {
		siftDownInterface(data, start, data.Len()-1)
	}
	for end := data.Len() - 1; end > 0; end-- {
		data.Swap(0, end)
		siftDownInterface(data, 0, end-1)
	}
}

func siftDownInts(data []int, start, end int) {
	for root := start; root*2+1 <= end; {
		child := root*2 + 1
		if child+1 <= end && (data[child] < data[child+1]) {
			child++
		}
		if !(data[root] < data[child]) {
			return
		}
		data[root], data[child] = data[child], data[root]
		root = child
	}
}

func siftDownInterface(data sort.Interface, start, end int) {
	for root := start; root*2+1 <= end; {
		child := root*2 + 1
		if child+1 <= end && data.Less(child, child+1) {
			child++
		}
		if !data.Less(root, child) {
			return
		}
		data.Swap(root, child)
		root = child
	}
}
