package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Add(2,2)", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("Expected '%d', but got '%d'", expected, sum)
		}
	})
	t.Run("Add(2,-2)", func(t *testing.T) {
		sum := Add(2, -2)
		expected := 0

		if sum != expected {
			t.Errorf("Expected '%d', but got '%d'", expected, sum)
		}
	})

}

func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Println(sum)
	// Output: 4
}
