package utils

import "math"

type Executor struct {
	a       float64
	rounder *Rounder
}

func MakeExecutor(a float64, rounder *Rounder) *Executor {
	executor := new(Executor)
	executor.a = a
	executor.rounder = rounder
	return executor
}

func (executor *Executor) FNext(uPrev []float64, curIndex int) float64 {
	return executor.fHalf(uPrev[curIndex+1], uPrev[curIndex])
}

func (executor *Executor) FPrev(uPrev []float64, curIndex int) float64 {
	return executor.fHalf(uPrev[curIndex], uPrev[curIndex-1])
}

func (executor *Executor) fHalf(leftVal float64, rightVal float64) float64 {
	x := 0.5*executor.a*(leftVal+rightVal) - 0.5*math.Abs(executor.a)*(leftVal-rightVal)
	return executor.rounder.ToFixed(x)
}
