package function_graph

import (
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/pkg/calculator"
	"github.com/danila-osin/ascii-3d/pkg/controls"
	"github.com/danila-osin/ascii-3d/pkg/geometry"
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

		if calculator.Eq(y, 0, 0) {
			return "."
		}

		if calculator.Eq(x, 0, 0) {
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
			x := (float64(rawCursor.X)/float64(f.screen.Size.W)*2.0 - 1) * f.state.scale * f.screen.Aspect * f.config.FontAspect
			y := (float64(rawCursor.Y)/float64(f.screen.Size.H)*2.0 - 1) * f.state.scale

			x += f.state.cameraPos.X
			y += f.state.cameraPos.Y

			if shapes.Circle(x, y, 0, 0, 0.2) {
				return "x"
			}

			if calculator.Eq(y, 0, 0.005) {
				return "."
			}

			if calculator.Eq(x, 0, 0.005) {
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
