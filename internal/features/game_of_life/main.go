package game_of_life

import (
	"fmt"
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	"github.com/danila-osin/ascii-3d/pkg/screen"
	"math"
	"math/rand"
	"time"
)

const (
	DeadCell  = "."
	AliveCell = "X"
)

type GameOfLife struct {
	config config.Config
	screen *screen.Screen

	DeadCell  string
	AliveCell string
}

func New(config config.Config, screen *screen.Screen, deadCell, aliveCell string) GameOfLife {
	screen.PixelSeparator = " "
	screen.EmptyPixel = "."

	return GameOfLife{
		config:    config,
		screen:    screen,
		DeadCell:  deadCell,
		AliveCell: aliveCell,
	}
}

func (g GameOfLife) Run() {
	g.setRandomInitialState()
	g.screen.Render(true, true)
	fmt.Println("Starting...")

	time.Sleep(2 * time.Second)

	g.startRenderLoop(func(frameCounter int) {
		fmt.Printf("Generation: %d\n", frameCounter)
	})
}

func (g GameOfLife) setRandomInitialState() {
	g.screen.IterateAndSet(func(_ geometry.Vec2[int], _ string) string {
		random := rand.Float64()
		chance := 1. / 3.

		if random < chance {
			return AliveCell
		} else {
			return DeadCell
		}
	})
}

func (g GameOfLife) startRenderLoop(frameFn func(frameCounter int)) {
	frameCounter := 1

	bRenderFn := screen.BRenderFn(func() {
		g.screen.IterateAndSet(func(cursor geometry.Vec2[int], value string) string {
			aliveCount, _ := g.countNeighbours(cursor)

			imAlive := value == AliveCell
			if imAlive {
				if aliveCount == 2 || aliveCount == 3 {
					return value
				}
			}

			if aliveCount == 3 {
				return AliveCell
			} else {
				return DeadCell
			}
		})
	})

	aRender := screen.ARenderFn(func() {
		frameFn(frameCounter)
		frameCounter += 1
	})

	g.screen.StartRenderLoop(true, &bRenderFn, &aRender)
}

func (g GameOfLife) countNeighbours(pos geometry.Vec2[int]) (alive, dead int) {
	x := pos.X
	y := pos.Y

	minI := int(math.Max(0, float64(y-1)))
	maxI := int(math.Min(float64(y+1), float64(g.screen.Size().H-1)))

	minJ := int(math.Max(0, float64(x-1)))
	maxJ := int(math.Min(float64(x+1), float64(g.screen.Size().W-1)))

	for i := minI; i <= maxI; i++ {
		for j := minJ; j <= maxJ; j++ {
			if i == y && j == x {
				continue
			}
			if g.screen.Get(geometry.Vec2[int]{X: j, Y: i}) == AliveCell {
				alive += 1
			} else {
				dead += 1
			}
		}
	}

	return alive, dead
}
