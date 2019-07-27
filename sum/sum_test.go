package sum

import "testing"

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
