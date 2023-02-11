package utils

import "math"

type Rounder struct {
	precision int
}

func MakeRounder(precision int) *Rounder {
	rounder := new(Rounder)
	rounder.precision = precision
	return rounder
}

func (rounder *Rounder) ToFixed(num float64) float64 {
	if rounder.precision == -1 {
		return num
	}
	out := math.Pow(10, float64(rounder.precision))
	return math.Round(num*out) / out
}
