package game

import (
	"github.com/CyrusRoshan/asteroids/objects"
	"github.com/CyrusRoshan/asteroids/screen"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	NEARBY_SECTORS = 1
)

type Game struct {
	window  *pixelgl.Window
	canvas  *pixelgl.Canvas
	sectors objects.SectorHolder
}

func NewGame(win *pixelgl.Window) *Game {
	g := Game{
		window: win,
		canvas: screen.NewCanvas(screen.ScreenBounds()),
	}

	return &g
}

func (g *Game) Run() {
	for !g.window.Closed() {
		screen.LimitFPS(30, func() {
			screen.ScaleWindowToCanvas(g.window, g.canvas)

			g.GetChanges()
			g.HandleCalculations()
			g.Draw()

			g.window.Update()
		})
	}
}

func (g *Game) GetChanges() {

}

func (g *Game) HandleCalculations() {

}

func (g *Game) Draw() {
	g.window.Clear(colornames.Red)
	g.canvas.Clear(colornames.Black)

	// actual draw stuff here

	// g.canvas.Draw(g.window, pixel.IM.Moved(g.canvas.Bounds().Center()))
	g.canvas.Draw(g.window, pixel.IM)
}
