package shapes

import "github.com/danila-osin/ascii-3d/pkg/geometry"

type Plane struct {
	// Orientation
	o geometry.Vec3

	// ???
	w float64
}

func NewPlane(o geometry.Vec3, w float64) Plane {
	return Plane{o: o, w: w}
}

func (p Plane) Intersect(ro, rd geometry.Vec3) float64 {
	return -(ro.Dot(p.o) + p.w) / rd.Dot(p.o)
}
