package main

import (
	"Game/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationPerSecond = math.Pi
)

type Player struct {
	position Vector
	sprite   *ebiten.Image
	game     *Game
	rotation float64
}

func NewPlayer(g *Game) *Player {
	sprite := assets.PlayerSprite
	p := &Player{
		game:   g,
		sprite: sprite,
	}
	return p
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{
		DisableMipmaps: true,
	}

	screen.DrawImage(p.sprite, op)

}

func (p *Player) Update() {
}
