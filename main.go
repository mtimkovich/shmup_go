package main

import (
	"log"

	"github.com/Bredgren/geo"
	"github.com/hajimehoshi/ebiten"
)

// A 16x9 resolution to mimic a smartphone (or an arcade cabinet).
const (
	SCREEN_WIDTH  = 240
	SCREEN_HEIGHT = 426
)

var ScreenRect geo.Rect = geo.RectWH(SCREEN_WIDTH, SCREEN_HEIGHT)

type Drawable interface {
	Update()
	Draw(*ebiten.Image)
	Index(int) // Set the location of the object in the drawable list
}

var Drawables []Drawable

func AddToDrawables(d Drawable) {
	Drawables = append(Drawables, d)
	d.Index(len(Drawables) - 1)
}

type Game struct {
	Title  string
	Player *Player
}

func NewGame() *Game {
	g := &Game{
		Title:  "shmup",
		Player: NewPlayer(),
	}

	AddToDrawables(g.Player)

	return g
}

func (g *Game) Update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	for _, d := range Drawables {
		d.Update()
		d.Draw(screen)
	}

	return nil
}

func main() {
	g := NewGame()
	if err := ebiten.Run(g.Update, SCREEN_WIDTH, SCREEN_HEIGHT, 1, g.Title); err != nil {
		log.Fatal(err)
	}
}
