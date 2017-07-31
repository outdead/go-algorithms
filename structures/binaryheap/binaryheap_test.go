//
// Package binaryheap provides heap operations for []int type.
// A heap is a tree with the property that each node is the
// minimum-valued/maximum-valued node in its subtree.
//
// The minimum/maximum element in the tree is the root, at index 0.
//
package binaryheap

import (
	"github.com/ganelon13/go-algorithms/utils/randomizer"
	"math/rand"
	"sort"
	"testing"
)

// verify verifies that the heap is filled correctly.
// Automatically recognizes min and max heap.
func (h *Heap) verify(t *testing.T, root int) {
	n := h.Len()
	left := 2*root + 1
	right := 2*root + 2
	if left < n {
		if h.less(left, root) {
			t.Fatalf("heap invariant invalidated [%d] = %d > [%d] = %d", root, h.fetch(root), left, h.fetch(left))
			//return
		}
		h.verify(t, left)
	}
	if right < n {
		if h.less(right, root) {
			t.Fatalf("heap invariant invalidated [%d] = %d > [%d] = %d", root, h.fetch(root), left, h.fetch(right))
			//return
		}
		h.verify(t, right)
	}
}

// Can panic
func TestIncorrectIndexes(t *testing.T) {
	h := New(MIN)

	//defer func() {
	//	if r := recover(); r != nil {
	//		t.Fatalf("Error recovered: %v, %s", r, debug.Stack())
	//	}
	//}()

	if _, err := h.Fetch(-1); err == nil {
		t.Error("Error expected, nil received")
	}

	if _, err := h.Fetch(0); err == nil {
		t.Error("Error expected, nil received")
	}

	if _, err := h.Pop(); err == nil {
		t.Error("Error expected, nil received")
	}

	if _, err := h.Extract(-1); err == nil {
		t.Error("Error expected, nil received")
	}

	if _, err := h.Extract(0); err == nil {
		t.Error("Error expected, nil received")
	}

	if _, err := h.Replace(-1, 0); err == nil {
		t.Error("Error expected, nil received")
	}

	if _, err := h.Replace(0, 0); err == nil {
		t.Error("Error expected, nil received")
	}
}

func TestMinInitSame(t *testing.T) {
	N := 20
	array := make([]int, 0)
	for i := N; i > 0; i-- {
		array = append(array, 0) // all elements are the same
	}
	h := Init(array, MIN)
	h.verify(t, 0)

	if h.Len() != N {
		t.Errorf("h.Len() = %d; want %d", h.Len(), N)
	}

	for i := 1; h.Len() > 0; i++ {
		x, err := h.Pop()
		if err != nil {
			t.Errorf("%d.th pop return error: %s", i, err.Error())
			return
		}
		h.verify(t, 0)
		if x != 0 {
			t.Errorf("%d.th pop got %d; want %d", i, x, 0)
		}
	}
}

