package controls_showcase

import (
	"fmt"
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/pkg/controls"
)

func createControls(config config.Config, state *State) *controls.Controls {
	actions := []controls.Action{
		{
			Keys:        []string{"a"},
			Description: "+X offset",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.offset.X -= 1
				},
			},
		},
		{
			Keys:        []string{"d"},
			Description: "-X offset",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.offset.X += 1
				},
			},
		},
		{
			Keys:        []string{"w"},
			Description: "-Y offset",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.offset.Y -= 1
				},
			},
		},
		{
			Keys:        []string{"s"},
			Description: "+Y offset",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.offset.Y += 1
				},
			},
		},
		{
			Keys:        []string{"q"},
			Description: "Exit",
			Handlers: []controls.ActionHandler{
				func(c *controls.Controls) {
					fmt.Println("Exit...")
					c.Close()
				},
			},
		},
	}

	return controls.New(config, actions)
}
