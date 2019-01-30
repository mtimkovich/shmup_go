package main

import (
	"github.com/Bredgren/geo"
	"github.com/hajimehoshi/ebiten"
)

const (
	X_BUFFER = 5
	Y_BUFFER = 10
)

// Padding around the screen
var PlayerArea geo.Rect = geo.RectCorners(
	X_BUFFER,
	Y_BUFFER,
	SCREEN_WIDTH-X_BUFFER,
	SCREEN_HEIGHT-Y_BUFFER,
)

type Player struct {
	img *ebiten.Image
	box geo.Rect
}

func NewPlayer() *Player {
	p := &Player{
		img: SHIP_PNG,
	}
	size := geo.VecXYi(p.img.Size())
	p.box = geo.RectWH(size.XY())

	return p
}

// Create a new missile and add it to the drawable map.
func (p *Player) Shoot() {
	Drawables[NewMissile(p.box.TopMid())] = true
}

func (p *Player) Update() error {
	// TODO: TouchPosition
	x, _ := ebiten.CursorPosition()
	p.box.SetMid(float64(x), SCREEN_HEIGHT)
	p.box.Clamp(PlayerArea)

	p.Shoot()

	return nil
}

func (p *Player) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.box.TopLeft())
	dst.DrawImage(p.img, op)
}
