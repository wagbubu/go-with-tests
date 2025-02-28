package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{width: 12.0, height: 6.0}, hasArea: 72.0},
		{shape: Circle{radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{base: 12, height: 6}, hasArea: 36.0},
	}

	checkArea := func(t testing.TB, shape Shape, hasArea float64) {
		t.Helper()
		got := shape.Area()
		if got != hasArea {
			t.Errorf("%#v got %.2f, has area %.2f", shape, got, hasArea)
		}
	}

	for _, test := range areaTests {
		checkArea(t, test.shape, test.hasArea)
	}
}
