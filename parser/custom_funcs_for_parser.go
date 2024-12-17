package parser

import (
	"math"
)

// logarithm function
func logOperation(logRoot float64) func(x float64) float64 {
	return func(x float64) float64 {
		return math.Log(x) / math.Log(logRoot)
	}
}

func Log(logRoot, x float64) float64 {
	return logOperation(logRoot)(x)
}

// CTan 1 / tan
func CTan(x float64) float64 {
	return 1 / math.Tan(x)
}

// ExpFunc custom exponential function
func ExpFunc(number float64, x float64) float64 {
	return math.Pow(number, x)
}
