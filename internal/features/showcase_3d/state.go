package showcase_3d

import (
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	rot "github.com/danila-osin/ascii-3d/pkg/rotation"
)

var (
	defaultColors = []string{" ", ".", ":", "!", "r", "/", "(", "l", "1", "Z", "4", "H", "9", "W", "8", "$", "@"}
)

type state struct {
	colors           []string
	colorsSize       int
	cameraRot        rot.Rotation
	cameraPos        geometry.Vec3
	initialCameraDir geometry.Vec3
}

func (s state) cameraDir() geometry.Vec3 {
	return rot.RotateVec3Intrinsic(s.cameraRot, s.initialCameraDir).Norm()
}