func TestMinInitDifferent(t *testing.T) {
	N := 20
	array := make([]int, 0)
	for i := N; i > 0; i-- {
		array = append(array, i) // all elements are different
	}
	h := Init(array, MIN)
	h.verify(t, 0)

	if h.Len() != N {
		t.Errorf("h.Len() = %d; want %d", h.Len(), N)
	}

	for i := 1; h.Len() > 0; i++ {
		x, err := h.Pop()
		if err != nil {
			t.Errorf("%d.th pop return error: %s", i, err.Error())
			return
		}
		h.verify(t, 0)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func TestMinPush(t *testing.T) {
	h := New(MIN)
	for i := 20; i > 10; i-- {
		h.Push(i)
		h.verify(t, 0)
	}

	for i := 10; i > 0; i-- {
		h.Push(i)
		h.verify(t, 0)
	}

	for i := 1; h.Len() > 0; i++ {
		x, err := h.Pop()
		if err != nil {
			t.Errorf("%d.th pop returns error: %v", i, err.Error())
			return
		}
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
		h.verify(t, 0)
		if i < 20 {
			h.Push(20 + i)
		}
		h.verify(t, 0)
	}
}

func TestMinSort(t *testing.T) {
	slice := randomizer.GenerateIntSlice(100, 999)
	h := Init(slice, MIN)
	h.verify(t, 0)

	for i := 0; h.Len() > 0; i++ {
		var err error
		slice[i], err = h.Pop()
		if err != nil {
			t.Errorf("%d.th pop returns error: %v", i, err.Error())
			return
		}
		h.verify(t, 0)
	}

	if !sort.IsSorted(sort.IntSlice(slice)) {
		t.Errorf("Data not sortded %v", slice)
	}
}

func TestMinExtractFirst(t *testing.T) {
	array := make([]int, 0)
	for i := 0; i < 10; i++ {
		array = append(array, i)
	}
	h := Init(array, MIN)
	h.verify(t, 0)

	for i := 0; h.Len() > 0; i++ {
		x, err := h.Extract(0)
		if err != nil {
			t.Errorf("Extract(0) returns error: %s", err.Error())
			return
		}
		if x != i {
			t.Errorf("Extract(0) got %d; want %d", x, i)
		}
		h.verify(t, 0)
	}
}

func TestMinExtractLast(t *testing.T) {
	array := make([]int, 0)
	for i := 0; i < 10; i++ {
		array = append(array, i)
	}
	h := Init(array, MIN)
	h.verify(t, 0)

	for h.Len() > 0 {
		i := h.Len() - 1
		x, err := h.Extract(i)
		if err != nil {
			t.Errorf("Extract(%d) returns error: %s", i, err.Error())
			return
		}
		if x != i {
			t.Errorf("Extract(%d) got %d; want %d", i, x, i)
		}
		h.verify(t, 0)
	}
}

func TestMinExtract(t *testing.T) {
	N := 10
	array := make([]int, 0)
	for i := 0; i < N; i++ {
		array = append(array, i)
	}
	h := Init(array, MIN)
	h.verify(t, 0)

	m := make(map[int]bool)
	for h.Len() > 0 {
		i := (h.Len() - 1) / 2
		x, err := h.Extract(i)
		if err == nil {
			m[x] = true
		} else {
			t.Errorf("Extract(%d) returns error: %s", i, err.Error())
			return
		}
		h.verify(t, 0)
	}
	if len(m) != N {
		t.Errorf("len(m) = %d; want %d", len(m), N)
	}
	for i := 0; i < len(m); i++ {
		if !m[i] {
			t.Errorf("m[%d] doesn't exist", i)
		}
	}
}

func TestMinReplace(t *testing.T) {
	array := make([]int, 0)
	for i := 200; i > 0; i -= 10 {
		array = append(array, i)
	}
	h := Init(array, MIN)
	h.verify(t, 0)

	x, err := h.Fetch(0)
	if err != nil {
		t.Errorf("Fetch returns error: %s", err.Error())
		return
	}
	if x != 10 {
		t.Errorf("Expected head to be 10, was %d", x)
		return
	}
	h.Replace(0, 210)
	h.verify(t, 0)

	for i := 100; i > 0; i-- {
		elem := rand.Intn(h.Len())
		if i&1 == 0 {
			if x, err := h.Fetch(elem); err != nil {
				t.Errorf("Fetch(%d) returns error: %s", i, err.Error())
				return
			} else if _, err := h.Replace(elem, x*2); err != nil {
				t.Errorf("Replace(%d) returns error: %s", i, err.Error())
				return
			}
		} else {
			if x, err := h.Fetch(elem); err != nil {
				t.Errorf("Fetch(%d) returns error: %s", i, err.Error())
				return
			} else if _, err := h.Replace(elem, x/2); err != nil {
				t.Errorf("Replace(%d) returns error: %s", i, err.Error())
				return
			}
		}
		h.verify(t, 0)
	}
}

func TestMaxPush(t *testing.T) {
	h := New(MAX)
	for i := 20; i > 10; i-- {
		h.Push(i)
		h.verify(t, 0)
	}

	for i := 10; i > 0; i-- {
		h.Push(i)
		h.verify(t, 0)
	}

	for i := 1; h.Len() > 0; i++ {
		if _, err := h.Pop(); err != nil {
			t.Errorf("%d.th pop returns error: %v", i, err.Error())
			return
		}
		h.verify(t, 0)
		if i < 20 {
			h.Push(20 + i)
		}
		h.verify(t, 0)
	}
}

func TestMaxInitSame(t *testing.T) {
	N := 20
	array := make([]int, 0)
	for i := N; i > 0; i-- {
		array = append(array, 0) // all elements are the same
	}
	h := Init(array, MAX)
	h.verify(t, 0)

	if h.Len() != N {
		t.Errorf("h.Len() = %d; want %d", h.Len(), N)
	}

	for i := 1; h.Len() > 0; i++ {
		x, err := h.Pop()
		if err != nil {
			t.Errorf("%d.th pop return error: %s", i, err.Error())
			return
		}
		h.verify(t, 0)
		if x != 0 {
			t.Errorf("%d.th pop got %d; want %d", i, x, 0)
		}
	}
}

func TestMaxInitDifferent(t *testing.T) {
	N := 20
	array := make([]int, 0)
	for i := 0; i < N; i++ {
		array = append(array, i) // all elements are different
	}
	h := Init(array, MAX)
	h.verify(t, 0)

	if h.Len() != N {
		t.Errorf("h.Len() = %d; want %d", h.Len(), N)
	}

	for i := h.Len() - 1; i > 0; i-- {
		x, err := h.Pop()
		if err != nil {
			t.Errorf("%d.th pop return error: %s", i, err.Error())
			return
		}
		h.verify(t, 0)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func TestMaxSort(t *testing.T) {
	slice := randomizer.GenerateIntSlice(100, 999)
	h := Init(slice, MAX)
	h.verify(t, 0)

	for i := 0; h.Len() > 0; i++ {
		var err error
		slice[i], err = h.Pop()
		if err != nil {
			t.Errorf("%d.th pop returns error: %v", i, err.Error())
			return
		}
		h.verify(t, 0)
	}

	if !sort.IsSorted(sort.Reverse(sort.IntSlice(slice))) {
		t.Errorf("Data not sortded %v", slice)
	}
}

func TestMaxExtractFirst(t *testing.T) {
	array := make([]int, 0)
	for i := 0; i < 10; i++ {
		array = append(array, i)
	}
	h := Init(array, MAX)
	h.verify(t, 0)

	for i := h.Len() - 1; i > 0; i-- {
		x, err := h.Extract(0)
		if err != nil {
			t.Errorf("Extract(0) returns error: %s", err.Error())
			return
		}
		if x != i {
			t.Errorf("Extract(0) got %d; want %d", x, i)
		}
		h.verify(t, 0)
	}
}

func TestMaxExtractLast(t *testing.T) {
	want := []int{9, 8, 6, 7, 4, 5, 2, 0, 3, 1}
	array := make([]int, 0)
	for i := 0; i < 10; i++ {
		array = append(array, i)
	}
	h := Init(array, MAX)
	h.verify(t, 0)

	for h.Len() > 0 {
		i := h.Len() - 1
		x, err := h.Extract(i)
		if err != nil {
			t.Errorf("Extract(%d) returns error: %s", i, err.Error())
			return
		}
		if len(want) > i {
			if x != want[i] {
				t.Errorf("Extract(%d) got %d; want %d", i, x, want[i])
			}
		} else {
			t.Errorf("Extract(%d) > len(want) = %d", i, len(want))
		}
		h.verify(t, 0)
	}
}

func TestMaxExtract(t *testing.T) {
	N := 10
	array := make([]int, 0)
	for i := 0; i < N; i++ {
		array = append(array, i)
	}
	h := Init(array, MAX)
	h.verify(t, 0)

	m := make(map[int]bool)
	for h.Len() > 0 {
		i := (h.Len() - 1) / 2
		x, err := h.Extract(i)
		if err == nil {
			m[x] = true
		} else {
			t.Errorf("extract %d error: %s", i, err.Error())
		}
		h.verify(t, 0)
	}
	if len(m) != N {
		t.Errorf("len(m) = %d; want %d", len(m), N)
	}
	for i := 0; i < len(m); i++ {
		if !m[i] {
			t.Errorf("m[%d] doesn't exist", i)
		}
	}
}

func TestMaxReplace(t *testing.T) {
	array := make([]int, 0)
	for i := 200; i > 0; i -= 10 {
		array = append(array, i)
	}
	h := Init(array, MAX)
	h.verify(t, 0)

	x, err := h.Fetch(0)
	if err != nil {
		t.Errorf("Fetch returns error: %s", err.Error())
		return
	}
	if x != 200 {
		t.Errorf("Expected head to be 200, was %d", x)
		return
	}
	h.Replace(19, 210)
	h.verify(t, 0)

	for i := 100; i > 0; i-- {
		elem := rand.Intn(h.Len())
		if i&1 == 0 {
			if x, err := h.Fetch(elem); err != nil {
				t.Errorf("Fetch(%d) returns error: %s", i, err.Error())
				return
			} else if _, err := h.Replace(elem, x*2); err != nil {
				t.Errorf("Replace(%d) returns error: %s", i, err.Error())
				return
			}
		} else {
			if x, err := h.Fetch(elem); err != nil {
				t.Errorf("Fetch(%d) returns error: %s", i, err.Error())
				return
			} else if _, err := h.Replace(elem, x/2); err != nil {
				t.Errorf("Replace(%d) returns error: %s", i, err.Error())
				return
			}
		}
		h.verify(t, 0)
	}
}

func benchmarkInit(b *testing.B, size int) {
	b.StopTimer()
	slice := randomizer.GenerateIntSlice(size, 999)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = Init(slice, MIN)
	}
}

func benchmarkPush(b *testing.B, h *Heap, size int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			h.Push(j)
		}
	}
}

func benchmarkPop(b *testing.B, size int) {
	b.StopTimer()
	slice := randomizer.GenerateIntSlice(size, 999)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		h := Init(slice, MIN)
		b.StartTimer()
		for h.Len() > 0 {
			h.Pop()
		}
		h = nil
	}
}

