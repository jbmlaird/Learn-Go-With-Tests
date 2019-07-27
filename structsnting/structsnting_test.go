package structsnting

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("test perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("Expected %.2f, got %.2f", want, got)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("correctly calc Rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("Expected %.2f, got %.2f", want, got)
		}
	})

	t.Run("correctly calc Circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}
