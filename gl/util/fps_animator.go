package util

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

var fps float64

var current_time float64
var last_time float64
var elapsed_time float64

func init() {
	fps = 30.0

	last_time = 0.0
	elapsed_time = 0.0
}

func UpdateTimer() {
	current_time = glfw.GetTime()
	elapsed_time = current_time - last_time
}

func IsElapsed() bool {
	if elapsed_time >= 1.0/fps {
		last_time = current_time
		elapsed_time = 0.0

		return true
	} else {
		return false
	}
}

func SetFPS(f float64) {
	fps = f
}
