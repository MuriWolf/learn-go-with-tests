package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	checkSlice := func(t *testing.T, expected, got []int) {
		t.Helper()
		if !slices.Equal(expected, got) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	}

	checkSum := func(t *testing.T, expected, got int) {
		t.Helper()
		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	}

	t.Run("sum array (fixed positions)", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		sum := Sum(numbers[:])
		expected := 15

		checkSum(t, expected, sum)
	})

	t.Run("sum slice (dynamic positions)", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		expected := 6

		checkSum(t, expected, sum)
	})

	t.Run("sum all", func(t *testing.T) {
		sums := SumAll([]int{1, 2}, []int{3, 4})
		expected := []int{3, 7}

		checkSlice(t, expected, sums)
	})

	t.Run("sum all tails", func(t *testing.T) {
		sums := SumAllTails([]int{2, 6}, []int{3, 7, 2})
		expected := []int{6, 9}

		checkSlice(t, expected, sums)
	})

	t.Run("sum all tails (inclusing empty slices)", func(t *testing.T) {
		sums := SumAllTails([]int{2, 6}, []int{3, 7, 2}, []int{})
		expected := []int{6, 9, 0}

		checkSlice(t, expected, sums)
	})
}
