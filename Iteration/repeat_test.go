package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat five times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected %q, but got %q", expected, repeated)
		}
	})
	t.Run("repeat six times", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"
		if repeated != expected {
			t.Errorf("expected %q, but got %q", expected, repeated)
		}
	})
	t.Run("repeat 1 times", func(t *testing.T) {
		repeated := Repeat("a", 1)
		expected := "a"
		if repeated != expected {
			t.Errorf("expected %q, but got %q", expected, repeated)
		}
	})
	t.Run("repeat zero times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		if repeated != expected {
			t.Errorf("expected %q, but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	//Output: aaaaa
}
