package simulations

import "math/rand"

type MonteCarloSimulation struct {
	upperLimit float64
	lowerLimit float64
	function   Function
}

//func NewMonteCarloSimulation(upperLimit, lowerLimit int, function Function) *MonteCarloSimulation {
//	return &MonteCarloSimulation{
//		upperLimit: float64(upperLimit),
//		lowerLimit: float64(lowerLimit),
//		function:   function,
//	}
//}

func (m MonteCarloSimulation) limitRange() float64 {
	return m.upperLimit - m.lowerLimit
}

func (m MonteCarloSimulation) simulate(nPoints int) float64 {
	function := m.function
	totalSum := 0.0
	for i := 0; i < nPoints; i++ {
		x := m.lowerLimit + rand.Float64()*m.limitRange()
		totalSum += function(x)
	}
	return totalSum
}
