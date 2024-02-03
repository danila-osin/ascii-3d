package shapes

import (
	"github.com/danila-osin/ascii-3d/pkg/mathx"
	"math"
)

func Line(x, y, k, c float64) bool {
	return mathx.Eq(math.Round(k*x+c), y, math.Abs(k)/2)
}

func Circle(x, y, a, b, r float64) bool {
	return mathx.Lte((x-a)*(x-a)+(y-b)*(y-b), r*r, 0)
}
