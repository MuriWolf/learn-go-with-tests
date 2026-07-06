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
}

func assertIsSumEqual(t *testing.T, expected, got int) {
	if got != expected {
		t.Errorf("expected %d, got %d", got, expected)
	}
}
