package geometry

import (
	"github.com/danila-osin/ascii-3d/pkg/mathx"
	"math"
)

var (
	ZeroVec3Float = Vec3{X: 0, Y: 0}
)

type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Len() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec3) Norm() Vec3 {
	l := v.Len()
	return Vec3{
		X: v.X / l,
		Y: v.Y / l,
		Z: v.Z / l,
	}
}

func (v Vec3) Abs() Vec3 {
	return Vec3{
		X: math.Abs(v.X),
		Y: math.Abs(v.Y),
		Z: math.Abs(v.Z),
	}
}

func (v Vec3) Add(o Vec3) Vec3 {
	return Vec3{
		X: v.X + o.X,
		Y: v.Y + o.Y,
		Z: v.Z + o.Z,
	}
}

func (v Vec3) Sub(o Vec3) Vec3 {
	return Vec3{
		X: v.X - o.X,
		Y: v.Y - o.Y,
		Z: v.Z - o.Z,
	}
}

func (v Vec3) Mul(o Vec3) Vec3 {
	return Vec3{
		X: v.X * o.X,
		Y: v.Y * o.Y,
		Z: v.Z * o.Z,
	}
}

func (v Vec3) Div(o Vec3) Vec3 {
	res := Vec3{}

	if o.X == 0 {
		res.X = math.MaxFloat64
	} else {
		res.X = v.X / o.X
	}

	if o.Y == 0 {
		res.Y = math.MaxFloat64
	} else {
		res.Y = v.Y / o.Y
	}

	if o.Z == 0 {
		res.Z = math.MaxFloat64
	} else {
		res.Z = v.Z / o.Z
	}

	return res
}

func (v Vec3) AddN(n float64) Vec3 {
	return Vec3{
		X: v.X + n,
		Y: v.Y + n,
		Z: v.Z + n,
	}
}

func (v Vec3) SubN(n float64) Vec3 {
	return Vec3{
		X: v.X - n,
		Y: v.Y - n,
		Z: v.Z - n,
	}
}

func (v Vec3) MulN(n float64) Vec3 {
	return Vec3{
		X: v.X * n,
		Y: v.Y * n,
		Z: v.Z * n,
	}
}

func (v Vec3) DivN(n float64) Vec3 {
	return Vec3{
		X: v.X / n,
		Y: v.Y / n,
		Z: v.Z / n,
	}
}

func (v Vec3) Dot(o Vec3) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

func (v Vec3) Reflect(o Vec3) Vec3 {
	return v.Sub(o.MulN(o.Dot(v) * 2))
}

func (v Vec3) Step(edge Vec3) Vec3 {
	return Vec3{
		X: mathx.Step(edge.X, v.X),
		Y: mathx.Step(edge.Y, v.Y),
		Z: mathx.Step(edge.Z, v.Z),
	}
}

func (v Vec3) Sign() Vec3 {
	return Vec3{
		X: mathx.Sign(v.X),
		Y: mathx.Sign(v.Y),
		Z: mathx.Sign(v.Z),
	}
}

func (v Vec3) MinAxis() Axis {
	minComp := math.Min(math.Min(v.X, v.Y), v.Z)

	if v.X == minComp {
		return XAxis
	}

	if v.Y == minComp {
		return YAxis
	} else {
		return ZAxis
	}
}
