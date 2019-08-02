package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("sum slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if want != got {
			t.Errorf("Wanted '%d' but got '%d', given '%v'", want, got, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wanted '%v' but got '%v'", want, got)
		}
	}

	t.Run("sum all the tails correclty", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 5})
		want := []int{5, 14}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
