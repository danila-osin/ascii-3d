package screen

// TODO create optimized matrix ???
type matrix [][]string

func newMatrix(size Size, emptyPixel string) matrix {
	m := make(matrix, size.H)

	for y := 0; y < size.H; y++ {
		l := make([]string, size.W)

		for x := 0; x < size.W; x++ {
			l[x] = emptyPixel
		}

		m[y] = l
	}

	return m
}
