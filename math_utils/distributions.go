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

func (d Distributions) GaussianDistribution(mean float64, standardDeviation float64) float64 {
	return mean + (rand.NormFloat64() * standardDeviation)
}

func (d Distributions) UniformDistribution(lowerLimit float64, upperLimit float64) float64 {
	return lowerLimit + (rand.Float64() * (upperLimit - lowerLimit))
}

func (d Distributions) CauchyDistribution(location float64, scale float64) float64 {
	return location + (rand.NormFloat64() / scale)
}

func (d Distributions) LogNormalDistribution(mean float64, standardDeviation float64) float64 {
	return mean + (rand.NormFloat64() * standardDeviation)
}

func (d Distributions) ExponentialDistribution(mean float64) float64 {
	return -mean * math.Log(rand.Float64())
}

func (d Distributions) WeibullDistribution(scale float64, shape float64) float64 {
	return scale * math.Pow(rand.Float64(), 1/shape)
}

func (d Distributions) ParetoDistribution(scale float64, shape float64) float64 {
	return scale / math.Pow(rand.Float64(), 1/shape)
}

func (d Distributions) GammaDistribution(shape float64, scale float64) float64 {
	return scale * math.Pow(rand.Float64(), 1/shape)
}
