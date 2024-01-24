package geometry

import (
	"github.com/danila-osin/ascii-3d/pkg/calculator"
	"github.com/danila-osin/ascii-3d/pkg/screen"
)

func Vec2FromSize[T calculator.Number](s screen.Size) Vec2[T] {
	return Vec2[T]{X: T(s.W), Y: T(s.H)}
}
