package utils

import (
	"gonum.org/v1/plot/plotter"
	"math"
	"sync"
)

type Executor struct {
	a       float64
	rounder *Rounder
	x       []float64
}

func MakeExecutor(a float64, rounder *Rounder, x []float64) *Executor {
	executor := new(Executor)
	executor.a = a
	executor.rounder = rounder
	executor.x = x
	return executor
}

func (executor *Executor) fNext(uPrev []float64, curIndex int) float64 {
	return executor.fHalf(uPrev[curIndex+1], uPrev[curIndex])
}

func (executor *Executor) fPrev(uPrev []float64, curIndex int) float64 {
	return executor.fHalf(uPrev[curIndex], uPrev[curIndex-1])
}

func (executor *Executor) fHalf(leftVal float64, rightVal float64) float64 {
	x := 0.5*executor.a*(leftVal+rightVal) - 0.5*math.Abs(executor.a)*(leftVal-rightVal)
	return executor.rounder.ToFixed(x)
}

func (executor *Executor) Calculate(u *[]float64, uPrev []float64, from int, to int, tCur float64, isLastDetour bool, pts *plotter.XYs, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := from; i <= to; i++ {
		hCur := executor.x[i]
		(*u)[i] = executor.rounder.ToFixed(uPrev[i] - tCur/hCur*(executor.fNext(uPrev, i)-executor.fPrev(uPrev, i)))
		if isLastDetour {
			(*pts)[i-1].X = executor.x[i]
			(*pts)[i-1].Y = (*u)[i]
		}
	}
}

func (executor *Executor) Initialize(uPrev *[]float64, from int, to int, pts *plotter.XYs, h float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := from; i <= to; i++ {
		executor.x[i] = executor.rounder.ToFixed(h * float64(i))
		(*uPrev)[i] = getUStartValue(executor.x[i])
		(*pts)[i-1].X = executor.x[i]
		(*pts)[i-1].Y = (*uPrev)[i]
	}
}

func getUStartValue(xValue float64) float64 {
	if xValue < 0.25 || xValue > 0.75 {
		return 0
	}
	return 1
}
