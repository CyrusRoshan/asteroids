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

	screenArea pixel.Rect
}

func NewGame(win *pixelgl.Window) *Game {
	g := Game{
		window: win,
		canvas: screen.NewCanvas(screen.ScreenBounds()),
	}

	g.screenArea = g.canvas.Bounds()
	g.setupDemo()

	return &g
}

func (g *Game) setupDemo() {
	g.sectors = objects.NewSectorHolder(1, 1, 500, 500)
	demoSector := g.sectors.GetSector(0, 0)
	g.screenArea = g.screenArea.Moved(g.sectors.GetCenterOfSector(0, 0))

	planet := objects.NewObject(100, colornames.Red, objects.PLANET_ONE)
	demoSector.Objects = append(demoSector.Objects, planet)
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

	drawSector := g.sectors.GetSector(0, 0)
	moveVec := g.screenArea.Center().Sub(g.sectors.GetCenterOfSector(0, 0))
	sectorMatrix := pixel.IM.Moved(moveVec)

	for _, object := range drawSector.Objects {
		matrix := object.Physics.GetMatrix().Chained(sectorMatrix)
		object.Sprite.DrawColorMask(g.canvas, matrix, object.Color)

	}

	g.canvas.Draw(g.window, pixel.IM.Moved(g.canvas.Bounds().Center()))
	// g.canvas.Draw(g.window, pixel.IM)
}
