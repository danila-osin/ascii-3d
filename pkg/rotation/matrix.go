package rot

import (
	"github.com/danila-osin/ascii-3d/pkg/angle"
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"github.com/danila-osin/ascii-3d/pkg/matrix"
	"math"
)

type RotateMatrixFn func(angle angle.Angle, m matrix.Matrix[float64]) matrix.Matrix[float64]

func RotateMatrixX(angle angle.Angle, m matrix.Matrix[float64]) matrix.Matrix[float64] {
	sin := math.Sin(angle.Rad())
	cos := math.Cos(angle.Rad())

	rm := matrix.New([][]float64{
		{1, 0, 0},
		{0, cos, -sin},
		{0, sin, cos},
	})

	return rm.Mul(m)
}

func RotateMatrixY(angle angle.Angle, m matrix.Matrix[float64]) matrix.Matrix[float64] {
	sin := math.Sin(angle.Rad())
	cos := math.Cos(angle.Rad())

	rm := matrix.New([][]float64{
		{cos, 0, sin},
		{0, 1, 0},
		{-sin, 0, cos},
	})

	return rm.Mul(m)
}

func RotateMatrixZ(angle angle.Angle, m matrix.Matrix[float64]) matrix.Matrix[float64] {
	sin := math.Sin(angle.Rad())
	cos := math.Cos(angle.Rad())

	rm := matrix.New([][]float64{
		{cos, -sin, 0},
		{sin, cos, 0},
		{0, 0, 1},
	})

	return rm.Mul(m)
}

func RotateMatrix(angle angle.Angle, m matrix.Matrix[float64], axis geometry.Axis) matrix.Matrix[float64] {
	switch axis {
	case geometry.XAxis:
		return RotateMatrixX(angle, m)
	case geometry.YAxis:
		return RotateMatrixY(angle, m)
	case geometry.ZAxis:
		return RotateMatrixZ(angle, m)
	}

	return m
}

func RotateMatrixIntrinsic(rot Rotation, m matrix.Matrix[float64]) matrix.Matrix[float64] {
	cosA := math.Cos(rot.Roll.Rad())
	sinA := math.Sin(rot.Roll.Rad())

	cosB := math.Cos(rot.Yaw.Rad())
	sinB := math.Sin(rot.Yaw.Rad())

	cosC := math.Cos(rot.Pitch.Rad())
	sinC := math.Sin(rot.Pitch.Rad())

	rm := matrix.New([][]float64{
		{cosB * cosC, sinA*sinB*cosC - cosA*sinC, cosA*sinB*cosC + sinA*sinC},
		{cosB * sinC, sinA*sinB*sinC + cosA*cosC, cosA*sinB*sinC - sinA*cosC},
		{-sinB, sinA * cosB, cosA * cosB},
	})

	return rm.Mul(m)
}

func RotateVec3(axis geometry.Axis, angle angle.Angle, v geometry.Vec3[float64]) geometry.Vec3[float64] {
	return RotateMatrix(angle, matrix.FromVec3(v), axis).Vec3()
}

func RotateVec3Intrinsic(rot Rotation, v geometry.Vec3[float64]) geometry.Vec3[float64] {
	return RotateMatrixIntrinsic(rot, matrix.FromVec3(v)).Vec3()
}
