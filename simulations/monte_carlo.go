package simulations

import (
	"var-simulation/types"
	"var-simulation/utils"
)

type MonteCarloSimulation struct {
	upperLimit float64
	lowerLimit float64
	function   types.Function
}

func NewMonteCarloSimulation[T types.Number](upperLimit, lowerLimit T, function types.Function) *MonteCarloSimulation {
	return &MonteCarloSimulation{
		upperLimit: float64(upperLimit),
		lowerLimit: float64(lowerLimit),
		function:   function,
	}
}

func (m MonteCarloSimulation) limitRange() float64 {
	return m.upperLimit - m.lowerLimit
}

func (m MonteCarloSimulation) simulateRandomValue(nPoints int) (*float64, error) {
	function := m.function
	totalSum := 0.0
	for i := 0; i < nPoints; i++ {
		x, err := utils.PseudoNumberGenerator(m.lowerLimit, m.upperLimit)
		if err != nil {
			return nil, err
		}
		totalSum += function(x)
	}
	return &totalSum, nil
}

func (m MonteCarloSimulation) Simulate(nPoints int) (*float64, error) {
	totalSum, err := m.simulateRandomValue(nPoints)
	if err != nil {
		return nil, err
	}
	limitRange := m.limitRange()
	simulationValue := (limitRange / float64(nPoints)) * *totalSum
	return &simulationValue, nil
}
