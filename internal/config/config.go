package config

import (
	"time"
)

type Config struct {
	ScreenHeight int
	ScreenWidth  int
	Framerate    int
	FrameTime    time.Duration
}

func New(screenHeight, screenWidth, framerate int) Config {
	config := Config{
		ScreenHeight: screenHeight,
		ScreenWidth:  screenWidth,
		Framerate:    framerate,
		FrameTime:    time.Second / time.Duration(framerate),
	}

	return config
}
