package main

import (
	"runtime"
	"testing"
	"sort"
	"fmt"
)

const N = 10000

//var initial = []int{25, 13, 99, 84, 99, 5, 21, 17, 12, 97, 45, 44, 1, 92, 25, 13, 5, 85, 2, 47, 56, 21, 66, 24, 75}
var initial = []int{
	273, 217, 481, 821, 502, 254, 4, 233, 674, 181, 733, 563, 571, 447, 772, 125, 29, 511, 354, 656, 216, 621, 229, 995, 95,
	25, 13, 99, 84, 99, 5, 21, 25, 12, 99, 45, 44, 1, 92, 25, 13, 5, 85, 2, 47, 56, 21, 666, 24, 75,
	679, 930, 500, 68, 486, 550, 707, 592, 142, 312, 131, 495, 468, 88, 160, 155, 274, 470, 262, 490, 54, 357, 937, 849, 919,
	732, 499, 525, 504, 412, 712, 119, 90, 142, 634, 534, 594, 188, 474, 854, 630, 359, 464, 525, 20, 401, 286, 328, 266, 545,
}

func init() {
	runtime.GOMAXPROCS(1)
}

// *** tests

func TestNativeSortInts(t *testing.T) {
	array := make([]int, len(initial))
	copy(array, initial)

	sort.Ints(array)

	if !sort.IsSorted(sort.IntSlice(array)) {
		t.Fatalf("Slice is not sorted %v", array)
	}
	fmt.Println("TestNativeSortInts .... OK")
}

func TestNativeQuickSort(t *testing.T) {
	array := make([]int, 100)
	copy(array, initial)

	sort.Sort(sort.IntSlice(array))

	if !sort.IsSorted(sort.IntSlice(array)) {
		t.Fatalf("Slice is not sorted %v", array)
	}
	fmt.Println("TestNativeQuickSort ... OK")
}

func TestBubbleSort(t *testing.T) {
	array := make([]int, len(initial))
	copy(array, initial)

	BubbleSort(array)

	if !sort.IsSorted(sort.IntSlice(array)) {
		t.Fatalf("Slice is not sorted %v", array)
	}
	fmt.Println("TestBubbleSort ........ OK")
}

func TestInsertionSort(t *testing.T) {
	array := make([]int, len(initial))
	copy(array, initial)

	InsertionSort(array)

	if !sort.IsSorted(sort.IntSlice(array)) {
		t.Fatalf("Slice is not sorted %v", array)
	}
	fmt.Println("TestInsertionSort ..... OK")
}

func TestShellSort(t *testing.T) {
	array := make([]int, len(initial))
	copy(array, initial)

	ShellSort(array)

	if !sort.IsSorted(sort.IntSlice(array)) {
		t.Fatalf("Slice is not sorted %v", array)
	}
	fmt.Println("TestShellSort ......... OK")
}

func TestSelectionSort(t *testing.T) {
	array := make([]int, len(initial))
	copy(array, initial)

	SelectionSort(array)

	if !sort.IsSorted(sort.IntSlice(array)) {
		t.Fatalf("Slice is not sorted %v", array)
	}
	fmt.Println("TestSelectionSort ..... OK")
}

func TestMergeSort(t *testing.T) {
	array := make([]int, len(initial))
	copy(array, initial)

	MergeSort(array)

	if !sort.IsSorted(sort.IntSlice(array)) {
		t.Fatalf("Slice is not sorted %v", array)
	}
	fmt.Println("TestMergeSort ......... OK")
}

func TestMergeBubbleSort(t *testing.T) {
	array := make([]int, len(initial))
	copy(array, initial)

	MergeBubbleSort(array)

	if !sort.IsSorted(sort.IntSlice(array)) {
		t.Fatalf("Slice is not sorted %v", array)
	}
	fmt.Println("TestMergeBubbleSort ... OK")
}

func TestMergeQuickSort(t *testing.T) {
	array := make([]int, len(initial))
	copy(array, initial)

	MergeQuickSort(array)

	if !sort.IsSorted(sort.IntSlice(array)) {
		t.Fatalf("Slice is not sorted %v", array)
	}
	fmt.Println("TestMergeQuickSort .... OK")
}

func TestQuickSort(t *testing.T) {
	array := make([]int, len(initial))
	copy(array, initial)

	QuickSort(array)

	if !sort.IsSorted(sort.IntSlice(array)) {
		t.Fatalf("Slice is not sorted %v", array)
	}
	fmt.Println("TestQuickSort ......... OK")
}

