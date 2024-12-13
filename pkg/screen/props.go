package screen

import (
	"time"
)

type Props struct {
	Size       Size
	Framerate  int
	FrameTime  time.Duration
	FontAspect float64
}
