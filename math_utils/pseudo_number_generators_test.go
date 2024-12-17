package math_utils

import (
	"testing"
)

func TestPseudoNumberGenerator(t *testing.T) {
	minVal, maxVal := -1.0, 1.0
	for i := 0; i < 100; i++ {
		num, err := PseudoNumberGenerator(minVal, maxVal)
		if err != nil {
			t.Fatalf("Error generating random number: %v", err)
		}
		if num < minVal || num > maxVal {
			t.Errorf("Generated number %f is out of bounds [%f, %f]", num, minVal, maxVal)
		}
	}
}
