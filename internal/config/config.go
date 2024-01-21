package config

import (
	"time"
)

type Config struct {
	ScreenHeight int
	ScreenWidth  int
	Framerate    int
	FrameTime    time.Duration
	FontAspect   float64
}

func New(screenHeight, screenWidth, framerate int, fontAspect float64) Config {
	config := Config{
		ScreenHeight: screenHeight,
		ScreenWidth:  screenWidth,
		Framerate:    framerate,
		FrameTime:    time.Second / time.Duration(framerate),
		FontAspect:   fontAspect,
	}

	return config
}
