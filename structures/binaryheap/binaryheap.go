//
// Package binaryheap provides heap operations for []int type.
// A heap is a tree with the property that each node is the
// minimum-valued/maximum-valued node in its subtree.
//
// The minimum/maximum element in the tree is the root, at index 0.
//
package binaryheap

import (
	"errors"
	"fmt"
)

type (
	sortType bool
	Heap     struct {
		container []int
		length    int
		min       sortType
	}
)

var (
	MIN sortType = true
	MAX sortType = false

	ErrEmptyHeap             = errors.New("Heap is empty")
	errIndexOutOfRangeFormat = "Index %d out of range"
)

// New allows to initialize an empty heap
func New(heapType sortType) *Heap {
	return &Heap{
		container: make([]int, 0),
		length:    0,
		min:       heapType,
	}
}

// A heap must be initialized before any of the heap operations
// can be used.
func Init(data []int, heapType sortType) *Heap {
	h := &Heap{
		container: append([]int{}, data...),
		// We can save on memory allocs, but in exchange we give unprotected access to the container
		//container: data,
		length: len(data),
		min:    heapType,
	}
	for i := h.length/2 - 1; i >= 0; i-- {
		h.down(i)
	}
	return h
}

func (h *Heap) Len() int {
	return h.length
}

func (h *Heap) Push(v int) {
	h.container = append(h.container, v)
	h.up(h.length)
	h.length = h.length + 1
}

func (h *Heap) Pop() (int, error) {
	if h.length == 0 {
		return 0, ErrEmptyHeap
	}

	return h.extract(0), nil
}

func (h *Heap) Fetch(i int) (int, error) {
	if i >= h.length || i < 0 {
		return 0, fmt.Errorf(errIndexOutOfRangeFormat, i)
	}
	return h.fetch(i), nil
}

// Replace replaces  the element at index i to new value
// and returns old one
func (h *Heap) Replace(i int, value int) (int, error) {
	if h.length == 0 {
		return 0, ErrEmptyHeap
	}

	if i >= h.length || i < 0 {
		return 0, fmt.Errorf(errIndexOutOfRangeFormat, i)
	}

	h.container[i], value = value, h.container[i]
	h.fix(i)
	return value, nil
}

// Extract fetch the element at index i from the heap and removes it.
func (h *Heap) Extract(i int) (int, error) {
	if h.length == 0 {
		return 0, ErrEmptyHeap
	}

	if i >= h.length || i < 0 {
		return 0, fmt.Errorf(errIndexOutOfRangeFormat, i)
	}

	return h.extract(i), nil
}

func (h *Heap) fetch(i int) int {
	return h.container[i]
}

func (h *Heap) extract(i int) (value int) {
	value = h.fetch(i)
	h.length = h.length - 1
	if h.length == 0 {
		h.container = nil
		return
	}

	h.swap(i, h.length)
	h.container = h.container[0:h.length]
	h.down(i)
	return
}

func (h *Heap) less(i, j int) bool {
	if h.min {
		return h.fetch(i) < h.fetch(j)
	} else {
		return h.fetch(i) > h.fetch(j)
	}
}

func (h *Heap) swap(i, j int) {
	h.container[i], h.container[j] = h.container[j], h.container[i]
}

// fix re-establishes the heap ordering after the element at index i has changed its value.
func (h *Heap) fix(i int) {
	if !h.down(i) {
		h.up(i)
	}
}

func (h *Heap) up(child int) {
	for {
		parent := (child - 1) / 2
		if parent == child || !h.less(child, parent) {
			break
		}
		h.swap(parent, child)
		child = parent
	}
}

func (h *Heap) down(parent int) bool {
	root := parent
	end := h.length - 1
	for root = parent; root*2+1 <= end; {
		child := (root << 1) + 1
		if child+1 <= end && !h.less(child, child+1) {
			child++
		}
		if h.less(root, child) {
			return root > parent
		}
		h.swap(root, child)
		root = child
	}
	return root > parent
}
