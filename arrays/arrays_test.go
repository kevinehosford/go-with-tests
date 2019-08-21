package arrays

import "testing"

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
