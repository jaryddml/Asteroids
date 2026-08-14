package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// main is the composition root: it creates the game, configures the window,
// and hands control to Ebitengine. Game rules belong in other files.
func main() {
	game := &Game{}
	game.player = NewPlayer(game)

	// The window can be larger than the logical canvas returned by Layout.
	// Ebitengine scales the game while all rules keep using stable coordinates.
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Top-Down Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
