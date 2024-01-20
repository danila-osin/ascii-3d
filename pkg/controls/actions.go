package controls

type ActionHandler func(c *Controls)

type Action struct {
	Keys        []string
	Description string
	Handlers    []ActionHandler
}
