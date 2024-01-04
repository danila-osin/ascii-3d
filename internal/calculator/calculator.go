package calculator

import "math"

type Number interface {
	int | float64
}

func Eq[A Number, B Number](a A, b B, p float64) bool {
	return Gte(a, b, p) && Lte(a, b, p)
}

func Gt[A Number, B Number](a A, b B, p float64) bool {
	return float64(a) > (float64(b) - p)
}

func Gte[A Number, B Number](a A, b B, p float64) bool {
	return float64(a) >= (float64(b) - p)
}

func Lt[A Number, B Number](a A, b B, p float64) bool {
	return float64(a) <= (float64(b) + p)
}

func Lte[A Number, B Number](a A, b B, p float64) bool {
	return float64(a) <= (float64(b) + p)
}

func Bt[N Number, A Number, B Number](n N, a A, b B) bool {
	if float64(a) > float64(b) {
		return Gt(n, b, 0) && Lt(n, a, 0)
	}

	if float64(b) > float64(a) {
		return Gt(n, a, 0) && Lt(n, b, 0)
	}

	return false
}

type Functions struct{}

func (f Functions) Line(x, y, k, c float64, p float64) bool {
	return Eq(math.Round(k*x+c), y, math.Abs(k)/2+p)
}

func (f Functions) Circle(x, y, a, b, r float64, p float64) bool {
	return Eq(math.Floor((x-a)*(x-a)+(y-b)*(y-b)), r*r, math.Abs(x-a)+math.Abs(y-b))
}
