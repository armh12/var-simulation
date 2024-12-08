package utils

import (
	"math"
	"math/rand"
)

func GaussianDistribution(mean float64, standardDeviation float64) float64 {
	return mean + (rand.NormFloat64() * standardDeviation)
}

func UniformDistribution(lowerLimit float64, upperLimit float64) float64 {
	return lowerLimit + (rand.Float64() * (upperLimit - lowerLimit))
}

func CauchyDistribution(location float64, scale float64) float64 {
	return location + (rand.NormFloat64() / scale)
}

func LogNormalDistribution(mean float64, standardDeviation float64) float64 {
	return mean + (rand.NormFloat64() * standardDeviation)
}

func ExponentialDistribution(mean float64) float64 {
	return -mean * math.Log(rand.Float64())
}

func WeibullDistribution(scale float64, shape float64) float64 {
	return scale * math.Pow(rand.Float64(), 1/shape)
}

func ParetoDistribution(scale float64, shape float64) float64 {
	return scale / math.Pow(rand.Float64(), 1/shape)
}

func GammaDistribution(shape float64, scale float64) float64 {
	return scale * math.Pow(rand.Float64(), 1/shape)
}
