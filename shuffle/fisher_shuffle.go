package shuffle

import (
	"math/rand"
	"time"
)

/*
ShuffleOptions to set special features:
- Seed: seed for random number generator
- ExcludeIndices: Slice of indices which are not desired to be shuffled
*/
type ShuffleOptions struct {
	Seed           int64
	ExcludeIndices []int
}

/*
Fisher Yates Shuffling method
*/
func FisherYatesShuffle[T any](arr []T, opts *ShuffleOptions) {
	seed := time.Now().UnixNano()
	excludeIndices := make(map[int]bool)

	if opts != nil {
		if opts.Seed != 0 {
			seed = opts.Seed
		}

		if opts.ExcludeIndices != nil {
			for _, ei := range opts.ExcludeIndices {
				excludeIndices[ei] = true
			}
		}
	}

	r := rand.New(rand.NewSource(seed))
	n := len(arr)

	validIndices := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if len(excludeIndices) == 0 || !excludeIndices[i] {
			validIndices = append(validIndices, i)
		}
	}

	for i := n - 1; i > 0; i-- {
		if len(excludeIndices) == 0 || !excludeIndices[i] {
			j := validIndices[r.Intn(len(validIndices))]
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
}
