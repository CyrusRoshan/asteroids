package physics

// physics should really only ever interact with other physics.
// Therefore, all struct values should be private, and methods should be used to read/write

import (
	"github.com/faiface/pixel"
)

type Physics struct {
	immobile bool // simplifies planet gravitational fields, for now
	mass     int

	x int
	y int

	vx float64
	vy float64

	radius int
	rect   pixel.Rect
	isRect bool // true if rect is being used, false if radius is being used

	rotation float64
}

type PhysicsBounds interface {
}
