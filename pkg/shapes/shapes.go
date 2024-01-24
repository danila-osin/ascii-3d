package shapes

import (
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"github.com/danila-osin/ascii-3d/pkg/mathx"
	"math"
)

func Line(x, y, k, c float64) bool {
	return mathx.Eq(math.Round(k*x+c), y, math.Abs(k)/2)
}

func Circle(x, y, a, b, r float64) bool {
	return mathx.Lte((x-a)*(x-a)+(y-b)*(y-b), r*r, 0)
}

func Sphere(ro, rd geometry.Vec3[float64], r float64) geometry.Vec2[float64] {
	b := ro.Dot(rd)
	c := ro.Dot(ro) - r*r
	h := b*b - c

	if h < 0 {
		return geometry.Vec2[float64]{X: -1, Y: -1}
	}

	h = math.Sqrt(h)
	return geometry.Vec2[float64]{X: -b - h, Y: -b + h}
}
