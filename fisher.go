package stats

import (
	"math/rand"
	"time"
)

type ShuffleOptions struct {
	Seed           int64
	ExcludeIndices map[int]bool
}

func FisherYatesShuffle(arr []any, opts *ShuffleOptions) {
	seed := time.Now().UnixNano()
	var excludeIndices map[int]bool

	if opts != nil {
		if opts.Seed != 0 {
			seed = opts.Seed
		}

		if opts.ExcludeIndices != nil {
			excludeIndices = opts.ExcludeIndices
		}
	}

	r := rand.New(rand.NewSource(seed))
	n := len(arr)

	validIndices := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if excludeIndices == nil || !excludeIndices[i] {
			validIndices = append(validIndices, i)
		}
	}

	for i := n - 1; i > 0; i-- {
		if excludeIndices == nil || !excludeIndices[i] {
			j := validIndices[r.Intn(len(validIndices))]
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
}
