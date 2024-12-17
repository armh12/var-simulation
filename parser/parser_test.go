package parser

import (
	"github.com/expr-lang/expr"
	"math"
	"testing"
)

func concatMaps(m1, m2 map[string]interface{}) map[string]interface{} {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

func TestUserInputBasicFunctions(t *testing.T) {
	tests := []struct {
		stringFunction   string
		expectedFunction func(x float64) float64
	}{
		{
			stringFunction: "sin(x)",
			expectedFunction: func(x float64) float64 {
				return math.Sin(x)
			},
		},
		{
			stringFunction: "cos(x)",
			expectedFunction: func(x float64) float64 {
				return math.Cos(x)
			},
		},
		{
			stringFunction: "exp(x)",
			expectedFunction: func(x float64) float64 {
				return math.Exp(x)
			},
		},
		{
			stringFunction: "sin(exp(x))",
			expectedFunction: func(x float64) float64 {
				return math.Sin(math.Exp(x))
			},
		},
	}

	for _, test := range tests {
		program, err := ParseUserInput(test.stringFunction)
		if err != nil {
			t.Errorf("Error parsing user input: %v", err)
		}
		parseEnv := GetEnvironmentForParse()
		runEnv := map[string]interface{}{
			"x": 1.0,
		}
		runEnv = concatMaps(runEnv, parseEnv)
		result, err := expr.Run(program, runEnv)
		if err != nil {
			t.Errorf("Error running user input: %v", err)
		}
		if result != test.expectedFunction(1.0) {
			t.Errorf("Expected %v, got %v", test.expectedFunction(1.0), result)
		}
	}
}

func TestUserComplexFunctions(t *testing.T) {
	//tests := []struct {
	//	stringFunction string
	//	expectedFunction func (float64) float64
	//}{
	//	//
	//}
}
