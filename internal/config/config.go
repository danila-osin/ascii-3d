package config

import (
	"time"
)

type Config struct {
	ScreenHeight int
	ScreenWidth  int
	FrameRate    int
	FrameTime    time.Duration
}

func New(screenHeight, screenWidth, frameRate int) Config {
	config := Config{
		ScreenHeight: screenHeight,
		ScreenWidth:  screenWidth,
		FrameRate:    frameRate,
		FrameTime:    time.Second / time.Duration(frameRate),
	}

	return config
}
