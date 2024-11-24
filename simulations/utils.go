package simulations

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

type Function func(x float64) float64

func pseudoNumberGenerator[T Number](lowerLimit, upperLimit T) float64 {
	return 0.0
}
