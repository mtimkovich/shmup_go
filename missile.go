package main

import (
	"fmt"

	"github.com/Bredgren/geo"
	"github.com/hajimehoshi/ebiten"
)

type Missile struct {
	img *ebiten.Image
	box geo.Rect
}

func NewMissile(x, y float64) *Missile {
	m := &Missile{
		img: MISSILE_PNG,
	}
	size := geo.VecXYi(m.img.Size())
	m.box = geo.RectWH(size.XY())
	m.box.SetMid(x, y)

	return m
}

func (m *Missile) Update() error {
	m.box.Move(0, -m.box.H-2)

	if !m.box.CollideRect(ScreenRect) {
		delete(Drawables, m)
		return fmt.Errorf("missile offscreen")
	}

	return nil
}

func (m *Missile) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.box.TopLeft())
	dst.DrawImage(m.img, op)
}
