package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	SCREEN_WIDTH  = 240
	SCREEN_HEIGHT = 426
	SCREEN_BUFFER = 20
)

var (
	ship *ebiten.Image
)

func ReadPngFromFile(filepath string) *ebiten.Image {
	imgFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()

	png, err := png.Decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}
	img, _ := ebiten.NewImageFromImage(png, ebiten.FilterDefault)
	return img
}

func (g *Game) drawShip(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	imgX, _ := g.Ship.Size()
	cenX := float64(x - imgX/2)
	cenY := float64(SCREEN_HEIGHT - SCREEN_BUFFER)
	op.GeoM.Translate(cenX, cenY)
	screen.DrawImage(g.Ship, op)
}

type Game struct {
	Title string
	Ship  *ebiten.Image
}

func NewGame() *Game {
	return &Game{
		Title: "shmup",
		Ship:  ReadPngFromFile("img/ship.png"),
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	// TODO: TouchPosition
	x, y := ebiten.CursorPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("(%v, %v)", 2, y))

	g.drawShip(screen, x, y)
	return nil
}

func main() {
	g := NewGame()
	if err := ebiten.Run(g.Update, SCREEN_WIDTH, SCREEN_HEIGHT, 2, g.Title); err != nil {
		log.Fatal(err)
	}
}
