package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image MyImage
type Image struct{}

// Bounds MyImage bounds
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

// ColorModel MyImage colormodel
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At MyImage at
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
