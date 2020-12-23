package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %v want %v given, %v", got, want, numbers)
		}
	})

}

func checkSums(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestSumAll(t *testing.T) {
	t.Run("two slices of two numbers", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("two slices of two numbers", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("A slice of one number ", func(t *testing.T) {
		got := SumAllTails([]int{9})
		want := []int{0}

		checkSums(t, got, want)
	})

	t.Run("An empty slice", func(t *testing.T) {
		got := SumAllTails([]int{})
		want := []int{0}

		checkSums(t, got, want)
	})
}
