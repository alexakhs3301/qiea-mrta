package quantum

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func PlotFitnessOverTime(fitness []float64, filename string) error {
	p := plot.New()
	p.Title.Text = "QIEA Best Fitness Over Time"
	p.X.Label.Text = "Generation"
	p.Y.Label.Text = "Fitness"

	points := make(plotter.XYs, len(fitness))
	for i := range fitness {
		points[i].X = float64(i)
		points[i].Y = fitness[i]
	}

	line, err := plotter.NewLine(points)
	if err != nil {
		return err
	}
	line.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255} // Blue

	p.Add(line)
	p.Legend.Add("Best Fitness", line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, filename); err != nil {
		return fmt.Errorf("failed to save plot: %w", err)
	}

	return nil
}
