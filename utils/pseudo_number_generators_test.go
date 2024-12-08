package utils

import (
	"testing"
	"var-simulation/utils"
)

func TestPseudoNumberGenerator(t *testing.T) {
	minVal, maxVal := -1.0, 1.0
	for i := 0; i < 100; i++ {
		num, err := utils.PseudoNumberGenerator(minVal, maxVal)
		if err != nil {
			t.Fatalf("Error generating random number: %v", err)
		}
		if num < minVal || num > maxVal {
			t.Errorf("Generated number %f is out of bounds [%f, %f]", num, minVal, maxVal)
		}
	}
}
