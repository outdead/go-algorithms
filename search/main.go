package main

import (
	"fmt"
	"sort"
)

func main() {
	var initial = []int{1, 2, 5, 5, 12, 13, 13, 17, 21, 21, 24, 25, 25, 44, 45, 47, 56, 66, 75, 84, 85, 92, 97, 99, 99}
	var needle = 25
	var want = 11 // 12 if unstable

	fmt.Printf("Initial array .......... %v\n", initial)
	fmt.Printf("Needle value ........... %v\n", needle)
	fmt.Printf("Want index ............. %v\n\n", want)

	fmt.Printf("Native SearchInts ...... %v\n", sort.SearchInts(initial, needle))
	fmt.Printf("LineSearch ............. %v\n", LineSearch(initial, needle))
	fmt.Printf("LinePredicateSearch .... %v\n", LinePredicateSearch(initial, needle))
	fmt.Printf("BinarySearch ........... %v\n", BinarySearch(initial, needle))
	fmt.Printf("BinaryRecursionSearch .. %v\n", BinaryRecursionSearch(initial, needle, 0, len(initial)))
	fmt.Printf("InterpolationSearch .... %v\n", InterpolationSearch(initial, needle))

	hashMap := make(map[int]int, len(initial))
	for key, value := range initial {
		hashMap[value] = key
	}
	index, ok := hashMap[needle];
	if !ok {
		index = -1
	}
	fmt.Printf("hashMap ................ %v\n", index)
}

func LineSearch(array []int, needle int) int {
	for key, value := range array {
		if value == needle {
			return key
		}
	}
	return -1
	////array[len(array)+1] = needle
	//var i int
	//for i = 0; array[i] != needle; i++ {}
	//return i
}

func LinePredicateSearch(array []int, needle int) int {
	return linePredicateSearch(len(array), func(i int) bool {
		return array[i] == needle
	})
}

func linePredicateSearch(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func BinarySearch(array []int, needle int) int {
	left := 0
	right := len(array)

	if right == 0 || needle < array[0] || needle > array[right - 1] {
		return -1
	}

	for left < right {
		mid := left + (right - left) / 2

		if needle <= array[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if array[right] == needle {
		return right
	} else {
		return -1
	}
}

func BinaryRecursionSearch(array []int, needle int, left int, right int) int {
	if left >= right {
		if array[left] == needle {
			return left
		}
		return -1
	}

	mid := left + (right - left) / 2

	if array[mid] == needle {
		return mid
	} else if array[mid] > needle {
		return BinaryRecursionSearch(array, needle, left, mid)
	} else {
		return BinaryRecursionSearch(array, needle, mid, right)
	}
}

func InterpolationSearch(array []int, needle int) int {
	left := 0
	right := len(array) - 1
	mid := 0

	if right == 0 || needle < array[0] || needle > array[right] {
		return -1
	}

	for array[left] < needle && array[right] > needle {
		mid = left + ((needle - array[left]) * (right - left)) / (array[right] - array[left]);

		if (array[mid] < needle) {
			left = mid + 1;
		} else if (array[mid] > needle) {
			right = mid - 1;
		} else {
			return mid;
		}
	}

	if (array[left] == needle) {
		return left;
	} else if (array[right] == needle) {
		return right;
	} else {
		return -1; // Not found
	}
}
