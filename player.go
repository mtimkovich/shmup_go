package main

import (
	_ "image/png"
	"log"

	"github.com/Bredgren/geo"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	X_BUFFER = 5
	Y_BUFFER = 10
)

var PlayerArea geo.Rect = geo.RectXYWH(
	float64(X_BUFFER),
	float64(Y_BUFFER),
	float64(SCREEN_WIDTH-X_BUFFER*2),
	float64(SCREEN_HEIGHT-Y_BUFFER*2),
)

type Player struct {
	x    float64
	y    float64
	img  *ebiten.Image
	rect geo.Rect
}

func NewPlayer() *Player {
	var err error
	p := &Player{}
	p.img, _, err = ebitenutil.NewImageFromFile("img/ship.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	size := geo.VecXYi(p.img.Size())
	p.rect = geo.RectWH(size.XY())

	return p
}

func (p *Player) Update() {
	// TODO: TouchPosition
	x, _ := ebiten.CursorPosition()
	p.rect.SetMid(float64(x), SCREEN_HEIGHT)
	p.rect.Clamp(PlayerArea)
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.rect.TopLeft())
	screen.DrawImage(p.img, op)
}
