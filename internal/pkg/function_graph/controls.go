package function_graph

import (
	"fmt"
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/pkg/controls"
)

func setupControls(config config.Config, state *State) *controls.Controls {
	actions := []controls.Action{
		{
			Keys:        []string{"w", "W"},
			Description: "Move Up",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.cameraPos.Y -= 0.1
				},
			},
		},
		{
			Keys:        []string{"s", "S"},
			Description: "Move Down",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.cameraPos.Y += 0.1
				},
			},
		},
		{
			Keys:        []string{"d", "D"},
			Description: "Move Right",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.cameraPos.X += 0.1
				},
			},
		},
		{
			Keys:        []string{"a", "A"},
			Description: "Move Left",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.cameraPos.X -= 0.1
				},
			},
		},
		{
			Keys:        []string{"="},
			Description: "Zoom In",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.scale -= state.scale / 50
				},
			},
		},
		{
			Keys:        []string{"-"},
			Description: "Zoom Out",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.scale += state.scale / 50
				},
			},
		},
		{
			Keys:        []string{"q", "Q"},
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
