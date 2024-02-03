package mathx

type Number interface {
	int | float64
}

func Eq[T Number](a, b T, p float64) bool {
	return Gte(a, b, p) && Lte(a, b, p)
}

func Gt[T Number](a, b T, p float64) bool {
	return float64(a) > (float64(b) - p)
}

func Gte[T Number](a, b T, p float64) bool {
	return float64(a) >= (float64(b) - p)
}

func Lt[T Number](a, b T, p float64) bool {
	return float64(a) <= (float64(b) + p)
}

func Lte[T Number](a, b T, p float64) bool {
	return float64(a) <= (float64(b) + p)
}

func Clamp[T Number](val, min, max T) T {
	if val < min {
		return min
	}

	if val > max {
		return max
	}

	return val
}

func Sign(a float64) float64 {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	}
	return 0
}

func Step(edge, x float64) float64 {
	if x > edge {
		return 1
	}
	return 0
}
