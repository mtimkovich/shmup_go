package main

import (
	"fmt"
	"image/color"
	"log"

	"golang.org/x/image/font"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
)

// A 16x9 resolution to mimic Drawables smartphone (or an arcade cabinet).
const (
	TITLE         = "shmup"
	SCREEN_WIDTH  = 240
	SCREEN_HEIGHT = 426
	FONT_SIZE     = 8
)

var (
	SHIP_PNG    *ebiten.Image
	MISSILE_PNG *ebiten.Image
	Score       int
	arcadeFont  font.Face
)

// Loop through Drawable objects to write to the screen.
type Drawable interface {
	Update() error
	Draw(*ebiten.Image)
}

var Drawables map[Drawable]bool

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw the score
	scoreStr := fmt.Sprintf("%02d", Score)
	scoreBuffer := 5
	text.Draw(screen, scoreStr, arcadeFont, scoreBuffer, FONT_SIZE+scoreBuffer, color.White)

	for d := range Drawables {
		err := d.Update()

		if err == nil {
			d.Draw(screen)
		}
	}

	return nil
}

func loadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func init() {
	// Load resources into memory
	SHIP_PNG = loadImage("img/ship.png")
	MISSILE_PNG = loadImage("img/missile.png")

	tt, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    FONT_SIZE,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

func main() {
	Drawables = map[Drawable]bool{
		NewPlayer(): true,
	}

	if err := ebiten.Run(update, SCREEN_WIDTH, SCREEN_HEIGHT, 2, TITLE); err != nil {
		log.Fatal(err)
	}
}
