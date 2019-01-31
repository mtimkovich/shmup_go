package main

import (
	"github.com/Bredgren/geo"
	"github.com/hajimehoshi/ebiten"
)

const (
	BUFFER = 5
)

// Padding around the screen
var PlayerArea geo.Rect = geo.RectCorners(
	BUFFER,
	BUFFER*2,
	SCREEN_WIDTH-BUFFER,
	SCREEN_HEIGHT-BUFFER,
)

type Player struct {
	img  *ebiten.Image
	box  geo.Rect
	tick int
}

func NewPlayer() *Player {
	p := &Player{
		img: SHIP_PNG,
	}
	size := geo.VecXYi(p.img.Size())
	p.box = geo.RectWH(size.XY())

	return p
}

// Create 2 new missiles and add them to the drawable map.
func (p *Player) Shoot() {
	// Shoot every n frames.
	p.tick = (p.tick + 1) % 2

	if p.tick != 0 {
		return
	}

	x, y := p.box.Mid()
	Drawables[NewMissile(x-4, y)] = true
	Drawables[NewMissile(x+4, y)] = true
	Score++
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
