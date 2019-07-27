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
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		// I could rewrite the below lines as:
		// shape: Rectangle{Width: 10.0, Height: 10.0}, want: 100.0}, etc.
		// but GoLand provides this automatically
		{Rectangle{10.0, 10.0}, 100.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}

}
