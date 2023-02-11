package main

import (
	"math"
	"strconv"
)

func toFixed(num float64, precision int) float64 {
	out := math.Pow(10, float64(precision))
	return math.Round(num*out) / out
}

func getUStartValue(xValue float64) float64 {
	if xValue < 0.25 || xValue > 0.75 {
		return 0
	}
	return 1
}

func fNext(uPrev []float64, a float64, curIndex int) float64 {
	return fHalf(uPrev[curIndex+1], uPrev[curIndex], a)
}

func fPrev(uPrev []float64, a float64, curIndex int) float64 {
	return fHalf(uPrev[curIndex], uPrev[curIndex-1], a)
}

func fHalf(leftVal float64, rightVal float64, a float64) float64 {
	x := 0.5*a*(leftVal+rightVal) - 0.5*math.Abs(a)*(leftVal-rightVal)
	return x
}

func main() {
	// a = 1 const
	//h = 1 / (n - 1)
	var n, tm int
	a := 1.0
	n = 10000
	//t : at/h < 1  -> t/h < 1, t < n
	tm = n * 2
	precisionN := len(strconv.Itoa(n))
	precisionT := len(strconv.Itoa(tm))

	h := toFixed(1/float64(n-1), precisionN)
	hTime := toFixed(1/float64(tm-1), precisionT)
	x := make([]float64, n+2)
	uPrev := make([]float64, n+2)
	for i := 1; i < n+1; i++ {
		x[i] = toFixed(h*float64(i), precisionN)
		uPrev[i] = getUStartValue(x[i])
	}
	uPrev[0] = uPrev[n]
	uPrev[n+1] = uPrev[1]

	t := make([]float64, tm)
	t[0] = 0
	for m := 1; m < tm; m++ {
		t[m] = toFixed(hTime*float64(m), precisionT)
		tCur := t[m]
		u := make([]float64, n+2)
		for i := 1; i < n+1; i++ {
			hCur := x[i]
			u[i] = toFixed(uPrev[i]-tCur/hCur*(fNext(uPrev, a, i)-fPrev(uPrev, a, i)), precisionN)
		}
		uPrev = u
	}
}
