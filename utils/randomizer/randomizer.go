package randomizer

import (
	"math/rand"
)

func Intnk(min, max int) int {
	return rand.Intn(max-min) + min
}

func GenerateIntSlice(length int, maxValue int) (slice []int) {
	for i := 0; i < length; i++ {
		slice = append(slice, rand.Intn(maxValue))
	}
	return
}

func GenerateSortedIntSlice(length int) (slice []int) {
	min, max := 0, 0
	for i := 0; i < length; i++ {
		max = min + 10
		min = Intnk(min, max)
		slice = append(slice, min)
	}
	return
}

func GenerateInt32Slice(length int, maxValue int32) (slice []int32) {
	for i := 0; i < length; i++ {
		slice = append(slice, rand.Int31n(maxValue))
	}
	return
}