func benchmarkExtract(b *testing.B, size int) {
	b.StopTimer()
	slice := randomizer.GenerateIntSlice(size, 999)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		h := Init(slice, MIN)
		b.StartTimer()
		for h.Len() > 0 {
			j := (h.Len() - 1) / 2
			if _, err := h.Extract(j); err != nil {
				b.Fatalf("Extract(%d) in %d.th iteraton get error: %s", j, i, err.Error())
			}
		}
		h = nil
	}
}

func benchmarkReplace(b *testing.B, size int) {
	b.StopTimer()
	slice := randomizer.GenerateIntSlice(size, 999)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		h := Init(slice, MIN)
		b.StartTimer()
		for c := h.Len(); c > 0; c-- {
			j := (h.Len() - 1) / 2
			if x, err := h.Fetch(j); err != nil {
				b.Fatalf("Fetch(%d) in %d.th iteraton get error: %s", j, i, err.Error())
			} else if _, err := h.Replace(j, x*2); err != nil {
				b.Fatalf("Replace(%d) in %d.th iteraton get error: %s", j, i, err.Error())
			}
		}
		h = nil
	}
}

func BenchmarkInit100(b *testing.B) {
	benchmarkInit(b, 100)
}

func BenchmarkInit100000(b *testing.B) {
	benchmarkInit(b, 100000)
}

func BenchmarkPush100(b *testing.B) {
	b.StopTimer()
	size := 100
	heap := New(MIN)
	b.StartTimer()
	benchmarkPush(b, heap, size)
}

func BenchmarkPush100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	heap := New(MIN)
	b.StartTimer()
	benchmarkPush(b, heap, size)
}

func BenchmarkPop100(b *testing.B) {
	benchmarkPop(b, 100)
}

func BenchmarkPop100000(b *testing.B) {
	benchmarkPop(b, 100000)
}

func BenchmarkExtract100(b *testing.B) {
	benchmarkExtract(b, 100)
}

func BenchmarkExtract100000(b *testing.B) {
	benchmarkExtract(b, 100000)
}

func BenchmarkReplace100(b *testing.B) {
	benchmarkReplace(b, 100)
}

func BenchmarkReplace100000(b *testing.B) {
	benchmarkReplace(b, 100000)
}
