package controls

import (
	"fmt"
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"strings"
)

const (
	descriptionTemplate = "| %s: (%s)%%s |" // | MoveLeft: (a, A) |
)

type Descriptions struct {
	text []string
	size DescriptionsSize
}

func newDescriptions(actions []Action) Descriptions {
	text := getText(actions)
	size := DescriptionsSize{H: len(text), W: len(text[0])}

	return Descriptions{
		text: text,
		size: size,
	}
}

func getText(actions []Action) []string {
	maxLen := 0
	descriptions := make([]string, len(actions)+2)

	for i := range actions {
		keyDescription := fmt.Sprintf(descriptionTemplate, actions[i].Description, strings.Join(actions[i].Keys, ", "))

		descriptions[i+1] = keyDescription

		if kdLen := len(keyDescription); kdLen > maxLen {
			maxLen = kdLen
		}
	}

	for i := range actions {
		if diff := maxLen - len(descriptions[i+1]); diff >= 0 {
			descriptions[i+1] = fmt.Sprintf(descriptions[i+1], strings.Repeat(" ", diff))
		}
	}

	horizontalLine := strings.Repeat("-", len(descriptions[1]))

	descriptions[0] = horizontalLine
	descriptions[len(descriptions)-1] = horizontalLine

	return descriptions
}

func (d Descriptions) Size() DescriptionsSize {
	return d.size
}

func (d Descriptions) Text() []string {
	return d.text
}

type DescriptionsSize struct {
	H, W int
}

func (ds DescriptionsSize) Vec2Int() geometry.Vec2[int] {
	return geometry.Vec2[int]{
		X: ds.W,
		Y: ds.H,
	}
}
