package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"var-simulation/types"
)

func PseudoNumberGenerator[T types.Number](lowerLimit, upperLimit T) (float64, error) {
	lowerValue := reflect.ValueOf(lowerLimit)
	upperValue := reflect.ValueOf(upperLimit)

	if lowerValue.Kind() != reflect.Int && lowerValue.Kind() != reflect.Float64 ||
		upperValue.Kind() != reflect.Int && upperValue.Kind() != reflect.Float64 {
		return 0, fmt.Errorf("unsupported type for lowerLimit or upperLimit")
	}

	lowerFloat := float64(lowerLimit)
	upperFloat := float64(upperLimit)
	x := lowerFloat + rand.Float64()*(upperFloat-lowerFloat)
	return x, nil
}
