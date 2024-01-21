package calculator

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
