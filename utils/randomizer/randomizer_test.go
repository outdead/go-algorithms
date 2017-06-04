package randomizer

import "testing"

func TestGenerateIntSlice(t *testing.T) {
	length, maxValue := 10, 100
	slice := GenerateIntSlice(length, maxValue)

	if len(slice) != length {
		t.Errorf("len(slice) = %d; expected %d", len(slice), length)
	}

	for i, v := range slice {
		if v > maxValue {
			t.Errorf("item %d invariant invalidated %d > %d", i, v, maxValue)
		}
	}
}

func TestGenerateInt32Slice(t *testing.T) {
	length, maxValue := 10, int32(100)
	slice := GenerateInt32Slice(length, maxValue)

	if len(slice) != length {
		t.Errorf("len(slice) = %d; expected %d", len(slice), length)
	}

	for i, v := range slice {
		if v > maxValue {
			t.Errorf("item %d invariant invalidated %d > %d", i, v, maxValue)
		}
	}
}
