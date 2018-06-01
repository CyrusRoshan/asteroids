package objects

import "image/color"

type Planet struct {
}

func NewPlanet(radius int, color color.Color) Planet {
	return Planet{}
}
