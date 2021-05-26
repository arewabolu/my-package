package calc

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("adds two positive numbers", func(t *testing.T) {
		val := Add(2, 4)
		if val != 6 {
			t.Fatalf("expected result to be 6, got %f", val)
		}
	})

	t.Run("adding 0 does nothing", func(t *testing.T) {
		x := 12.0
		val := Add(x, 0)
		if val != x {
			t.Fatalf("expected result to be %f, got %f", x, val)
		}
	})

	t.Run("adding a negative subtracts", func(t *testing.T) {
		val := Add(12, -4)
		if val >= 12 {
			t.Fatalf("expected result to be less than 12, got %f", val)
		}
	})

	t.Run("adding two negatives adds their magnitudes", func(t *testing.T) {
		x, y := -3.0, -5.0
		val := Add(x, y)
		abs := math.Abs(x) + math.Abs(y)
		if val != -abs {
			t.Fatalf("expected result to be %f, got %f", val, abs)
		}
	})
}
