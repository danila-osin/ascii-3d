package geometry

import "math"

var (
	ZeroVec2Int   = Vec2[int]{X: 0, Y: 0}
	ZeroVec2Float = Vec2[float64]{X: 0, Y: 0}
)

type Vec2[T Number] struct {
	X, Y T
}

func (v Vec2[T]) Len() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v Vec2[T]) Norm() Vec2[float64] {
	l := v.Len()
	return Vec2[float64]{
		X: float64(v.X) / l,
		Y: float64(v.Y) / l,
	}
}

func (v Vec2[T]) Add(o Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}

func (v Vec2[T]) Sub(o Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v.X - o.X,
		Y: v.Y - o.Y,
	}
}

func (v Vec2[T]) Mul(o Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v.X * o.X,
		Y: v.Y * o.Y,
	}
}

func (v Vec2[T]) Div(o Vec2[T]) Vec2[T] {
	return Vec2[T]{
		X: v.X / o.X,
		Y: v.Y / o.Y,
	}
}
