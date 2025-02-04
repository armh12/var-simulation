package math_utils

import (
	"fmt"
	"math"
	"testing"
)

// Define test functions for integration
func squareFunction(x float64) float64 {
	return x * x // f(x) = x^2
}

func sineFunction(x float64) float64 {
	return math.Sin(x) // f(x) = sin(x)
}

func expFunction(x float64) float64 {
	return math.Exp(x) // f(x) = e^x
}

func TestBasicFunctionsIntegration(t *testing.T) { // do not use Trapezoidal, use Simpsons in code
	tests := []struct {
		name      string
		function  func(float64) float64
		x         []float64
		method    IntegrateMethod
		expected  float64
		tolerance float64
	}{
		{
			name:      "Trapezoidal Method - Square Function",
			function:  squareFunction,
			x:         []float64{0, 1, 2, 3, 4},
			method:    Trapezoidal,
			expected:  21.3333333, // Expected value for x^2 from 0 to 4
			tolerance: 0.001,
		},
		{
			name:      "Simpsons Method - Square Function",
			function:  squareFunction,
			x:         []float64{0, 1, 2, 3, 4},
			method:    Simpsons,
			expected:  21.3333333,
			tolerance: 0.001,
		},
		{
			name:      "Trapezoidal Method - Sine Function",
			function:  sineFunction,
			x:         []float64{0, math.Pi / 8, math.Pi / 4, 3 * math.Pi / 8, math.Pi / 2},
			method:    Trapezoidal,
			expected:  1.0, // Integral of sin(x) from 0 to π/2 is 1
			tolerance: 0.02,
		},
		{
			name:      "Simpsons Method - Sine Function",
			function:  sineFunction,
			x:         []float64{0, math.Pi / 8, math.Pi / 4, 3 * math.Pi / 8, math.Pi / 2},
			method:    Simpsons,
			expected:  1.0,
			tolerance: 0.002,
		},
		{
			name:      "Trapezoidal Method - Exp Function",
			function:  expFunction,
			x:         []float64{0, 0.5, 1, 1.5, 2, 2.5, 3},
			method:    Trapezoidal,
			expected:  19.08554,
			tolerance: 0.4,
		},
		{
			name:      "Simpsons Method - Exp Function",
			function:  expFunction,
			x:         []float64{0, 0.5, 1, 1.5, 2, 2.5, 3},
			method:    Simpsons,
			expected:  19.08554,
			tolerance: 0.01,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Integrate(tt.x, tt.function, tt.method)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if math.Abs(result-tt.expected) > tt.tolerance {
				t.Errorf("Integration failed for %s. Expected: %f, Got: %f", tt.name, tt.expected, result)
			}
			fmt.Printf("Integration success for %s. Expected: %f, Got: %f", tt.name, tt.expected, result)
		})
	}
}

func TestSortSlice(t *testing.T) {
	t.Run("Test Sort Slice", func(t *testing.T) {
		// Test case for sorting the slice
		x := []float64{3, 1, 4, 1, 5, 9}
		expected := []float64{1, 1, 3, 4, 5, 9}

		sortedX := sortSlice(x)
		for i, val := range sortedX {
			if val != expected[i] {
				t.Errorf("Sorting failed. Expected: %v, Got: %v", expected, sortedX)
				break
			}
		}
	})
}

func TestGetFunctionValues(t *testing.T) {
	t.Run("Test Get Function Values", func(t *testing.T) {
		// Test case for getFunctionValues
		x := []float64{1, 2, 3}
		expected := []float64{1, 4, 9} // f(x) = x^2

		result := getFunctionValues(x, squareFunction)
		for i, val := range result {
			if val != expected[i] {
				t.Errorf("Function values failed. Expected: %v, Got: %v", expected, result)
				break
			}
		}
	})
}

// complex functions integration
func squareAndExpFunction(x float64) float64 {
	return math.Exp(x) * (x * x) // f(x) = e^x *x^2
}

func sineAndExpFunction(x float64) float64 {
	return math.Sin(x) * math.Exp(x) // f(x) = sin(x) * math.Exp(x)
}

func expSquaredFunction(x float64) float64 {
	return math.Exp(x) * math.Exp(x) // f(x) = e^x * e^x
}

func twoExpAndExpFunction(x float64) float64 {
	return math.Exp2(x) * math.Exp(x) // f(x) = e^x * 2^x
}

func TestComplexFunctionsIntegration(t *testing.T) {
	tests := []struct {
		name      string
		function  func(float64) float64
		x         []float64
		method    IntegrateMethod
		expected  float64
		tolerance float64
	}{
		{
			name:      "Trapezoidal Method - Square and Exp Function",
			function:  squareAndExpFunction,
			x:         []float64{0, 0.5, 1, 1.5, 2, 2.5, 3},
			method:    Trapezoidal,
			expected:  98.4,
			tolerance: 0.2,
		},
		{
			name:      "Simpsons Method - Square and Exp Function",
			function:  squareAndExpFunction,
			x:         []float64{0, 0.5, 1, 1.5, 2, 2.5, 3},
			method:    Simpsons,
			expected:  98.42768,
			tolerance: 0.25,
		},
		{
			name:      "Trapezoidal Method - Sine and Exp Function",
			function:  sineAndExpFunction,
			x:         []float64{0, 0.25, 0.5, 0.75, 1, 1.25, 1.5, 1.75, 2, 2.25, 2.5, 2.75, 3},
			method:    Trapezoidal,
			expected:  11.8595,
			tolerance: 0.1,
		},
		{
			name:      "Simpsons Method - Sine and Exp Function",
			function:  sineAndExpFunction,
			x:         []float64{0, 0.25, 0.5, 0.75, 1, 1.25, 1.5, 1.75, 2, 2.25, 2.5, 2.75, 3},
			method:    Trapezoidal,
			expected:  11.8595,
			tolerance: 0.1,
		},
		{
			name:      "Trapezoidal Method - Exp Squared Function",
			function:  expSquaredFunction,
			x:         []float64{0, 0.25, 0.5, 0.75, 1, 1.25, 1.5, 1.75, 2, 2.25, 2.5, 2.75, 3},
			method:    Trapezoidal,
			expected:  201.2144,
			tolerance: 0.4,
		},
		{
			name:      "Simpsons Method - Exp Squared Function",
			function:  expSquaredFunction,
			x:         []float64{0, 0.25, 0.5, 0.75, 1, 1.25, 1.5, 1.75, 2, 2.25, 2.5, 2.75, 3},
			method:    Simpsons,
			expected:  201.2144,
			tolerance: 0.4,
		},
		{
			name:      "Trapezoidal Method - Two Exp and Exp Function",
			function:  twoExpAndExpFunction,
			x:         []float64{0, 0.5, 1, 1.5, 2, 2.5, 3},
			method:    Trapezoidal,
			expected:  19.08554,
			tolerance: 0.4,
		},
		{
			name:      "Simpsons Method - Two Exp and Exp Function",
			function:  twoExpAndExpFunction,
			x:         []float64{0, 0.5, 1, 1.5, 2, 2.5, 3},
			method:    Simpsons,
			expected:  19.08554,
			tolerance: 0.4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Integrate(tt.x, tt.function, tt.method)
			if err != nil {
				t.Errorf("Error while integrating: %v", err)
				return
			}
			if math.Abs(result-tt.expected) > tt.tolerance {
				t.Errorf("Integration failed. Expected: %v, Got: %v", tt.expected, result)
			}
			fmt.Printf("Integration successed. Expected: %v, Got: %v", tt.expected, result)
		})
	}
}
