package math_utils

import (
	"fmt"
	"gonum.org/v1/gonum/integrate"
	"sort"
	"var-simulation/types"
)

type IntegrateMethod int

// Define constants for the methods
const (
	Trapezoidal IntegrateMethod = iota
	Simpsons
)

func Integrate(x []float64, f types.Function, method IntegrateMethod) (float64, error) {
	x = sortSlice(x)
	switch method {
	case Trapezoidal:
		return integrateTrapezoidal(x, f), nil
	case Simpsons:
		return integrateSimpsons(x, f), nil
	default:
		return 0, fmt.Errorf("error while integrating")
	}
}

func integrateSimpsons(x []float64, f types.Function) float64 {
	funcValues := getFunctionValues(x, f)
	integrateResult := integrate.Simpsons(x, funcValues)
	return integrateResult
}

func integrateTrapezoidal(x []float64, f types.Function) float64 {
	funcValues := getFunctionValues(x, f)
	integrateResult := integrate.Trapezoidal(x, funcValues)
	return integrateResult
}

func getFunctionValues(x []float64, f types.Function) []float64 {
	xSliceCapacity := cap(x)
	ySlice := make([]float64, xSliceCapacity)

	for i := 0; i < xSliceCapacity; i++ {
		ySlice[i] = f(x[i])
	}
	return ySlice
}

func sortSlice(x []float64) []float64 {
	sort.Float64s(x)
	return x
}