func TestCountingSort(t *testing.T) {
	array := make([]int, len(initial))
	copy(array, initial)

	CountingSort(array, 0, 1000)

	if !sort.IsSorted(sort.IntSlice(array)) {
		t.Fatalf("Slice is not sorted %v", array)
	}
	fmt.Println("TestCountingSort ...... OK")
}

func TestRadixSort(t *testing.T) {
	var array = []int32{
		273, 217, 481, 821, 502, 254, 4, 233, 674, 181, 733, 563, 571, 447, 772, 125, 29, 511, 354, 656, 216, 621, 229, 995, 95,
		25, 13, 99, 84, 99, 5, 21, 25, 12, 99, 45, 44, 1, 92, 25, 13, 5, 85, 2, 47, 56, 21, 666, 24, 75,
		679, 930, 500, 68, 486, 550, 707, 592, 142, 312, 131, 495, 468, 88, 160, 155, 274, 470, 262, 490, 54, 357, 937, 849, 919,
		732, 499, 525, 504, 412, 712, 119, 90, 142, 634, 534, 594, 188, 474, 854, 630, 359, 464, 525, 20, 401, 286, 328, 266, 545,
	}

	RadixSort(array, 4, 10)

	array2 := make([]int, 100)
	for k, v := range array {
		array2[k] = int(v)
	}

	if !sort.IsSorted(sort.IntSlice(array2)) {
		t.Fatalf("Slice is not sorted %v", array)
	}
	fmt.Println("TestRadixSort ......... OK")
}

// *** benchmarks

func BenchmarkNativeSortInts(b *testing.B) {
	array := make([]int, len(initial))
	b.N = N

	for i := 1; i <= b.N; i++ {
		copy(array, initial)
		sort.Ints(array)
	}
}

func BenchmarkNativeQuickSort(b *testing.B) {
	array := make([]int, len(initial))
	b.N = N

	for i := 1; i <= b.N; i++ {
		copy(array, initial)
		sort.Sort(sort.IntSlice(array))
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	array := make([]int, len(initial))
	b.N = N

	for i := 1; i <= b.N; i++ {
		copy(array, initial)
		BubbleSort(array)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	array := make([]int, len(initial))
	b.N = N

	for i := 1; i <= b.N; i++ {
		copy(array, initial)
		InsertionSort(array)
	}
}

func BenchmarkShellSort(b *testing.B) {
	array := make([]int, len(initial))
	b.N = N

	for i := 1; i <= b.N; i++ {
		copy(array, initial)
		ShellSort(array)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	array := make([]int, len(initial))
	b.N = N

	for i := 1; i <= b.N; i++ {
		copy(array, initial)
		SelectionSort(array)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	array := make([]int, len(initial))
	b.N = N

	for i := 1; i <= b.N; i++ {
		copy(array, initial)
		MergeSort(array)
	}
}

func BenchmarkMergeBubbleSort(b *testing.B) {
	array := make([]int, len(initial))
	b.N = N

	for i := 1; i <= b.N; i++ {
		copy(array, initial)
		MergeBubbleSort(array)
	}
}

func BenchmarkMergeQuickSort(b *testing.B) {
	array := make([]int, len(initial))
	b.N = N

	for i := 1; i <= b.N; i++ {
		copy(array, initial)
		MergeQuickSort(array)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	array := make([]int, len(initial))
	b.N = N

	for i := 1; i <= b.N; i++ {
		copy(array, initial)
		QuickSort(array)
	}
}

func BenchmarkCountingSort(b *testing.B) {
	array := make([]int, len(initial))
	b.N = N

	for i := 1; i <= b.N; i++ {
		copy(array, initial)
		CountingSort(array, 0, 1000)
	}
}

func BenchmarkRadixSort(b *testing.B) {
	b.N = N

	for i := 1; i <= b.N; i++ {
		var array = []int32{
			273, 217, 481, 821, 502, 254, 4, 233, 674, 181, 733, 563, 571, 447, 772, 125, 29, 511, 354, 656, 216, 621, 229, 995, 95,
			25, 13, 99, 84, 99, 5, 21, 25, 12, 99, 45, 44, 1, 92, 25, 13, 5, 85, 2, 47, 56, 21, 666, 24, 75,
			679, 930, 500, 68, 486, 550, 707, 592, 142, 312, 131, 495, 468, 88, 160, 155, 274, 470, 262, 490, 54, 357, 937, 849, 919,
			732, 499, 525, 504, 412, 712, 119, 90, 142, 634, 534, 594, 188, 474, 854, 630, 359, 464, 525, 20, 401, 286, 328, 266, 545,
		}
		RadixSort(array, 4, 10)
	}
}
