package stats

import (
	"math/rand"
)

func FisherYatesShuffleWithExclusion(arr []any, excludeIndices map[int]bool, seed int64) {
	r := rand.New(rand.NewSource(seed))
	n := len(arr)

	validIndices := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if !excludeIndices[i] {
			validIndices = append(validIndices, i)
		}
	}

	for i := n - 1; i > 0; i-- {
		if !excludeIndices[i] {
			j := validIndices[r.Intn(len(validIndices))]
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
}
