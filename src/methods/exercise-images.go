package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
	// return color.RGBA{uint8(x*x), uint8(y*y), 255, 255}
	// return color.RGBA{uint8(x*y), uint8(x*y), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
