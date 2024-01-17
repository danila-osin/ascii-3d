package function_graph

import (
	"fmt"
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/pkg/calculator"
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"github.com/danila-osin/ascii-3d/pkg/screen"
	"math"
	"time"
)

var functions = calculator.Functions{}

type FunctionGraph struct {
	config config.Config
	screen *screen.Screen
}

func New(config config.Config, screen *screen.Screen) FunctionGraph {
	screen.PixelSeparator = " "

	return FunctionGraph{
		config: config,
		screen: screen,
	}
}

func (f FunctionGraph) Run() {
	f.setInitialState()
	f.screen.Render(true, false)
	time.Sleep(2 * time.Second)

	f.startRenderLoop()
}

func (f FunctionGraph) setInitialState() {
	f.screen.IterateAndSet(func(rawCursor geometry.Vec2[int], value string) string {
		x := float64((rawCursor.X) - (f.screen.Size.W-1)/2)
		y := float64((f.screen.Size.H-1)/2 - (rawCursor.Y))

		if calculator.Eq(y, 0, 0) {
			return "."
		}

		if calculator.Eq(x, 0, 0) {
			return "."
		}

		return value
	})
}

func (f FunctionGraph) startRenderLoop() {
	frameCounter := 1
	a := 0.
	direction := 1.

	f.screen.StartRenderLoop(true, func() {
		f.screen.IterateAndSet(func(rawCursor geometry.Vec2[int], value string) string {
			y := float64((f.screen.Size.H-1)/2-(rawCursor.Y)) + 5
			x := float64((rawCursor.X)-(f.screen.Size.W-1)/2) + 5

			if functions.Circle(x, y, 5, 2, math.Abs(a*2), 0) {
				return "x"
			}
			if functions.Circle(x, y, -5, -2, math.Abs(a*2), 0) {
				return "x"
			}

			if functions.Line(x, y, a, 0, 0) {
				return "x"
			}
			if functions.Line(x, y, -a, 0, 0) {
				return "x"
			}

			if calculator.Eq(y, 0, 0) {
				return "."
			}

			if calculator.Eq(x, 0, 0) {
				return "."
			}

			return f.screen.EmptyPixel
		})
	}, func() {
		frameCounter += 1

		if math.Abs(a) > 10 {
			direction = -direction
		}

		a += direction * 0.1
		fmt.Println(a)
	})
}
