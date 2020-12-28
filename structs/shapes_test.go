package shapes

import "testing"

func TestAreaAndPerimeter(t *testing.T) {

	areaAndPerimeterTests := []struct {
		name         string
		shape        Shape
		hasArea      float64
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0, hasPerimeter: 36.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793, hasPerimeter: 62.83185307179586},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0, hasPerimeter: 0.0},
	}

	for _, tt := range areaAndPerimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			gotArea := tt.shape.Area()
			gotPerimeter := tt.shape.Perimeter()
			if gotArea != tt.hasArea || gotPerimeter != tt.hasPerimeter {
				t.Errorf("%#v got %g and %g want %g and %g", tt.shape, gotArea, gotPerimeter, tt.hasArea, tt.hasPerimeter)
			}
		})
	}
}
