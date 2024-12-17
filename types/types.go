package types

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

type Function func(float64) float64
