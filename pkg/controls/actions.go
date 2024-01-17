package controls

import "github.com/eiannone/keyboard"

type ActionHandler func(c *Controls)

type RuneAction struct {
	Runes       []rune
	Description string
	Handlers    []ActionHandler
}

type KeyAction struct {
	Keys        []keyboard.Key
	Description string
	Handlers    []ActionHandler
}
