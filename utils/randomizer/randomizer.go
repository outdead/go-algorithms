package randomizer

import "math/rand"

func GenerateIntSlice(count, n int) (slice []int) {
	for i := 0; i < count; i++ {
		slice = append(slice, rand.Intn(n))
	}
	return
}

func GenerateInt32Slice(count int, n int32) (slice []int32) {
	for i := 0; i < count; i++ {
		slice = append(slice, rand.Int31n(n))
	}
	return
}
