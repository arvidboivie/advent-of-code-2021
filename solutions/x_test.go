package solutions

import "testing"

func TestX(t *testing.T) {
	X := 1
	Y := 1

	if X != Y {
		t.Fatalf("Expected %d, got %d", X, Y)
	}
}
