package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Test slice sum", func(t *testing.T) {

		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got '%d' wanted '%d' given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Test slices sum", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 4, 5})
		want := []int{3, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%v is not expected %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Test sum of tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{0, 1}, []int{2, 9, 1})
		want := []int{1, 10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%v is not expected %v", got, want)
		}
	})
}
