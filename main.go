package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	SCREEN_WIDTH  = 240
	SCREEN_HEIGHT = 426
)

// var ScreenRect geo.Rect = geo.RectWH(SCREEN_WIDTH, SCREEN_HEIGHT)

type Game struct {
	Title  string
	Player *Player
}

func NewGame() *Game {
	g := &Game{
		Title:  "shmup",
		Player: NewPlayer(),
	}

	return g
}

func (g *Game) Update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	g.Player.Update()
	g.Player.Draw(screen)

	return nil
}

func main() {
	g := NewGame()
	if err := ebiten.Run(g.Update, SCREEN_WIDTH, SCREEN_HEIGHT, 1, g.Title); err != nil {
		log.Fatal(err)
	}
}
