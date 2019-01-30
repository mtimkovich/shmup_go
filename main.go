package main

import (
	"log"

	"github.com/Bredgren/geo"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// A 16x9 resolution to mimic Drawables smartphone (or an arcade cabinet).
const (
	SCREEN_WIDTH  = 240
	SCREEN_HEIGHT = 426
)

var (
	ScreenRect  geo.Rect = geo.RectWH(SCREEN_WIDTH, SCREEN_HEIGHT)
	SHIP_PNG    *ebiten.Image
	MISSILE_PNG *ebiten.Image
)

// Loop through Drawable objects to write to the screen.
type Drawable interface {
	Update() error
	Draw(*ebiten.Image)
}

var Drawables map[Drawable]bool

type Game struct {
	Title  string
	Player *Player
}

func NewGame() *Game {
	g := &Game{
		Title:  "shmup",
		Player: NewPlayer(),
	}

	Drawables = map[Drawable]bool{
		g.Player: true,
	}

	return g
}

func (g *Game) Update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	for d, _ := range Drawables {
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
}

func main() {
	g := NewGame()
	if err := ebiten.Run(g.Update, SCREEN_WIDTH, SCREEN_HEIGHT, 2, g.Title); err != nil {
		log.Fatal(err)
	}
}
