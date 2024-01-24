package geometry

import (
	"github.com/danila-osin/ascii-3d/pkg/calculator"
	"math"
)

var (
	ZeroVec2Int   = Vec2[int]{X: 0, Y: 0}
	ZeroVec2Float = Vec2[float64]{X: 0, Y: 0}
)

type Vec2[T calculator.Number] struct {
	X, Y T
}

func (v Vec2[T]) Int() Vec2[int] {
	return Vec2[int]{
		X: int(v.X),
		Y: int(v.Y),
	}
}

func (v Vec2[T]) Float64() Vec2[float64] {
	return Vec2[float64]{
		X: float64(v.X),
		Y: float64(v.Y),
	}
}

func (v Vec2[T]) Len() float64 {
	return math.Sqrt(float64(v.Dot(v)))
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

func (v Vec2[T]) AddN(n T) Vec2[T] {
	return Vec2[T]{
		X: v.X + n,
		Y: v.Y + n,
	}
}

func (v Vec2[T]) SubN(n T) Vec2[T] {
	return Vec2[T]{
		X: v.X - n,
		Y: v.Y - n,
	}
}

func (v Vec2[T]) MulN(n T) Vec2[T] {
	return Vec2[T]{
		X: v.X * n,
		Y: v.Y * n,
	}
}

func (v Vec2[T]) DivN(n T) Vec2[T] {
	return Vec2[T]{
		X: v.X / n,
		Y: v.Y / n,
	}
}

func (v Vec2[T]) Dot(o Vec2[T]) T {
	return v.X*o.X + v.Y*o.Y
}
