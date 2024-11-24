package simulations

type MonteCarloSimulation struct {
	upperLimit float64
	lowerLimit float64
	function   Function
}

func NewMonteCarloSimulation[T Number](upperLimit, lowerLimit T, function Function) *MonteCarloSimulation {
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
		x, err := PseudoNumberGenerator(m.lowerLimit, m.upperLimit)
		if err != nil {
			return nil, err
		}
		totalSum += function(x)
	}
	return &totalSum, nil
}

func (m MonteCarloSimulation) Integrate(nPoints int) (*float64, error) {
	totalSum, err := m.simulateRandomValue(nPoints)
	if err != nil {
		return nil, err
	}
	limitRange := m.limitRange()
	integrationValue := (limitRange / float64(nPoints)) * *totalSum
	return &integrationValue, nil
}
