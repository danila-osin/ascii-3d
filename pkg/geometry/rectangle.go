package geometry

import "github.com/danila-osin/ascii-3d/pkg/mathx"

type Rectangle[T mathx.Number] struct {
	a Vec2[T]
	b Vec2[T]
}
