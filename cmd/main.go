package main

import (
	"flag"
	"fmt"
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/internal/pkg/controls_showcase"
	"github.com/danila-osin/ascii-3d/internal/pkg/function_graph"
	"github.com/danila-osin/ascii-3d/internal/pkg/game_of_life"
	"github.com/danila-osin/ascii-3d/pkg/screen"
)

type appFlags struct {
	screenHeight int
	screenWidth  int
	frameRate    int
	mode         string
}

func main() {
	flags := parseFlags()

	c := config.New(flags.screenHeight, flags.screenWidth, flags.frameRate)
	s := screen.New(c, "", " ")

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
	default:
		fmt.Println("Unknown Mode '" + flags.mode + "'")
	}
}

func parseFlags() appFlags {
	screenHeight := flag.Int("h", 50, "Screen Height")
	screenWidth := flag.Int("w", 50, "Screen Width")
	frameRate := flag.Int("fr", 20, "Frame Rate")
	mode := flag.String("m", "unknown", "App Mode [life, graph]")

	flag.Parse()

	return appFlags{
		screenHeight: *screenHeight,
		screenWidth:  *screenWidth,
		frameRate:    *frameRate,
		mode:         *mode,
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
