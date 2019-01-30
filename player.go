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
var PlayerArea geo.Rect = geo.RectXYWH(
	X_BUFFER,
	Y_BUFFER,
	SCREEN_WIDTH-X_BUFFER*2,
	SCREEN_HEIGHT-Y_BUFFER*2,
)

type Player struct {
	img  *ebiten.Image
	rect geo.Rect
}

func NewPlayer() *Player {
	p := &Player{
		img: SHIP_PNG,
	}
	size := geo.VecXYi(p.img.Size())
	p.rect = geo.RectWH(size.XY())

	return p
}

func (p *Player) shoot() {
	Drawables[NewMissile(p.rect.TopMid())] = true
}

func (p *Player) Update() error {
	// TODO: TouchPosition
	x, _ := ebiten.CursorPosition()
	p.rect.SetMid(float64(x), SCREEN_HEIGHT)
	p.rect.Clamp(PlayerArea)

	p.shoot()

	return nil
}

func (p *Player) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.rect.TopLeft())
	dst.DrawImage(p.img, op)
}
