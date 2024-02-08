package shapes

import (
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"github.com/danila-osin/ascii-3d/pkg/mathx"
	"math"
)

type Box struct {
	// Size
	Size geometry.Vec3

	// Position
	Pos geometry.Vec3
}

func NewBox(size, position geometry.Vec3) Box {
	return Box{Size: size, Pos: position}
}

func (r Box) Intersect(ro, rd geometry.Vec3) (geometry.Vec2[float64], geometry.Vec3) {
	var oN geometry.Vec3

	m := geometry.Vec3{X: 1, Y: 1, Z: 1}.Div(rd)
	n := m.Mul(ro.Sub(r.Pos))
	k := m.Abs().Mul(r.Size)
	t1 := n.MulN(-1).Sub(k)
	t2 := n.MulN(-1).Add(k)

	tN := math.Max(math.Max(t1.X, t1.Y), t1.Z)
	tF := math.Min(math.Min(t2.X, t2.Y), t2.Z)

	if tN > tF || tF < 0.0 {
		return geometry.Vec2[float64]{X: -1.0, Y: -1.0}, oN
	}

	oN = geometry.Vec3{
		X: -mathx.Sign(rd.X) * mathx.Step(t1.Y, t1.X) * mathx.Step(t1.Z, t1.X),
		Y: -mathx.Sign(rd.Y) * mathx.Step(t1.Z, t1.Y) * mathx.Step(t1.X, t1.Y),
		Z: -mathx.Sign(rd.Z) * mathx.Step(t1.X, t1.Z) * mathx.Step(t1.Y, t1.Z),
	}

	return geometry.Vec2[float64]{X: tN, Y: tF}, oN
}
