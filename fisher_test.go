package stats

import (
	"reflect"
	"testing"
)

func TestFisherYatesShuffleWithExclusion(t *testing.T) {
	input := []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []any{6, 1, 2, 8, 0, 5, 4, 7, 9, 3}
	opts := &ShuffleOptions{Seed: int64(3), ExcludeIndices: map[int]bool{1: true, 2: true, 7: true}}

	// Copying original input
	shuffled := make([]any, len(input))
	copy(shuffled, input)

	FisherYatesShuffle(shuffled, opts)

	// Check if the result matches with the expected one
	if !reflect.DeepEqual(shuffled, expected) {
		t.Errorf("Expected %v, got %v", expected, shuffled)
	}

	// Checking that exclusions has not moved
	for i := range opts.ExcludeIndices {
		if shuffled[i] != input[i] {
			t.Errorf("Excluded index %d was altered: expected %v, got %v", i, input[i], shuffled[i])
		}
	}

	// Checking that all elements are conserved
	originalMap := make(map[any]int)
	shuffledMap := make(map[any]int)

	for _, in := range input {
		originalMap[in]++
	}

	for _, sh := range shuffled {
		shuffledMap[sh]++
	}

	if !reflect.DeepEqual(originalMap, shuffledMap) {
		t.Errorf("Mismatch in elements: original %v, shuffled %v", originalMap, shuffledMap)
	}
}

func TestFisherYatesShuffle(t *testing.T) {
	input := []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	opts := &ShuffleOptions{Seed: int64(3)}
	expected := []any{9, 8, 5, 0, 1, 6, 3, 2, 7, 4}

	// Copying original input
	shuffled := make([]any, len(input))
	copy(shuffled, input)

	FisherYatesShuffle(shuffled, opts)

	// Check if the result matches with the expected one
	if !reflect.DeepEqual(shuffled, expected) {
		t.Errorf("Expected %v, got %v", expected, shuffled)
	}

	// Checking that all elements are conserved
	originalMap := make(map[any]int)
	shuffledMap := make(map[any]int)

	for _, in := range input {
		originalMap[in]++
	}

	for _, sh := range shuffled {
		shuffledMap[sh]++
	}

	if !reflect.DeepEqual(originalMap, shuffledMap) {
		t.Errorf("Mismatch in elements: original %v, shuffled %v", originalMap, shuffledMap)
	}
}
