package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 12.0}
	got := rectangle.Perimeter()
	want := 44.0

	if got != want {
		t.Errorf("wanted %.f go %.f", want, got)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			t.Errorf("wanted %v got %v", want, got)
		}
	}

	tests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 12.0}, 120.0},
		{Circle{10.0}, 314.16},
		{Triangle{12.5, 10.0}, 62.5},
	}

	for _, test := range tests {
		checkArea(t, test.shape, test.want)
	}
}
