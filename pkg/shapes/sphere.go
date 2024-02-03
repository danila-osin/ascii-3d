package shapes

import (
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"math"
)

type Sphere struct {
	center geometry.Vec3[float64]
	radius float64
}

func NewSphere(center geometry.Vec3[float64], radius float64) Sphere {
	return Sphere{center: center, radius: radius}
}

func (s Sphere) Intersect(ro, rd geometry.Vec3[float64]) geometry.Vec2[float64] {
	l := ro.Sub(s.center)
	b := l.Dot(rd)
	c := l.Dot(l) - s.radius*s.radius
	h := b*b - c

	if h < 0 {
		return geometry.Vec2[float64]{X: -1, Y: -1}
	}

	h = math.Sqrt(h)
	return geometry.Vec2[float64]{X: -b - h, Y: -b + h}
}
