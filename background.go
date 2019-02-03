package main

import (
	"image/gif"

	"github.com/Bredgren/geo"
	"github.com/hajimehoshi/ebiten"
)

// Fill screen with tiling background GIF.
type Background struct {
	img      *gif.GIF
	frameImg *ebiten.Image
	frame    int
	tick     int
}

func (b *Background) Update() error {
	b.frameImg, _ = ebiten.NewImageFromImage(b.img.Image[b.frame], ebiten.FilterDefault)

	if b.tick == 0 {
		b.frame = (b.frame + 1) % len(b.img.Image)
	}

	b.tick = (b.tick + 1) % b.img.Delay[b.frame]

	return nil
}

func (b *Background) Draw(dst *ebiten.Image) {
	sizeX, sizeY := geo.VecXYi(b.frameImg.Size()).XY()
	var width, height float64
	opts := &ebiten.DrawImageOptions{}

	for height < SCREEN_HEIGHT {
		for width < SCREEN_WIDTH {
			opts.GeoM.Translate(width, height)
			dst.DrawImage(b.frameImg, opts)
			width += sizeX
		}

		width = 0
		height += sizeY
	}
}

func (b *Background) UpdateAndDraw(dst *ebiten.Image) {
	b.Update()
	b.Draw(dst)
}
