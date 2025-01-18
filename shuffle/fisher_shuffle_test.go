package shuffle

import (
	"reflect"
	"testing"
)

func TestFisherYatesShuffleWithExclusion(t *testing.T) {
	input := []any{0, "foo", 2, 3.2, 4, true, 6, 7, 8, "bar"}
	expected := []any{6, "foo", 2, 8, 0, true, 4, 7, "bar", 3.2}
	opts := &ShuffleOptions{Seed: int64(3), ExcludeIndices: []int{1, 2, 7}}

	// Copying original input
	shuffled := make([]any, len(input))
	copy(shuffled, input)

	FisherYatesShuffle(shuffled, opts)

	// Check if the result matches with the expected one
	if !reflect.DeepEqual(shuffled, expected) {
		t.Errorf("Expected %v, got %v", expected, shuffled)
	}

	// Checking that exclusions has not moved
	for _, ei := range opts.ExcludeIndices {
		if shuffled[ei] != input[ei] {
			t.Errorf("Excluded index %d was altered: expected %v, got %v", ei, input[ei], shuffled[ei])
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
	input := []any{0, "foo", 2, 3, 4, true, 6, 7.2, 8, "bar"}
	opts := &ShuffleOptions{Seed: int64(3)}
	expected := []any{"bar", 8, true, 0, "foo", 6, 3, 2, 7.2, 4}

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

func TestFisherYatesShuffleEmptySlice(t *testing.T) {
	input := []any{}
	opts := &ShuffleOptions{Seed: int64(3)}
	expected := []any{}

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
