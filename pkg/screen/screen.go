package screen

import (
	"fmt"
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Screen struct {
	props  Props
	matrix matrix

	PixelSeparator string
	EmptyPixel     string
	Aspect         float64
}

func New(props Props, pixelSeparator, emptyPixel string) *Screen {
	return &Screen{
		props:          props,
		matrix:         newMatrix(props.Size, emptyPixel),
		PixelSeparator: pixelSeparator,
		EmptyPixel:     emptyPixel,
		Aspect:         float64(props.Size.W) / float64(props.Size.H),
	}
}

func (s *Screen) Size() Size {
	return s.props.Size
}

func (s *Screen) Render(clear bool, byLine bool) {
	if clear {
		s.Clear()
	}

	var matrixText string
	for y := 0; y < s.props.Size.H; y++ {
		lineText := strings.Join(s.matrix[y], s.PixelSeparator)

		if byLine {
			fmt.Println(lineText)
		} else {
			matrixText = matrixText + lineText + "\n"
		}

	}

	if !byLine {
		fmt.Print(matrixText)
	}
}

func (s *Screen) StartRenderLoop(clear bool, beforeRenderFn *BRenderFn, afterRenderFn *ARenderFn) {
	for {
		if beforeRenderFn != nil {
			(*beforeRenderFn)()
		}

		s.Render(clear, true)

		if afterRenderFn != nil {
			(*afterRenderFn)()
		}

		time.Sleep(s.props.FrameTime)
	}
}

// Iterate TODO Create IterateRange method
func (s *Screen) Iterate(iteratorFn iteratorFn) {
	for y := 0; y < s.props.Size.H; y++ {
		for x := 0; x < s.props.Size.W; x++ {
			iteratorFn(geometry.Vec2[int]{X: x, Y: y}, s.matrix[y][x])
		}
	}
}

func (s *Screen) Set(pos geometry.Vec2[int], val string) {
	if pos.Y >= s.props.Size.H || pos.X >= s.props.Size.W {
		return
	}

	s.matrix[pos.Y][pos.X] = val
}

func (s *Screen) Get(pos geometry.Vec2[int]) string {
	return s.matrix[pos.Y][pos.X]
}

func (s *Screen) IterateAndSet(iteratorFn iteratorSetFn) {
	for y := 0; y < s.props.Size.H; y++ {
		for x := 0; x < s.props.Size.W; x++ {
			pos := geometry.Vec2[int]{X: x, Y: y}

			s.Set(pos, iteratorFn(pos, s.matrix[y][x]))
		}
	}
}

func (s *Screen) AddText(from geometry.Vec2[int], text []string) {
	for lineIdx, line := range text {
		for symbolIdx, symbolRune := range line {
			setPosition := geometry.Vec2[int]{
				X: symbolIdx + from.X,
				Y: lineIdx + from.Y,
			}

			s.Set(setPosition, string(symbolRune))
		}
	}
}

func (s *Screen) Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		panic(err.Error())
	}
}
