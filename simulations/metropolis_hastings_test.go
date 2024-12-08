package simulations

import (
	"math"
	"testing"
)

func TestMetropolisHastingSimulationBasic(t *testing.T) {
	targetDistribution := func(x float64) float64 {
		return math.Exp(-x * x / 2)
	}
	proposalDistribution := func(x float64) float64 {
		return x
	}

	delta := 1.0
	lowerLimit := -10.0
	upperLimit := 10.0
	simulation := NewMetropolisHastingSimulation(targetDistribution, proposalDistribution, delta, lowerLimit, upperLimit)

	// Run the simulation
	nPoints := 1000
	samples, err := simulation.Simulate(nPoints)
	if err != nil {
		t.Fatalf("Error during simulation: %v", err)
	}

	if len(samples) != nPoints {
		t.Errorf("Expected %d samples, got %d", nPoints, len(samples))
	}
}

func TestMetropolisHastingSimulationOutOfBounds(t *testing.T) {
	targetDistribution := func(x float64) float64 {
		return math.Exp(-x * x / 2)
	}
	proposalDistribution := func(x float64) float64 {
		return x
	}

	delta := 15.0 // Large delta to frequently generate out-of-bounds points
	lowerLimit := -10.0
	upperLimit := 10.0
	simulation := NewMetropolisHastingSimulation(targetDistribution, proposalDistribution, delta, lowerLimit, upperLimit)

	nPoints := 100
	samples, err := simulation.Simulate(nPoints)
	if err != nil {
		t.Fatalf("Error during simulation: %v", err)
	}

	for _, sample := range samples {
		if sample < lowerLimit || sample > upperLimit {
			t.Errorf("Sample %f is out of bounds [%f, %f]", sample, lowerLimit, upperLimit)
		}
	}
}

func TestMetropolisHastingSimulationOptimizationCheck(t *testing.T) {
	targetDistribution := func(x float64) float64 {
		return math.Exp(-x * x / 2)
	}
	proposalDistribution := func(x float64) float64 {
		return x
	}

	delta := 0.1 // Small delta to increase acceptance ratio
	lowerLimit := -10.0
	upperLimit := 10.0
	simulation := NewMetropolisHastingSimulation(targetDistribution, proposalDistribution, delta, lowerLimit, upperLimit)

	nPoints := 1000
	_, err := simulation.Simulate(nPoints)
	if err != nil {
		t.Fatalf("Error during simulation: %v", err)
	}

}

func TestMetropolisHastingSimulationInvalidInput(t *testing.T) {
	targetDistribution := func(x float64) float64 {
		return math.Exp(-x * x / 2)
	}
	proposalDistribution := func(x float64) float64 {
		return x
	}

	delta := 1.0
	lowerLimit := -10.0
	upperLimit := 10.0
	simulation := NewMetropolisHastingSimulation(targetDistribution, proposalDistribution, delta, lowerLimit, upperLimit)

	_, err := simulation.Simulate(0)
	if err == nil {
		t.Errorf("Expected an error for zero points, but got none")
	}
}

func TestExtremeDelta(t *testing.T) {
	targetDistribution := func(x float64) float64 {
		return math.Exp(-x * x / 2)
	}
	proposalDistribution := func(x float64) float64 {
		return x
	}

	simulation := NewMetropolisHastingSimulation(targetDistribution, proposalDistribution, 0.00001, -10.0, 10.0)
	samples, err := simulation.Simulate(100)
	if err != nil {
		t.Fatalf("Error during simulation: %v", err)
	}
	if len(samples) != 100 {
		t.Errorf("Expected 100 samples, got %d", len(samples))
	}

	simulation = NewMetropolisHastingSimulation(targetDistribution, proposalDistribution, 100, -10.0, 10.0)
	samples, err = simulation.Simulate(100)
	if err != nil {
		t.Fatalf("Error during simulation: %v", err)
	}
	if len(samples) != 100 {
		t.Errorf("Expected 100 samples, got %d", len(samples))
	}
}
