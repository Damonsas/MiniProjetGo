package main

import (
	"log"
	"projet3-graphicInterface/gameInterface"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Test :D")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	err := ebiten.RunGame(gameInterface.NewGame())
	if err != nil {
		log.Fatal(err)
	}
}
