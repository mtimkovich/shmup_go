package main

import (
	"fmt"
	"image/gif"
	"log"

	"golang.org/x/image/colornames"
	"golang.org/x/image/font"

	"github.com/Bredgren/geo"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
)

const (
	TITLE = "shmup"
	// A 16x9 resolution to mimic a smartphone (or an arcade cabinet).
	SCREEN_WIDTH  = 243
	SCREEN_HEIGHT = 432
	FONT_SIZE     = 8
	TEXT_ROW1     = FONT_SIZE + 5
	TEXT_ROW2     = FONT_SIZE*2 + 5
)

var (
	SHIP_PNG    *ebiten.Image
	MISSILE_PNG *ebiten.Image
	STAR_FIELD  *gif.GIF
	ScreenRect  geo.Rect = geo.RectWH(SCREEN_WIDTH, SCREEN_HEIGHT)
	Score       int
	arcadeFont  font.Face
	background  *Background
)

// Loop through Drawable objects to write to the screen.
type Drawable interface {
	Update() error
	Draw(*ebiten.Image)
}

var Drawables map[Drawable]bool

func drawScore(dst *ebiten.Image) {
	scoreStr := fmt.Sprintf("%02d", Score)
	text.Draw(dst, "1UP", arcadeFont, FONT_SIZE*3, TEXT_ROW1, colornames.Red)
	// TODO: This'll break if score is longer than 7 characters.
	text.Draw(dst, scoreStr, arcadeFont, FONT_SIZE*(7-len(scoreStr)), TEXT_ROW2, colornames.White)
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	background.UpdateAndDraw(screen)

	for d := range Drawables {
		err := d.Update()

		if err == nil {
			d.Draw(screen)
		}
	}

	drawScore(screen)

	return nil
}

func loadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func loadGIF(path string) *gif.GIF {
	file, err := ebitenutil.OpenFile(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := gif.DecodeAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func init() {
	// Load resources into memory
	SHIP_PNG = loadImage("img/ship.png")
	MISSILE_PNG = loadImage("img/missile.png")
	STAR_FIELD = loadGIF("img/starfield.gif")

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
	background = &Background{img: STAR_FIELD}
	Drawables = map[Drawable]bool{
		NewPlayer(): true,
	}

	if err := ebiten.Run(update, SCREEN_WIDTH, SCREEN_HEIGHT, 2, TITLE); err != nil {
		log.Fatal(err)
	}
}
