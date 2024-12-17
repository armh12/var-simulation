package math_utils

import (
	"math"
	"math/rand"
)

type DistributionType int

const (
	Gaussian DistributionType = iota
	Uniform
	Cauchy
	LogNormal
	Exponential
	Weibull
	Pareto
)

type Distributions struct{}

func (d Distributions) gaussianDistribution(mean float64, standardDeviation float64) float64 {
	return mean + (rand.NormFloat64() * standardDeviation)
}

func (d Distributions) uniformDistribution(lowerLimit float64, upperLimit float64) float64 {
	return lowerLimit + (rand.Float64() * (upperLimit - lowerLimit))
}

func (d Distributions) cauchyDistribution(location float64, scale float64) float64 {
	return location + (rand.NormFloat64() / scale)
}

func (d Distributions) logNormalDistribution(mean float64, standardDeviation float64) float64 {
	return mean + (rand.NormFloat64() * standardDeviation)
}

func (d Distributions) exponentialDistribution(mean float64) float64 {
	return -mean * math.Log(rand.Float64())
}

func (d Distributions) weibullDistribution(scale float64, shape float64) float64 {
	return scale * math.Pow(rand.Float64(), 1/shape)
}

func (d Distributions) paretoDistribution(scale float64, shape float64) float64 {
	return scale / math.Pow(rand.Float64(), 1/shape)
}

func (d Distributions) gammaDistribution(shape float64, scale float64) float64 {
	return scale * math.Pow(rand.Float64(), 1/shape)
}
