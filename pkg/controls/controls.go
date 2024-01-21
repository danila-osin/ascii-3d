package controls

import (
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/eiannone/keyboard"
	"slices"
)

type Controls struct {
	config  config.Config
	actions []Action

	Closed       bool
	Descriptions Descriptions
}

func New(config config.Config, actions []Action) *Controls {
	return &Controls{
		config:  config,
		actions: actions,

		Descriptions: newDescriptions(actions),
	}
}

func (c *Controls) Listen() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	c.Closed = false
	for {
		if c.Closed {
			break
		}

		char, _, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		for i := range c.actions {
			if slices.Contains(c.actions[i].Keys, string(char)) {
				for _, handler := range c.actions[i].Handlers {
					handler(c)
				}
			}
		}
	}
}

func (c *Controls) Close() {
	_ = keyboard.Close()
	c.Closed = true
}
