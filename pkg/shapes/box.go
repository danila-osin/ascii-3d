package shapes

import (
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"math"
)

type Box struct {
	// Size
	size geometry.Vec3[float64]

	// Position
	position geometry.Vec3[float64]

	// Out Normal
	normal geometry.Vec3[float64]
}

func NewBox(size, position geometry.Vec3[float64]) Box {
	return Box{size: size, position: position, normal: geometry.ZeroVec3Float}
}

func (r Box) Intersect(ro, rd geometry.Vec3[float64]) (geometry.Vec2[float64], geometry.Vec3[float64]) {
	m := geometry.Vec3[float64]{X: 1, Y: 1, Z: 1}.Div(rd)
	n := m.Mul(ro.Sub(r.position))
	k := m.Abs().Mul(r.size)

	t1 := n.MulN(-1).Sub(k)
	t2 := n.MulN(-1).Add(k)

	tN := math.Max(math.Max(t1.X, t1.Y), t1.Z)
	tF := math.Min(math.Min(t2.X, t2.Y), t2.Z)

	if tN > tF || tF < 0.0 {
		return geometry.Vec2[float64]{X: -1, Y: -1}, geometry.Vec3[float64]{X: 0, Y: 0, Z: 0}
	}

	yzx := geometry.Vec3[float64]{X: t1.Y, Y: t1.Z, Z: t1.X}
	zxy := geometry.Vec3[float64]{X: t2.Z, Y: t2.X, Z: t2.Y}

	outNormal := rd.Sign().MulN(-1).Mul(t1.Step(yzx)).Mul(t1.Step(zxy))

	return geometry.Vec2[float64]{X: tN, Y: tF}, outNormal
}
