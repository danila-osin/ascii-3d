package angle

import "math"

type Angle struct {
	deg float64
	rad float64
}

func Radian(a float64) Angle {
	return Angle{
		deg: RadToDeg(a),
		rad: a,
	}
}

func Degree(a float64) Angle {
	return Angle{
		deg: a,
		rad: DegToRad(a),
	}
}

func (a Angle) Rad() float64 {
	return a.rad
}

func (a Angle) Deg() float64 {
	return a.deg
}

func (a Angle) Add(o Angle) Angle {
	return Angle{
		deg: a.deg + o.deg,
		rad: a.rad + o.rad,
	}
}

func (a Angle) Sub(o Angle) Angle {
	return Angle{
		deg: a.deg - o.deg,
		rad: a.rad - o.rad,
	}
}

func DegToRad(d float64) float64 {
	return d * math.Pi / 180.0
}

func RadToDeg(r float64) float64 {
	return r * 180.0 / math.Pi
}
