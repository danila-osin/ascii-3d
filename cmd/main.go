package main

import (
	"flag"
	"fmt"
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/internal/features/controls_showcase"
	"github.com/danila-osin/ascii-3d/internal/features/function_graph"
	"github.com/danila-osin/ascii-3d/internal/features/game_of_life"
	"github.com/danila-osin/ascii-3d/internal/features/showcase_3d"
	"github.com/danila-osin/ascii-3d/pkg/screen"
)

type appFlags struct {
	screenHeight int
	screenWidth  int
	frameRate    int
	fontAspect   float64
	mode         string
}

func main() {
	flags := parseFlags()

	c := config.New(flags.screenHeight, flags.screenWidth, flags.frameRate, flags.fontAspect)
	s := screen.New(screen.Props{
		Size: screen.Size{
			W: c.ScreenWidth,
			H: c.ScreenHeight,
		},
		Framerate:  c.Framerate,
		FrameTime:  c.FrameTime,
		FontAspect: c.FontAspect,
	}, "", " ")

	switch flags.mode {
	case "life":
		runGameOfLife(c, s)
		return
	case "graph":
		runFunctionGraph(c, s)
		return
	case "controls":
		runControlsShowcase(c, s)
		return
	case "3d":
		run3dShowcase(c, s)
		return
	default:
		fmt.Println("Unknown Mode '" + flags.mode + "'")
	}
}

func runGameOfLife(c config.Config, s *screen.Screen) {
	gameOfLife := game_of_life.New(c, s, ".", "X")
	gameOfLife.Run()
}

func runFunctionGraph(c config.Config, s *screen.Screen) {
	functionGraph := function_graph.New(c, s)
	functionGraph.Run()
}

func runControlsShowcase(c config.Config, s *screen.Screen) {
	functionGraph := controls_showcase.New(c, s)
	functionGraph.Run()
}

func run3dShowcase(c config.Config, s *screen.Screen) {
	showcase3d := showcase_3d.New(c, s)
	showcase3d.Run()
}

func parseFlags() appFlags {
	screenHeight := flag.Int("h", 50, "Screen Height")
	screenWidth := flag.Int("w", 50, "Screen Width")
	frameRate := flag.Int("fr", 20, "Frame Rate")
	fontAspect := flag.Float64("fa", 0.4, "Font Aspect")
	mode := flag.String("m", "unknown", "App Mode [life, graph, controls, 3d]")

	flag.Parse()

	return appFlags{
		screenHeight: *screenHeight,
		screenWidth:  *screenWidth,
		frameRate:    *frameRate,
		mode:         *mode,
		fontAspect:   *fontAspect,
	}
}
