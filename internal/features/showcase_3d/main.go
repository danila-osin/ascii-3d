package showcase_3d

import (
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"github.com/danila-osin/ascii-3d/pkg/mathx"
	"github.com/danila-osin/ascii-3d/pkg/screen"
	"github.com/danila-osin/ascii-3d/pkg/shapes"
	"math"
)

var defaultColors = []string{" ", ".", ":", "!", "/", "r", "(", "l", "1", "Z", "4", "H", "9", "W", "8", "$", "@"}

type state struct {
	colors     []string
	colorsSize int
	cameraPos  geometry.Vec3[float64]
}

type Showcase3D struct {
	config config.Config
	screen *screen.Screen
	state  *state
}

func New(config config.Config, screen *screen.Screen) Showcase3D {
	st := &state{
		colors:     defaultColors,
		colorsSize: len(defaultColors),
		cameraPos:  geometry.Vec3[float64]{X: 0, Y: 0, Z: 2},
	}

	return Showcase3D{
		config: config,
		screen: screen,
		state:  st,
	}
}

func (s Showcase3D) Run() {
	s.startRenderLoop()
}

func (s Showcase3D) startRenderLoop() {
	light := geometry.Vec3[float64]{X: 0, Y: 0, Z: 0}.Norm()
	fr := 0
	brFn := screen.BRenderFn(func() {
		s.screen.IterateAndSet(func(rawCursor geometry.Vec2[int], value string) string {
			ssVec := sizeToVec2[float64](s.screen.Size)

			cursor := rawCursor.Float64().Div(ssVec).MulN(2).SubN(1)
			cursor.X *= s.screen.Aspect * s.config.FontAspect

			ro := geometry.Vec3[float64]{X: 0, Y: 0, Z: -2}
			rd := cursor.Vec3(1).Norm()

			i := shapes.Sphere(ro, rd, 1)
			if i.X > 0 {
				it := ro.Add(rd.MulN(i.X))
				n := it.Norm()
				d := n.Dot(light)
				color := mathx.Clamp(int(d*20), 0, s.state.colorsSize-1)
				return s.state.colors[color]
			}

			return s.screen.EmptyPixel
		})
	})

	aRenderFn := screen.ARenderFn(func() {
		fr += 1

		light = geometry.Vec3[float64]{X: math.Sin(float64(fr) * 0.01), Y: math.Sin(float64(fr) * 0.01), Z: -1}.Norm()
	})

	s.screen.StartRenderLoop(true, &brFn, &aRenderFn)
}

func sizeToVec2[T mathx.Number](s screen.Size) geometry.Vec2[T] {
	return geometry.Vec2[T]{X: T(s.W), Y: T(s.H)}
}
