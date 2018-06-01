package physics

// physics should really only ever interact with other physics.
// Therefore, all struct values should be private, and methods should be used to read/write

import (
	"github.com/faiface/pixel"
)

type Physics struct {
	immobile bool // simplifies planet gravitational fields, for now
	mass     int

	x float64
	y float64

	vx float64
	vy float64

	radius float64
	rect   pixel.Rect
	isRect bool

	rotation float64
}

func NewPhysics(x float64, y float64, rotation float64, mass int, immobile bool, radius float64) Physics {
	return Physics{
		immobile: immobile,
		mass:     mass,

		x: x,
		y: x,

		radius: radius,
		isRect: false,

		rotation: rotation,
	}
}

func (p *Physics) GetMatrix() pixel.Matrix {
	m := pixel.IM
	m = m.Rotated(pixel.ZV, p.rotation)
	m = m.Moved(pixel.Vec{X: p.x, Y: p.y})

	return m
}
