package main

import (
	"flag"
	"fmt"
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/internal/pkg/game_of_life"
	"github.com/danila-osin/ascii-3d/internal/screen"
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
	s := screen.New(c)

	switch flags.mode {
	case "life":
		runLife(c, s)
		return
	default:
		fmt.Println("Unknown Mode '" + flags.mode + "'")
	}
}

func parseFlags() appFlags {
	screenHeight := flag.Int("h", 50, "Screen Height")
	screenWidth := flag.Int("w", 50, "Screen Width")
	frameRate := flag.Int("fr", 20, "Frame Rate")
	mode := flag.String("m", "unknown", "App Mode [life]")

	flag.Parse()

	return appFlags{
		screenHeight: *screenHeight,
		screenWidth:  *screenWidth,
		frameRate:    *frameRate,
		mode:         *mode,
	}
}

func runLife(c config.Config, s *screen.Screen) {
	gameOfLife := game_of_life.New(c, s, ".", "X")
	gameOfLife.Run()
}