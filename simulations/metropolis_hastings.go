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

func (m *MetropolisHastingSimulation) optimizationCheck(acceptedPoints int, nPoints int) {
	acceptanceRatio := float64(acceptedPoints) / float64(nPoints)
	if acceptanceRatio < 0.33 || acceptanceRatio > 0.5 {
		fmt.Printf("Warning: Acceptance ratio (%.2f) is outside the recommended range (0.33 - 0.5)\n", acceptanceRatio)
	}
}

func (m *MetropolisHastingSimulation) Simulate(nPoints int) ([]float64, error) {
	if nPoints <= 0 {
		return nil, fmt.Errorf("number of points must be greater than zero")
	}

	samples := make([]float64, 0, nPoints)
	var acceptedPoints int

	x, err := utils.PseudoNumberGenerator(m.LowerLimit, m.UpperLimit)
	if err != nil {
		return nil, fmt.Errorf("error generating initial random point")
	}

	for i := 0; i < nPoints; i++ {
		trialPoint, err := m.selectTrialPoint(x)
		if err != nil {
			return nil, fmt.Errorf("error selecting trial point: %v", err)
		}

		if *trialPoint < m.LowerLimit || *trialPoint > m.UpperLimit {
			continue
		}
		acceptanceProbability := m.calculateAcceptanceProbability(x, *trialPoint)

		if acceptanceProbability >= 1 {
			x = *trialPoint
			acceptedPoints++
		} else {
			randomNumber, err := utils.PseudoNumberGenerator(0, 1)
			if err != nil {
				return nil, fmt.Errorf("error generating random number for acceptance check: %v", err)
			}
			if randomNumber <= acceptanceProbability {
				x = *trialPoint
				acceptedPoints++
			}
		}
		samples = append(samples, x)
	}

	m.optimizationCheck(acceptedPoints, nPoints)

	return samples, nil
}
