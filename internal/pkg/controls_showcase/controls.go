package controls_showcase

import (
	"fmt"
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/pkg/controls"
	"github.com/danila-osin/ascii-3d/pkg/geometry"
)

func createControls(config config.Config, offset *geometry.Vec2[int]) *controls.Controls {
	actions := []controls.Action{
		{
			Keys:        []string{"a"},
			Description: "+X offset",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					offset.X -= 1
				},
			},
		},
		{
			Keys:        []string{"d"},
			Description: "-X offset",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					offset.X += 1
				},
			},
		},
		{
			Keys:        []string{"w"},
			Description: "-Y offset",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					offset.Y -= 1
				},
			},
		},
		{
			Keys:        []string{"s"},
			Description: "+Y offset",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					offset.Y += 1
				},
			},
		},
		{
			Keys:        []string{"+"},
			Description: "ZoomIn",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					fmt.Println("Zoom In")
				},
			},
		},
		{
			Keys:        []string{"-"},
			Description: "ZoomOut",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					fmt.Println("Zoom Out")
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
