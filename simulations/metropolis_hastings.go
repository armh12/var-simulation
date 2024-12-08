package simulations

import (
	"fmt"
	"var-simulation/types"
	"var-simulation/utils"
)

type MetropolisHastingSimulation struct {
	TargetDistribution   types.Function
	ProposalDistribution types.Function
	Delta                float64
	LowerLimit           float64
	UpperLimit           float64
}

func NewMetropolisHastingSimulation[T types.Number](targetDistribution, proposalDistribution types.Function, delta, lowerLimit, upperLimit T) *MetropolisHastingSimulation {
	return &MetropolisHastingSimulation{
		TargetDistribution:   targetDistribution,
		ProposalDistribution: proposalDistribution,
		Delta:                float64(delta),
		LowerLimit:           float64(lowerLimit),
		UpperLimit:           float64(upperLimit),
	}
}

func (m *MetropolisHastingSimulation) selectTrialPoint(x float64) (*float64, error) {
	randomNumber, err := utils.PseudoNumberGenerator(-m.Delta, m.Delta)
	if err != nil {
		return nil, fmt.Errorf("error while generating random number")
	}
	trialPoint := x + randomNumber
	return &(trialPoint), nil
}

func (m *MetropolisHastingSimulation) calculateAcceptanceProbability(x, trialPoint float64) float64 {
	acceptanceProbability := m.TargetDistribution(x) / m.TargetDistribution(trialPoint)
	return acceptanceProbability
}
