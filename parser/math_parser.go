package parser

import (
	"math"
	"var-simulation/types"
)

func parseNumber[T types.Number](num T) float64 {
	return float64(num)
}

// logarithm function
func logOperation[T types.Number](logRoot T) func(x T) float64 {
	conversedLogRoot := parseNumber(logRoot)
	return func(x T) float64 {
		conversedX := parseNumber(x)
		return math.Log(conversedX) / math.Log(conversedLogRoot)
	}
}

func Log[T types.Number](logRoot T, x T) float64 {
	return logOperation(logRoot)(x)
}

// CTan 1 / tan
func CTan[T types.Number](x T) float64 {
	return 1 / math.Tan(parseNumber(x))
}

// ExpFunc custom exponential function
func ExpFunc[T types.Number](number T, x int) float64 {
	conversedNumber := float64(number)
	return math.Pow(conversedNumber, float64(x))
}
