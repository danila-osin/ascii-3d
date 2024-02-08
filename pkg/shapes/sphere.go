package shapes

import (
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"math"
)

type Sphere struct {
	Pos geometry.Vec3
	Rad float64
}

func NewSphere(radius float64, position geometry.Vec3) Sphere {
	return Sphere{Pos: position, Rad: radius}
}

func (s Sphere) Intersect(ro, rd geometry.Vec3) geometry.Vec2[float64] {
	l := ro.Sub(s.Pos)
	b := l.Dot(rd)
	c := l.Dot(l) - s.Rad*s.Rad
	h := b*b - c

	if h < 0 {
		return geometry.Vec2[float64]{X: -1, Y: -1}
	}

	h = math.Sqrt(h)
	return geometry.Vec2[float64]{X: -b - h, Y: -b + h}
}
