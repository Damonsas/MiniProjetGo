package menu

import (
	"image"
	"image/color"
)

type MenuBG struct {
	backgroundMenu color.Color
}

type Sprite struct {
	size image.Point
}

func NewMenuBG() *MenuBG {
	return &MenuBG{
		backgroundMenu: color.Black,
	}
}
