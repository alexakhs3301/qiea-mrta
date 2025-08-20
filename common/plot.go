package common

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func PlotFitnessHistory(history []float64, title, filename string) error {
	pts := make(plotter.XYs, len(history))
	for i := range history {
		pts[i].X = float64(i)
		pts[i].Y = history[i]
	}

	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "Generation"
	p.Y.Label.Text = "Fitness"

	line, err := plotter.NewLine(pts)
	if err != nil {
		return err
	}
	line.Color = color.RGBA{B: 255, A: 255}
	p.Add(line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, filename); err != nil {
		return err
	}
	return nil
}
