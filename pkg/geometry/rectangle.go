package geometry

import "github.com/danila-osin/ascii-3d/pkg/calculator"

type Rectangle[T calculator.Number] struct {
	a Vec2[T]
	b Vec2[T]
}
