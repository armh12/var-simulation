package visualization

import (
	"bytes"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"image"
	"image/png"
)

func GetImageBuffer(widthInch, heightInch float64, p *plot.Plot) (image.Image, error) {
	wt, err := p.WriterTo(
		vg.Length(widthInch)*vg.Inch,
		vg.Length(heightInch)*vg.Inch,
		"png",
	)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	_, err = wt.WriteTo(buf)
	if err != nil {
		return nil, err
	}

	img, err := png.Decode(buf)
	if err != nil {
		return nil, err
	}

	return img, nil
}
