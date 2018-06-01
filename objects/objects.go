package objects

import (
	"image/color"

	"github.com/CyrusRoshan/asteroids/physics"
	"github.com/CyrusRoshan/asteroids/sprites"

	"github.com/faiface/pixel"
)

type objectSpriteConst struct {
	pixel.Sprite
}

var (
	PLANET_ONE = objectSpriteConst{*sprites.LoadSprite("img/planet_one.png")}
)

type Object struct {
	Sprite  *objectSpriteConst
	Physics *physics.Physics
	Color   color.Color
}

func NewObject(radius float64, color color.Color, sprite objectSpriteConst) Object {
	p := physics.NewPhysics(0, 0, 0, 0, true, radius)

	return Object{
		Sprite:  &sprite,
		Physics: &p,
		Color:   color,
	}
}
