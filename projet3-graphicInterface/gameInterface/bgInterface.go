package gameInterface

import (
	"image"
	"image/color"
)

type game struct {
	backgroundColor color.RGBA
}

type Image struct {
	size image.Point
}

func NewGame() *game {
	return &game{
		backgroundColor: color.RGBA{120, 0, 235, 255},
	}
}
