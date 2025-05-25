package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		// Run individual tests: go test -run TestArea/{testName}
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 5, Height: 7}, hasArea: 17.5},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) { // tt.name is used as the test case name
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.name, got, tt.hasArea)
			}
		})
	}
	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }
	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	checkArea(t, rectangle, 100.0)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10.0}
	// 	checkArea(t, circle, 314.1592653589793)
	// })
}
