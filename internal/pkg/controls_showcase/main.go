package controls_showcase

import (
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/pkg/controls"
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"github.com/danila-osin/ascii-3d/pkg/screen"
	"time"
)

type State struct {
	offset geometry.Vec2[int]
}

type ControlsShowcase struct {
	config   config.Config
	screen   *screen.Screen
	controls *controls.Controls
	state    *State
}

func New(config config.Config, screen *screen.Screen) ControlsShowcase {
	state := &State{offset: geometry.Vec2[int]{X: 0, Y: 0}}

	return ControlsShowcase{
		config:   config,
		screen:   screen,
		controls: createControls(config, state),
		state:    state,
	}
}

func (cs ControlsShowcase) Run() {
	cs.setInitialState()
	cs.screen.AddText(geometry.Vec2[int]{X: 5, Y: 5}, cs.controls.Descriptions.Text())
	cs.screen.Render(true, false)
	time.Sleep(2 * time.Second)

	go cs.controls.Listen()
	cs.startRenderLoop()
}

func (cs ControlsShowcase) setInitialState() {
	cs.screen.IterateAndSet(func(rawCursor geometry.Vec2[int], value string) string {
		return "."
	})
}

func (cs ControlsShowcase) startRenderLoop() {
	cs.screen.StartRenderLoop(true, func() {
		cs.screen.IterateAndSet(func(rawCursor geometry.Vec2[int], value string) string {
			if (rawCursor.X+cs.state.offset.X)%5 == 0 {
				return "x"
			}

			if (rawCursor.Y+cs.state.offset.Y)%5 == 0 {
				return "x"
			}

			return cs.screen.EmptyPixel
		})

		cs.screen.AddText(geometry.Vec2[int]{X: 5, Y: 5}, cs.controls.Descriptions.Text())
	}, func() {})
}
