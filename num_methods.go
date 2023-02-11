package main

import (
	"gonum.org/v1/plot/plotter"
	"num_methods/utils"
)

func getUStartValue(xValue float64) float64 {
	if xValue < 0.25 || xValue > 0.75 {
		return 0
	}
	return 1
}

func main() {
	// a = 1 const
	//h = 1 / (n - 1)
	//t : at/h < 1  -> t/h < 1, t < h
	a := 1.0

	x0 := 0.0
	xN := 1.0
	n := 100

	t0 := 0.0
	tM := 0.1
	m := 100

	rounder := utils.MakeRounder(-1)
	executor := utils.MakeExecutor(a, rounder)

	h := rounder.ToFixed((xN - x0) / float64(n-1))
	hTime := rounder.ToFixed((tM - t0) / float64(m-1))
	x := make([]float64, n+2)
	uPrev := make([]float64, n+2)

	pts := make(plotter.XYs, n)
	for i := 1; i < n+1; i++ {
		x[i] = rounder.ToFixed(h * float64(i))
		uPrev[i] = getUStartValue(x[i])
		pts[i-1].X = x[i]
		pts[i-1].Y = uPrev[i]
	}
	utils.DrawBase(pts)

	uPrev[0] = uPrev[n]
	uPrev[n+1] = uPrev[1]

	pts = make(plotter.XYs, n)
	var tCur float64
	for j := 1; j < m; j++ {
		tCur = rounder.ToFixed(hTime * float64(j))
		u := make([]float64, n+2)
		for i := 1; i < n+1; i++ {
			hCur := x[i]
			u[i] = rounder.ToFixed(uPrev[i] - tCur/hCur*(executor.FNext(uPrev, i)-executor.FPrev(uPrev, i)))
			if j+1 == m {
				pts[i-1].X = x[i]
				pts[i-1].Y = u[i]
			}
		}
		uPrev = u
	}

	utils.DrawShift(pts)
}
