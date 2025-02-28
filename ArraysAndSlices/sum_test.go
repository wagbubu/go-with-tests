package main

import (
	"slices"
	"testing"
)

func TestSUM(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, given, %v", got, want, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	assertCorrectSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("sum of 2 slices", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{0, 9}
		got := SumAllTails(numbers1, numbers2)
		want := []int{5, 9}
		assertCorrectSum(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		var numbers1 []int
		numbers2 := []int{0, 9}
		got := SumAllTails(numbers1, numbers2)
		want := []int{0, 9}
		assertCorrectSum(t, got, want)
	})
}
