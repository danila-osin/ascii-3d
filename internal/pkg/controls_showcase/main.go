package controls_showcase

import (
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/pkg/controls"
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"github.com/danila-osin/ascii-3d/pkg/screen"
	"time"
)

type ControlsShowcase struct {
	config   config.Config
	screen   *screen.Screen
	controls *controls.Controls
	offset   *geometry.Vec2[int]
}

func New(config config.Config, screen *screen.Screen) ControlsShowcase {
	offset := &geometry.Vec2[int]{X: 0, Y: 0}

	return ControlsShowcase{
		config:   config,
		screen:   screen,
		controls: createControls(config, offset),
		offset:   offset,
	}
}

func (cs ControlsShowcase) Run() {
	cs.setInitialState()
	cs.screen.AddText(geometry.Vec2[int]{X: 5, Y: 5}, cs.controls.Descriptions)
	cs.screen.Render(false, false)
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
			if (rawCursor.X+cs.offset.X)%5 == 0 {
				return "x"
			}

			if (rawCursor.Y+cs.offset.Y)%5 == 0 {
				return "x"
			}

			return cs.screen.EmptyPixel
		})

		cs.screen.AddText(geometry.Vec2[int]{X: 5, Y: 5}, cs.controls.Descriptions)
	}, func() {})
}
