package shapes

import (
	"github.com/danila-osin/ascii-3d/pkg/calculator"
	"math"
)

func Line(x, y, k, c float64) bool {
	return calculator.Eq(math.Round(k*x+c), y, math.Abs(k)/2)
}

func Circle(x, y, a, b, r float64) bool {
	return calculator.Lte((x-a)*(x-a) + (y-b)*(y-b), r*r, 0)
}
