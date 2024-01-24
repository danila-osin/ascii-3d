package geometry

import (
	"github.com/danila-osin/ascii-3d/pkg/calculator"
	"math"
)

var (
	ZeroVec3Int   = Vec3[int]{X: 0, Y: 0}
	ZeroVec3Float = Vec3[float64]{X: 0, Y: 0}
)

type Vec3[T calculator.Number] struct {
	X, Y, Z T
}

func (v Vec3[T]) Int() Vec3[int] {
	return Vec3[int]{
		X: int(v.X),
		Y: int(v.Y),
		Z: int(v.Z),
	}
}

func (v Vec3[T]) Float64() Vec3[float64] {
	return Vec3[float64]{
		X: float64(v.X),
		Y: float64(v.Y),
		Z: float64(v.Z),
	}
}

func (v Vec3[T]) Len() float64 {
	return math.Sqrt(float64(v.Dot(v)))
}

func (v Vec3[T]) Norm() Vec3[float64] {
	l := v.Len()
	return Vec3[float64]{
		X: float64(v.X) / l,
		Y: float64(v.Y) / l,
		Z: float64(v.Z) / l,
	}
}

func (v Vec3[T]) Add(o Vec3[T]) Vec3[T] {
	return Vec3[T]{
		X: v.X + o.X,
		Y: v.Y + o.Y,
		Z: v.Z + o.Z,
	}
}

func (v Vec3[T]) Sub(o Vec3[T]) Vec3[T] {
	return Vec3[T]{
		X: v.X - o.X,
		Y: v.Y - o.Y,
		Z: v.Z - o.Z,
	}
}

func (v Vec3[T]) Mul(o Vec3[T]) Vec3[T] {
	return Vec3[T]{
		X: v.X * o.X,
		Y: v.Y * o.Y,
		Z: v.Z * o.Z,
	}
}

func (v Vec3[T]) Div(o Vec3[T]) Vec3[T] {
	return Vec3[T]{
		X: v.X / o.X,
		Y: v.Y / o.Y,
		Z: v.Z / o.Z,
	}
}

func (v Vec3[T]) AddN(n T) Vec3[T] {
	return Vec3[T]{
		X: v.X + n,
		Y: v.Y + n,
		Z: v.Z + n,
	}
}

func (v Vec3[T]) SubN(n T) Vec3[T] {
	return Vec3[T]{
		X: v.X - n,
		Y: v.Y - n,
		Z: v.Z - n,
	}
}

func (v Vec3[T]) MulN(n T) Vec3[T] {
	return Vec3[T]{
		X: v.X * n,
		Y: v.Y * n,
		Z: v.Z * n,
	}
}

func (v Vec3[T]) DivN(n T) Vec3[T] {
	return Vec3[T]{
		X: v.X / n,
		Y: v.Y / n,
		Z: v.Z / n,
	}
}

func (v Vec3[T]) Dot(o Vec3[T]) T {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}
