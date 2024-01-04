package function_graph

import (
	"fmt"
	"github.com/danila-osin/ascii-3d/internal/calculator"
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/internal/screen"
	"math"
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
	//time.Sleep(2 * time.Second)

	f.startRenderLoop()
}

func (f FunctionGraph) setInitialState() {
	f.screen.IterateAndSet(func(rawY, rawX int, value string) string {
		y := float64((f.screen.Height-1)/2 - (rawY))
		x := float64((rawX) - (f.screen.Width-1)/2)

		if calculator.Eq(y, 0, 0) {
			return "."
		}

		if calculator.Eq(x, 0, 0) {
			return "."
		}

		return value
	})
}

//func (f FunctionGraph) testLInes() {
//	y0 := 2
//	x0 := 2
//
//	y1 := 20
//	x1 := 20
//
//	dx := int(math.Abs(float64(x0 - x1)))
//	dy := int(math.Abs(float64(y0 - y1)))
//
//	f.screen.IterateAndSet(func(rawY, rawX int, value string) string {
//		if calculator.Bt(rawX, x0, x1) && calculator.Bt(rawY, y0, y1) {
//
//		}
//
//		return " "
//	})
//}

func (f FunctionGraph) startRenderLoop() {
	frameCounter := 1
	a := 0.
	direction := 1.

	f.screen.StartRenderLoop(true, func() {
		f.screen.IterateAndSet(func(rawY, rawX int, value string) string {
			y := float64((f.screen.Height-1)/2 - (rawY))
			x := float64((rawX) - (f.screen.Width-1)/2)

			if functions.Circle(x, y, 5, 2, math.Abs(a*2), 0) {
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
