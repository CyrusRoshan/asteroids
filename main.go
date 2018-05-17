package main

import (
	"github.com/CyrusRoshan/asteroids/game"
	"github.com/CyrusRoshan/asteroids/screen"
	"github.com/CyrusRoshan/asteroids/utils"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	width, height := screen.ScreenBounds()
	cfg := pixelgl.WindowConfig{
		Title:     "Asteroids",
		Bounds:    pixel.R(0, 0, width/2, height/2),
		VSync:     true,
		Resizable: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	utils.PanicIf(err)

	asteroids := game.NewGame(win)
	asteroids.Run()
}
