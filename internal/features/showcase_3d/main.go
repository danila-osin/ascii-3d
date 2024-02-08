package showcase_3d

import (
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/pkg/controls"
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"github.com/danila-osin/ascii-3d/pkg/mathx"
	rot "github.com/danila-osin/ascii-3d/pkg/rotation"
	"github.com/danila-osin/ascii-3d/pkg/screen"
	"github.com/danila-osin/ascii-3d/pkg/shapes"
)

var (
	controlsPosition = geometry.Vec2[int]{X: 0, Y: 0}
	maxDist          = 999.9
)

type Showcase3D struct {
	config   config.Config
	screen   *screen.Screen
	state    *state
	controls *controls.Controls
}

func New(config config.Config, screen *screen.Screen) Showcase3D {
	st := &state{
		colors:           defaultColors,
		colorsSize:       len(defaultColors),
		cameraRot:        rot.ZeroRotation,
		cameraPos:        geometry.Vec3{X: -5, Y: 0, Z: 0},
		initialCameraDir: geometry.Vec3{X: 1, Y: 0, Z: 0},
	}

	return Showcase3D{
		config:   config,
		screen:   screen,
		state:    st,
		controls: setupControls(config, st),
	}
}

func (s Showcase3D) Run() {
	go s.controls.Listen()
	s.startRenderLoop()
}

func (s Showcase3D) startRenderLoop() {
	light := geometry.Vec3{X: -20, Y: 70, Z: -100}.Norm()

	sphere1 := shapes.NewSphere(1, geometry.Vec3{X: 0, Y: 1, Z: 0})
	sphere2 := shapes.NewSphere(2, geometry.Vec3{X: 0, Y: 5, Z: 0})
	box := shapes.NewBox(geometry.Vec3{X: 1, Y: 2, Z: 0.5}, geometry.Vec3{X: 0, Y: -3, Z: 0})
	//plane := shapes.NewPlane(geometry.Vec3{X: 0, Y: 0, Z: -1}, 1)

	brFn := screen.BRenderFn(func() {
		s.screen.IterateAndSet(func(rawCursor geometry.Vec2[int], value string) string {
			ssVec := sizeToVec2[float64](s.screen.Size)

			cursor := rawCursor.Float64().Div(ssVec).MulN(2).SubN(1)
			cursor.X *= s.screen.Aspect * s.config.FontAspect

			camRayDir := rot.RotateVec3Intrinsic(s.state.cameraRot, s.state.initialCameraDir.Add(cursor.Vec3(0))).Norm()

			minIt := geometry.Vec2[float64]{X: maxDist, Y: maxDist}
			var n geometry.Vec3

			// Sphere 1
			sphere1It := sphere1.Intersect(s.state.cameraPos, camRayDir)
			if sphere1It.X > 0 && sphere1It.X < minIt.X {
				minIt = sphere1It
				it := s.state.cameraPos.Sub(sphere1.Pos).Add(camRayDir.MulN(sphere1It.X))
				n = it.Norm()
			}

			sphere2It := sphere2.Intersect(s.state.cameraPos, camRayDir)
			if sphere2It.X > 0 && sphere2It.X < minIt.X {
				minIt = sphere2It
				it := s.state.cameraPos.Sub(sphere2.Pos).Add(camRayDir.MulN(sphere2It.X))
				n = it.Norm()
			}

			// Box
			boxIt, outNormal := box.Intersect(s.state.cameraPos, camRayDir)
			if boxIt.X > 0 && boxIt.X < minIt.X {
				minIt = boxIt
				n = outNormal.Norm()
			}

			if minIt.X < maxDist {

				d := n.Dot(light)
				color := mathx.Clamp(int(d*20), 2, s.state.colorsSize-1)
				return s.state.colors[color]
			}

			//Plane
			//planeIt := plane.Intersect(s.state.cameraPos, camRayDir)
			//if planeIt > 0 {
			//	return s.state.colors[1]
			//}

			return s.screen.EmptyPixel
		})

		s.screen.AddText(controlsPosition, s.controls.Descriptions.Text())
	})

	s.screen.StartRenderLoop(true, &brFn, nil)
}

func sizeToVec2[T mathx.Number](s screen.Size) geometry.Vec2[T] {
	return geometry.Vec2[T]{X: T(s.W), Y: T(s.H)}
}
