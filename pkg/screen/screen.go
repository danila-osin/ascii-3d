package screen

import (
	"fmt"
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Screen struct {
	config config.Config

	Matrix         Matrix
	Size           Size
	PixelSeparator string
	EmptyPixel     string
	Aspect         float64
}

func New(config config.Config, pixelSeparator, emptyPixel string) *Screen {
	size := Size{W: config.ScreenWidth, H: config.ScreenHeight}

	return &Screen{
		config: config,

		Matrix:         newMatrix(size, emptyPixel),
		Size:           size,
		PixelSeparator: pixelSeparator,
		EmptyPixel:     emptyPixel,
		Aspect:         float64(size.W) / float64(size.H),
	}
}

func (s Screen) Render(clear bool, byLine bool) {
	if clear {
		s.Clear()
	}

	var matrixText string
	for y := 0; y < s.Size.H; y++ {
		lineText := strings.Join(s.Matrix[y], s.PixelSeparator)

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

func (s Screen) StartRenderLoop(clear bool, beforeRenderFn func(), afterRenderFn func()) {
	for {
		beforeRenderFn()
		s.Render(clear, true)
		afterRenderFn()

		time.Sleep(s.config.FrameTime)
	}
}

func (s Screen) Iterate(iteratorFn iteratorFunc) {
	for y := 0; y < s.Size.H; y++ {
		for x := 0; x < s.Size.W; x++ {
			iteratorFn(geometry.Vec2[int]{X: x, Y: y}, s.Matrix[y][x])
		}
	}
}

func (s Screen) Set(position geometry.Vec2[int], value string) {
	s.Matrix[position.Y][position.X] = value
}

func (s Screen) IterateAndSet(iteratorFn iteratorSetFunc) {
	for y := 0; y < s.Size.H; y++ {
		for x := 0; x < s.Size.W; x++ {
			setPosition := geometry.Vec2[int]{X: x, Y: y}

			s.Set(setPosition, iteratorFn(setPosition, s.Matrix[y][x]))
		}
	}
}

func (s Screen) AddText(from geometry.Vec2[int], text []string) {
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

func (s Screen) Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		panic(err.Error())
	}
}

// Matrix TODO create separate package matrix
type Matrix [][]string

func newMatrix(size Size, emptyPixel string) Matrix {
	var matrix Matrix

	for y := 0; y < size.H; y++ {
		matrix = append(matrix, []string{})
		for x := 0; x < size.W; x++ {
			matrix[y] = append(matrix[y], emptyPixel)
		}
	}

	return matrix
}
