package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Get perimeter of a retangle", func(t *testing.T) {
		retangle := Rectangle{10.0, 10.0}
		perimeter := retangle.Perimeter()
		expected := 40.0
		assertFloat(t, expected, perimeter)
	})

	t.Run("Get perimeter of a circle", func(t *testing.T) {
		circle := Circle{10}
		perimeter := circle.Perimeter()
		expected := 62.83185307179586
		assertFloat(t, expected, perimeter)
	})
}

func TestArea(t *testing.T) {
	t.Run("Get area of a retangle", func(t *testing.T) {
		retangle := Rectangle{10.0, 10.0}
		area := retangle.Area()
		expected := 100.0
		assertFloat(t, expected, area)
	})

	t.Run("Get area of a circle", func(t *testing.T) {
		circle := Circle{10}
		area := circle.Area()
		expected := 314.1592653589793
		assertFloat(t, expected, area)
	})

}

func assertFloat(t *testing.T, expected, got float64) {
	t.Helper()
	if expected != got {
		t.Errorf("Expected %g got %g", expected, got)
	}
}
