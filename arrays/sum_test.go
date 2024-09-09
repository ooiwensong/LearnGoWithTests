package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	arr1 := []int{1, 2}
	arr2 := []int{0, 9}

	got := SumAll(arr1, arr2)
	want := []int{3, 9}

	// DeepEqual compares if two variables of same type are deeply equal.
	// In the case of arrays/ slices, they are deeply equal if their corresponding
	// elements are deeply equal.
	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("got %v want %v, given %v and %v", got, want, arr1, arr2)
	// }

	// slices.Equal do a simple shallow compare of slices
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v, given %v and %v", got, want, arr1, arr2)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		arr1 := []int{1, 2}
		arr2 := []int{0, 9}

		got := SumAllTails(arr1, arr2)
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
