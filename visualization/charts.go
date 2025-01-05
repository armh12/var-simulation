package visualization

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"image"
)

func PlotHistogram(values []float64, bins int, widthInch, heightInch float64) (image.Image, error) {
	p := plot.New()
	p.Title.Text = "Histogram"
	p.X.Label.Text = "Value"
	p.Y.Label.Text = "Frequency"

	hist, err := plotter.NewHist(plotter.Values(values), bins)
	if err != nil {
		return nil, err
	}
	// hist.Color = color.RGBA{R: 0, G: 128, B: 255, A: 255}
	// hist.FillColor = color.RGBA{R: 0, G: 128, B: 255, A: 100}

	p.Add(hist)

	img, err := GetImageBuffer(widthInch, heightInch, p)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func PlotLine(xValues, yValues []float64, widthInch, heightInch float64) (image.Image, error) {
	p := plot.New()
	p.Title.Text = "Line Plot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	pts := make(plotter.XYs, len(xValues))
	for i := range xValues {
		pts[i].X = xValues[i]
		pts[i].Y = yValues[i]
	}
	line, err := plotter.NewLine(pts)
	if err != nil {
		return nil, err
	}
	p.Add(line)

	img, err := GetImageBuffer(widthInch, heightInch, p)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func PlotRunningMean(samples []float64, widthInch, heightInch float64) (image.Image, error) {
	n := len(samples)

	running := make([]float64, n)
	sum := 0.0
	for i, v := range samples {
		sum += v
		running[i] = sum / float64(i+1)
	}

	p := plot.New()
	p.Title.Text = "Running Mean"
	p.X.Label.Text = "Iteration"
	p.Y.Label.Text = "Cumulative Mean"

	pts := make(plotter.XYs, n)
	for i, v := range running {
		pts[i].X = float64(i)
		pts[i].Y = v
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		return nil, err
	}
	p.Add(line)

	img, err := GetImageBuffer(widthInch, heightInch, p)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func PlotTrace(samples []float64, widthInch, heightInch float64) (image.Image, error) {
	p := plot.New()
	p.Title.Text = "Trace Plot"
	p.X.Label.Text = "Iteration"
	p.Y.Label.Text = "Value"

	// Преобразуем samples в формат, удобный для Gonum (XY-данные).
	pts := make(plotter.XYs, len(samples))
	for i, v := range samples {
		pts[i].X = float64(i)
		pts[i].Y = v
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		return nil, err
	}
	// Можно настроить цвет / стиль
	// line.Color = color.RGBA{R: 255, A: 255}

	p.Add(line)

	// Рендерим в память (PNG):
	img, err := GetImageBuffer(widthInch, heightInch, p)
	if err != nil {
		return nil, err
	}
	return img, nil
}
