package gui

import (
	"var-simulation/simulations"
)

func SimulateMH(targetDistribution, proposalDistribution func(float64) float64, delta, lowerLimit, upperLimit float64,
	numOfSamples int) ([]float64, error) {
	mh := simulations.NewMetropolisHastingSimulation(
		targetDistribution,
		proposalDistribution,
		delta,
		lowerLimit,
		upperLimit,
	)
	simulateResult, err := mh.Simulate(numOfSamples)
	if err != nil {
		return nil, err
	}
	return simulateResult, nil
}
