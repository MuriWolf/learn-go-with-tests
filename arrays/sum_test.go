package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum array (fixed positions)", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		sum := Sum(numbers[:])
		expected := 15

		assertIsSumEqual(t, expected, sum)
	})

	t.Run("sum slice (dynamic positions)", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		expected := 6

		assertIsSumEqual(t, expected, sum)
	})

	t.Run("sum all", func(t *testing.T) {
		sums := SumAll([]int{1, 2}, []int{3, 4})
		expected := []int{3, 7}

		assertIsSliceEqual(t, expected, sums)
	})
}

func assertIsSumEqual(t *testing.T, expected, got int) {
	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func assertIsSliceEqual(t *testing.T, expected, got []int) {
	if !slices.Equal(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
