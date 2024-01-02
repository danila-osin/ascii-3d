package screen

import (
	"fmt"
	"github.com/danila-osin/ascii-3d/internal/config"
	"os"
	"os/exec"
	"strings"
	"time"
)

const EmptyPixel = "."
const PixelSeparator = " "

type Screen struct {
	config config.Config

	Matrix Matrix
	Height int
	Width  int
}

type iteratorFunc func(y, x int, value string)
type iteratorSetFunc func(y, x int, value string) string

func New(config config.Config) *Screen {
	return &Screen{
		config: config,
		Matrix: newMatrix(config.ScreenHeight, config.ScreenWidth),
		Height: config.ScreenHeight,
		Width:  config.ScreenWidth,
	}
}

func (s Screen) Render(clear bool) {
	if clear {
		s.Clear()
	}

	for y := 0; y < s.Height; y++ {
		fmt.Println(strings.Join(s.Matrix[y], PixelSeparator))
	}
}

func (s Screen) StartRenderLoop(clear bool, beforeRenderFn func(), afterRenderFn func()) {
	for {
		beforeRenderFn()
		s.Render(clear)
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

func newMatrix(height, width int) Matrix {
	var matrix Matrix

	for y := 0; y < height; y++ {
		matrix = append(matrix, []string{})
		for x := 0; x < width; x++ {
			matrix[y] = append(matrix[y], EmptyPixel)
		}
	}

	return matrix
}
