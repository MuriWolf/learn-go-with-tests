package structs

import "testing"

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		perimeter := shape.Perimeter()
		assertFloat(t, want, perimeter)
	}

	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{width: 10.0, height: 10.0}, 40.0},
		{Circle{radius: 10}, 62.83185307179586},
		{Triangle{height: 10, base: 10, sideA: 10, sideB: 10}, 30.0},
	}

	for _, tt := range perimeterTests {
		checkPerimeter(t, tt.shape, tt.want)
	}

}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		area := shape.Area()
		assertFloat(t, want, area)
	}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{width: 10.0, height: 10.0}, 100.0},
		{Circle{radius: 10}, 314.1592653589793},
		{Triangle{height: 10, base: 10, sideA: 10, sideB: 10}, 50.0},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}

}

func assertFloat(t *testing.T, expected, got float64) {
	t.Helper()
	if expected != got {
		t.Errorf("Expected %g got %g", expected, got)
	}
}
