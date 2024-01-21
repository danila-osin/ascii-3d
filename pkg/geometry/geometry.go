package geometry

var (
	ZeroVec2Int   = Vec2[int]{X: 0, Y: 0}
	ZeroVec2Float = Vec2[float64]{X: 0, Y: 0}
)

type Number interface {
	int | float64
}

type Vec2[T Number] struct {
	X, Y T
}

func (v Vec2[T]) Add(other Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}
