package randomizer

import "testing"

func TestGenerateIntSlice(t *testing.T) {
	count, n := 10, 100
	slice := GenerateIntSlice(count, n)

	if len(slice) != count {
		t.Errorf("len(slice) = %d; expected %d", len(slice), count)
	}

	for i, v := range slice {
		if v > n {
			t.Errorf("item %d invariant invalidated %d > %d", i, v, n)
		}
	}
}
