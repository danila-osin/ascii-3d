package matrix

import (
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"github.com/danila-osin/ascii-3d/pkg/mathx"
)

type Matrix[T mathx.Number] struct {
	inner [][]T
	size  Size
}

func New[T mathx.Number](m [][]T) Matrix[T] {
	return Matrix[T]{
		inner: m,
		size: Size{
			M: len(m),
			N: len(m[0]),
		},
	}
}

func FromVec3[T mathx.Number](v geometry.Vec3[T]) Matrix[T] {
	return New([][]T{
		{v.X},
		{v.Y},
		{v.Z},
	})
}

func (m Matrix[T]) Vec3() geometry.Vec3[T] {
	return geometry.Vec3[T]{
		X: m.Get(0, 0),
		Y: m.Get(1, 0),
		Z: m.Get(2, 0),
	}
}

func (m Matrix[T]) Size() Size {
	return m.size
}

func (m Matrix[T]) Get(i, j int) T {
	return m.inner[i][j]
}

func (m Matrix[T]) Set(i, j int, val T) {
	m.inner[i][j] = val
}

func (m Matrix[T]) Add(o Matrix[T]) Matrix[T] {
	r := make([][]T, m.size.M)

	for i := range m.inner {
		l := make([]T, m.size.N)
		for j := range m.inner[i] {
			l[j] = m.Get(i, j) + o.Get(i, j)
		}
		r[i] = l
	}

	return New[T](r)
}

func (m Matrix[T]) Sub(o Matrix[T]) Matrix[T] {
	r := make([][]T, m.size.M)

	for i := range m.inner {
		l := make([]T, m.size.N)
		for j := range m.inner[i] {
			l[j] = m.Get(i, j) - o.Get(i, j)
		}
		r[i] = l
	}

	return New[T](r)
}

func (m Matrix[T]) Mul(o Matrix[T]) Matrix[T] {
	r := make([][]T, m.size.M)

	for i := 0; i < m.size.M; i++ {
		l := make([]T, o.size.N)
		for j := 0; j < o.size.N; j++ {
			ls := T(0)
			for k := 0; k < o.size.M; k++ {
				ls += m.inner[i][k] * o.inner[k][j]
			}
			l[j] = ls
		}
		r[i] = l
	}

	return New(r)
}
