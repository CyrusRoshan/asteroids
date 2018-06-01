package physics

import (
	"math"

	"github.com/faiface/pixel"
)

func unitVector(v pixel.Vec) pixel.Vec {
	return pixel.V(1, 1).Project(v)
}

type Touch struct {
	Touching       bool
	TouchDirection pixel.Vec
	TouchStrength  float64
}

func isTouching(p1 Physics, p2 Physics) Touch {
	if p1.isRect && p2.isRect {
		c1 := p1.rect.Center()
		c2 := p2.rect.Center()
		dx := c1.X - c2.X
		dy := c1.Y - c2.Y

		return isTouchingRectRect(p1, p2, dx, dy)
	}

	if !p1.isRect && !p2.isRect {
		dx := float64(p1.x - p2.x)
		dy := float64(p1.y - p2.y)

		return isTouchingCircleCircle(p1, p2, dx, dy)
	}

	return Touch{}

	// TODO: handle rect-circle intersections
	// if (p1.isRect && !p2.isRect) {

	// }
}

func isTouchingCircleCircle(p1 Physics, p2 Physics, dx float64, dy float64) Touch {
	combinedRadii := float64(p1.radius + p2.radius)
	boundsOverlap := pythagoreanTheorem(dx, dy) - combinedRadii

	if boundsOverlap < 0 {
		return Touch{Touching: false}
	}

	return Touch{
		Touching:       true,
		TouchDirection: unitVector(pixel.V(dx, dy)),
		TouchStrength:  boundsOverlap,
	}
}

// TODO: handle rotated rects
func isTouchingRectRect(p1 Physics, p2 Physics, dx float64, dy float64) Touch {
	intersection := p1.rect.Intersect(p2.rect)

	if (intersection == pixel.Rect{}) {
		return Touch{}
	}

	intersectionX := math.Copysign(intersection.Max.X-intersection.Min.X, dx)
	intersectionY := math.Copysign(intersection.Max.Y-intersection.Min.X, dy)

	return Touch{
		Touching:       true,
		TouchDirection: unitVector(pixel.V(intersectionX, intersectionY)),
		TouchStrength:  pythagoreanTheorem(intersectionX, intersectionY),
	}
}

// TODO: complete function
func isTouchingCircleRect(p1 Physics, p2 Physics, dx float64, dy float64) Touch {
	return Touch{}
}
