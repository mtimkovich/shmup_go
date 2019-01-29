package main

import (
	"github.com/Bredgren/geo"
	"github.com/hajimehoshi/ebiten"
)

type Missile struct {
	img   *ebiten.Image
	rect  geo.Rect
	index int
}

func NewMissile(x, y float64) {
}
