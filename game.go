package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	worldMargin  = 32
)

type Game struct {
	player *Player
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
