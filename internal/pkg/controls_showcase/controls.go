package controls_showcase

import (
	"fmt"
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/pkg/controls"
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"github.com/eiannone/keyboard"
)

func createControls(config config.Config, offset *geometry.Vec2[int]) *controls.Controls {
	runeActions := []controls.RuneAction{
		{
			Runes:       []rune{'a'},
			Description: "+X offset",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					offset.X -= 1
				},
			},
		},
		{
			Runes:       []rune{'d'},
			Description: "-X offset",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					offset.X += 1
				},
			},
		},
		{
			Runes:       []rune{'w'},
			Description: "-Y offset",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					offset.Y -= 1
				},
			},
		},
		{
			Runes:       []rune{'s'},
			Description: "+Y offset",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					offset.Y += 1
				},
			},
		},
		{
			Runes:       []rune{'+'},
			Description: "ZoomIn",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					fmt.Println("Zoom In")
				},
			},
		},
		{
			Runes:       []rune{'-'},
			Description: "ZoomOut",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					fmt.Println("Zoom Out")
				},
			},
		},

		{
			Runes:       []rune{'q'},
			Description: "Exit",
			Handlers: []controls.ActionHandler{
				func(c *controls.Controls) {
					fmt.Println("Exit...")
					c.Close()
				},
			},
		},
	}

	keyActions := []controls.KeyAction{
		{
			Keys:        []keyboard.Key{keyboard.KeyEsc, keyboard.KeyDelete},
			Description: "Exit",
			Handlers: []controls.ActionHandler{
				func(c *controls.Controls) {
					fmt.Println("Exit...")
					c.Close()
				},
			},
		},
	}

	return controls.New(config, runeActions, keyActions)
}
