package controls

import (
	"fmt"
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/eiannone/keyboard"
	"slices"
	"strings"
)

type Controls struct {
	config      config.Config
	runeActions []RuneAction
	keyActions  []KeyAction

	Closed       bool
	Descriptions []string
}

func New(config config.Config, runeActions []RuneAction, keyActions []KeyAction) *Controls {
	return &Controls{
		config:      config,
		runeActions: runeActions,
		keyActions:  keyActions,

		Descriptions: createDescriptions(runeActions, keyActions),
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

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		for _, action := range c.runeActions {
			if slices.Contains(action.Runes, char) {
				for _, handler := range action.Handlers {
					handler(c)
				}
			}
		}

		for _, action := range c.keyActions {
			if slices.Contains(action.Keys, key) {
				for _, handler := range action.Handlers {
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

func createDescriptions(runeActions []RuneAction, _ []KeyAction) []string {
	var descriptions []string

	maxLen := 0
	for _, action := range runeActions {
		var descriptionKeys []string
		for _, r := range action.Runes {
			descriptionKeys = append(descriptionKeys, string(r))
		}

		description := fmt.Sprintf("%s: (%s)", action.Description, strings.Join(descriptionKeys, ", "))
		descriptions = append(descriptions, description)

		if descLen := len(description); descLen > maxLen {
			maxLen = descLen
		}
	}

	for i, description := range descriptions {
		if diff := maxLen - len(description); diff > 0 {
			tail := strings.Repeat(" ", diff)
			description = description + tail
		}

		descriptions[i] = fmt.Sprintf("| %s |", description)
	}

	line := strings.Repeat("-", maxLen+4)
	return append(append([]string{line}, descriptions...), line)
}
