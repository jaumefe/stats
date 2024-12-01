package stats

import (
	"math/rand"
	"time"
)

func FihserYatesShuffleWithExclusion(arr []int, excludeIndices map[int]bool) {
	rand.NewSource(time.Now().UnixNano())
	n := len(arr)

	validIndices := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if !excludeIndices[i] {
			validIndices = append(validIndices, i)
		}
	}

	for i := n - 1; i > 0; i-- {
		j := validIndices[rand.Intn(len(validIndices))]
		arr[i], arr[j] = arr[j], arr[i]
	}
}
