package randomizer

import (
	"sort"
	"testing"
)

func TestGenerateIntSlice(t *testing.T) {
	length, maxValue := 10, 100
	data := GenerateIntSlice(length, maxValue)

	if len(data) != length {
		t.Errorf("len(data) = %d; expected %d", len(data), length)
	}

	for i, v := range data {
		if v > maxValue {
			t.Errorf("item %d invariant invalidated %d > %d", i, v, maxValue)
		}
	}
}

func TestGenerateSortedIntSlice(t *testing.T) {
	length := 1000
	data := GenerateSortedIntSlice(length)
	if len(data) != length {
		t.Errorf("len(data) = %d; expected %d", len(data), length)
	}

	if !sort.IntsAreSorted(data) {
		t.Error("Data in not sorted")
	}
}

func TestGenerateInt32Slice(t *testing.T) {
	length, maxValue := 10, int32(100)
	data := GenerateInt32Slice(length, maxValue)

	if len(data) != length {
		t.Errorf("len(data) = %d; expected %d", len(data), length)
	}

	for i, v := range data {
		if v > maxValue {
			t.Errorf("item %d invariant invalidated %d > %d", i, v, maxValue)
		}
	}
}
