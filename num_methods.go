package main

import (
	"gonum.org/v1/plot/plotter"
	"num_methods/utils"
	"sync"
)

func main() {
	// a = 1 const
	//h = 1 / (n - 1)
	//t : at/h < 1  -> t/h < 1, t < h
	a := 1.0

	x0 := 0.0
	xN := 1.0
	n := 10000

	t0 := 0.0
	tM := 0.1
	m := 100

	var wg sync.WaitGroup

	rounder := utils.MakeRounder(-1)
	h := rounder.ToFixed((xN - x0) / float64(n-1))
	hTime := rounder.ToFixed((tM - t0) / float64(m-1))
	x := make([]float64, n+2)
	uPrev := make([]float64, n+2)

	threader := utils.MakeThreader(n)
	executor := utils.MakeExecutor(a, rounder, x)

	pts := make(plotter.XYs, n)
	for idx := 1; idx <= threader.GetThreadsCount(); idx++ {
		from, to := threader.GetInterval(idx)
		wg.Add(1)
		go executor.Initialize(&uPrev, from, to, &pts, h, &wg)
	}
	wg.Wait()
	utils.DrawBase(pts)

	uPrev[0] = uPrev[n]
	uPrev[n+1] = uPrev[1]

	pts = make(plotter.XYs, n)
	for j := 1; j < m; j++ {
		tCur := rounder.ToFixed(hTime * float64(j))
		u := make([]float64, n+2)

		for idx := 1; idx <= threader.GetThreadsCount(); idx++ {
			from, to := threader.GetInterval(idx)
			wg.Add(1)
			go executor.Calculate(&u, uPrev, from, to, tCur, j+1 == m, &pts, &wg)
		}
		wg.Wait()
		uPrev = u
	}
	utils.DrawShift(pts)
}
