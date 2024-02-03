package rot

import "github.com/danila-osin/ascii-3d/pkg/angle"

var (
	ZeroRotation = Rotation{
		Pitch: angle.Radian(0),
		Yaw:   angle.Radian(0),
		Roll:  angle.Radian(0),
	}
)

type Rotation struct {
	// Y Axis
	Pitch angle.Angle

	// Z Axis
	Yaw angle.Angle

	// X Axis
	Roll angle.Angle
}
