package function_graph

import (
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/pkg/controls"
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"github.com/danila-osin/ascii-3d/pkg/mathx"
	"github.com/danila-osin/ascii-3d/pkg/screen"
	"github.com/danila-osin/ascii-3d/pkg/shapes"
	"math"
	"time"
)

var ControlsPosition = geometry.Vec2[int]{X: 2, Y: 2}

type state struct {
	cameraPos geometry.Vec2[float64]
	scale     float64
}

type FunctionGraph struct {
	config   config.Config
	screen   *screen.Screen
	state    *state
	controls *controls.Controls
}

func New(config config.Config, screen *screen.Screen) FunctionGraph {
	st := &state{cameraPos: geometry.ZeroVec2Float, scale: 1}
	screen.PixelSeparator = ""

	return FunctionGraph{
		config:   config,
		screen:   screen,
		state:    st,
		controls: setupControls(config, st),
	}
}

func (f FunctionGraph) Run() {
	f.setInitialState()
	f.screen.Render(true, false)
	time.Sleep(2 * time.Second)

	go f.controls.Listen()
	f.startRenderLoop()
}

func (f FunctionGraph) setInitialState() {
	f.screen.IterateAndSet(func(rawCursor geometry.Vec2[int], value string) string {
		x := float64(rawCursor.X)/float64(f.screen.Size.W)*2.0 - 1
		y := float64(rawCursor.Y)/float64(f.screen.Size.H)*2.0 - 1

		if mathx.Eq(y, 0, 0) {
			return "."
		}

		if mathx.Eq(x, 0, 0) {
			return "."
		}

		return value
	})

	f.screen.AddText(ControlsPosition, f.controls.Descriptions.Text())
}

func (f FunctionGraph) startRenderLoop() {
	a := 0.0
	direction := 1.0

	brFn := screen.BRenderFn(func() {
		f.screen.IterateAndSet(func(rawCursor geometry.Vec2[int], value string) string {
			ssVec := sizeToVec2[float64](f.screen.Size)

			cursor := rawCursor.Float64().Div(ssVec).MulN(2).SubN(1).MulN(f.state.scale)
			cursor.X *= f.screen.Aspect * f.config.FontAspect

			cursor = cursor.Add(f.state.cameraPos)

			if shapes.Circle(cursor.X, cursor.Y, 0, 0, 0.2) {
				return "x"
			}

			if mathx.Eq(cursor.Y, 0, 0.005) {
				return "."
			}

			if mathx.Eq(cursor.X, 0, 0.005) {
				return "."
			}

			return f.screen.EmptyPixel
		})

		f.screen.AddText(ControlsPosition, f.controls.Descriptions.Text())
	})

	arFn := screen.ARenderFn(func() {
		if math.Abs(a) > 1 {
			direction = -direction
		}

		a += direction * 0.01
	})

	f.screen.StartRenderLoop(true, &brFn, &arFn)
}

func sizeToVec2[T mathx.Number](s screen.Size) geometry.Vec2[T] {
	return geometry.Vec2[T]{X: T(s.W), Y: T(s.H)}
}
