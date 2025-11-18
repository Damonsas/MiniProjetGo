package gameInterface

import (
	_ "image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game interface {
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
	Update() error
	Draw(screen *Image)
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
