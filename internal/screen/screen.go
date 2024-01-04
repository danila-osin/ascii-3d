package screen

import (
	"fmt"
	"github.com/danila-osin/ascii-3d/internal/config"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Screen struct {
	config config.Config

	Matrix         Matrix
	Height         int
	Width          int
	PixelSeparator string
	EmptyPixel     string
}

type iteratorFunc func(y, x int, value string)
type iteratorSetFunc func(y, x int, value string) string

func New(config config.Config, pixelSeparator, emptyPixel string) *Screen {
	return &Screen{
		config: config,

		Matrix:         newMatrix(config.ScreenHeight, config.ScreenWidth, emptyPixel),
		Height:         config.ScreenHeight,
		Width:          config.ScreenWidth,
		PixelSeparator: pixelSeparator,
		EmptyPixel:     emptyPixel,
	}
}

func (s Screen) Render(clear bool, byLine bool) {
	if clear {
		s.Clear()
	}

	var matrixText string
	for y := 0; y < s.Height; y++ {
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
	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
			iteratorFn(y, x, s.Matrix[y][x])
		}
	}
}

func (s Screen) IterateAndSet(iteratorFn iteratorSetFunc) {
	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
			s.Matrix[y][x] = iteratorFn(y, x, s.Matrix[y][x])
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

type Matrix [][]string

func newMatrix(height, width int, emptyPixel string) Matrix {
	var matrix Matrix

	for y := 0; y < height; y++ {
		matrix = append(matrix, []string{})
		for x := 0; x < width; x++ {
			matrix[y] = append(matrix[y], emptyPixel)
		}
	}

	return matrix
}
