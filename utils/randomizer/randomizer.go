package randomizer

import "math/rand"

func GenerateIntSlice(count, n int) (slice []int) {
	for i := 0; i < count; i++ {
		slice = append(slice, rand.Intn(n))
	}
	return
}
