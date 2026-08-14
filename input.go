package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Input is a per-tick snapshot of player intent. Game objects consume intent,
// not hardware keys. This seam later allows gamepads, remapping, replays, and
// tests without rewriting Player.
type Input struct {
	Move Vector
}

func ReadInput() Input {
	var move Vector
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		move.X--
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		move.X++
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		move.Y--
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		move.Y++
	}

	return Input{
		Move: move,
	}
}
