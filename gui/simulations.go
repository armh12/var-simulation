package gui

import (
	"var-simulation/simulations"
	"var-simulation/types"
)

func SimulateMH(targetDistribution, proposalDistribution types.Function, delta, lowerLimit, upperLimit float64,
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

func BackgroundSimulateMC(upperLimit, lowerLimit float64, function types.Function, numOfSamples int) (*float64, error) {
	monteCarlo := simulations.NewMonteCarloSimulation(upperLimit, lowerLimit, function)
	simulateResult, err := monteCarlo.Simulate(numOfSamples)
	if err != nil {
		return nil, err
	}
	return simulateResult, nil
}
