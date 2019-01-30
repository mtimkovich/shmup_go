package main

import (
	"fmt"

	"github.com/Bredgren/geo"
	"github.com/hajimehoshi/ebiten"
)

type Missile struct {
	img  *ebiten.Image
	rect geo.Rect
}

func NewMissile(x, y float64) *Missile {
	m := &Missile{
		img: MISSILE_PNG,
	}
	size := geo.VecXYi(m.img.Size())
	m.rect = geo.RectWH(size.XY())
	m.rect.SetMid(x, y)

	return m
}

func (m *Missile) Update() error {
	_, dy := m.rect.Size()
	m.rect.Move(0, -dy-2)

	if !ScreenRect.Contains(m.rect) {
		delete(Drawables, m)
		return fmt.Errorf("missile offscreen")
	}

	return nil
}

func (m *Missile) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.rect.TopLeft())
	dst.DrawImage(m.img, op)
}
