package gameInterface

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	backgroundColor color.RGBA
}

type Image struct {
	size image.Point
}

type Game interface {
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
	Update() error
	Draw(screen *Image)
}

func NewGame() *game {
	return &game{
		backgroundColor: color.RGBA{120, 0, 235, 255},
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 480, 270
}

func (g *game) Update() error {
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(g.backgroundColor)
}
