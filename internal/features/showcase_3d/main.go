package showcase_3d

import (
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/pkg/screen"
)

var defaultColors = []string{" ", ".", ":", "!", "/", "r", "(", "l", "1", "Z", "4", "H", "9", "W", "8", "$", "@"}

type state struct {
	colors     []string
	colorsSize int
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
	}

	return Showcase3D{
		config: config,
		screen: screen,
		state:  st,
	}
}

func (s Showcase3D) Run() {}

func (s Showcase3D) startRenderLoop() {
}
