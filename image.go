package goca

import (
	"image"
	"image/color"
	"math"
)

type StateToColor interface {
	color.Model
	Color(state State) color.Color
}

type MonochroicConverter struct {
	MinState State
	MaxState State
}

func (c MonochroicConverter) Color(s State) color.Color {
	if s <= c.MinState {
		return color.White
	}

	if s >= c.MaxState {
		return color.Black
	}

	ratio := float64(s-c.MinState) / float64(c.MaxState-c.MinState)
	return color.Gray16{
		uint16(float64(math.MaxUint16) * ratio),
	}
}

func (c MonochroicConverter) Convert(cl color.Color) color.Color {
	gray, _ := color.Gray16Model.Convert(cl).(color.Gray16)
	if gray == color.White || gray == color.Black {
		return gray
	}

	for s := c.MinState; s <= c.MaxState; s++ {
		ratio := float64(s-c.MinState) / float64(c.MaxState-c.MinState)
		y := uint16(float64(math.MaxUint16) * ratio)
		if y <= gray.Y {
			return color.Gray16{y}
		}
	}
	return color.Black
}

type SquareLattice1DImage struct {
	Patterns     []Pattern
	StateToColor StateToColor
}

func (img SquareLattice1DImage) ColorModel() color.Model {
	return img.StateToColor
}

func (img SquareLattice1DImage) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: int(img.Patterns[0].Topology.NumCells()) - 1,
			Y: len(img.Patterns) - 1,
		},
	}
}

func (img SquareLattice1DImage) At(x, y int) color.Color {
	p := img.Patterns[y]
	return img.StateToColor.Color(p.Get(int64(x)))
}
