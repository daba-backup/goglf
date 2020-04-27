package util

import (
	"math"
	"time"
)

var fps float64
var spf float64

func init() {
	fps = 30.0
	spf = 1.0 / fps
}

func Sleep() {
	ms := time.Duration(math.Round(spf * 1000.0))
	time.Sleep(ms * time.Millisecond)
}

func SetFPS(f float64) {
	fps = f
	spf = 1.0 / fps
}
