package utils

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

func draw(title string, imageName string, pts plotter.XYs) {
	p := plot.New()
	p.Add(plotter.NewGrid())
	p.Title.Text = title
	p.X.Label.Text = "X"
	p.Y.Label.Text = "U(x)"
	lpLine, lpPoints, err := plotter.NewLinePoints(pts)
	if err != nil {
		panic(err)
	}
	lpLine.Color = color.RGBA{B: 255, A: 255}
	lpPoints.Color = color.RGBA{R: 255, A: 255}
	p.Legend.Add("Line", lpLine)
	p.Legend.Add("Point", lpPoints)
	p.Add(lpLine, lpPoints)
	if err := p.Save(10*vg.Inch, 10*vg.Inch, imageName); err != nil {
		panic(err)
	}
}

func DrawBase(pts plotter.XYs) {
	draw("Base", "base.png", pts)
}

func DrawShift(pts plotter.XYs) {
	draw("Shift", "shift.png", pts)
}
