package number

import (
	"math/rand"
	"time"
)

func ArrayShuffler(array *[25]int) [25]int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(array),
		func(i, j int) {
			array[i], array[j] = array[j], array[i]
		})
	return *array
}
