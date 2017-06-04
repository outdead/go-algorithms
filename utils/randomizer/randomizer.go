package randomizer

import "math/rand"

func GenerateIntSlice(length int, maxValue int) (slice []int) {
	for i := 0; i < length; i++ {
		slice = append(slice, rand.Intn(maxValue))
	}
	return
}

func GenerateSortedIntSlice(length int) (slice []int) {
	for i := 0; i < length; i++ {
		slice = append(slice, i)
	}
	return
}

func GenerateInt32Slice(length int, maxValue int32) (slice []int32) {
	for i := 0; i < length; i++ {
		slice = append(slice, rand.Int31n(maxValue))
	}
	return
}
