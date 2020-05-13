package window

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type keyCountsAndFlags struct {
	pressing_counts  map[glfw.Key]int
	releasing_counts map[glfw.Key]int
	pressing_flags   map[glfw.Key]bool
}
type mouseButtonCountsAndFlags struct {
	pressing_counts  map[glfw.MouseButton]int
	releasing_counts map[glfw.MouseButton]int
	pressing_flags   map[glfw.MouseButton]bool
}

func newKeyCountsAndFlags() *keyCountsAndFlags {
	ret := new(keyCountsAndFlags)
	ret.pressing_counts = make(map[glfw.Key]int)
	ret.releasing_counts = make(map[glfw.Key]int)
	ret.pressing_flags = make(map[glfw.Key]bool)

	return ret
}
func (cf *keyCountsAndFlags) reset() {
	cf.pressing_counts = make(map[glfw.Key]int)
	cf.releasing_counts = make(map[glfw.Key]int)
	cf.pressing_flags = make(map[glfw.Key]bool)
}

func newMouseButtonCountsAndFlags() *mouseButtonCountsAndFlags {
	ret := new(mouseButtonCountsAndFlags)
	ret.pressing_counts = make(map[glfw.MouseButton]int)
	ret.releasing_counts = make(map[glfw.MouseButton]int)
	ret.pressing_flags = make(map[glfw.MouseButton]bool)

	return ret
}
func (cf *mouseButtonCountsAndFlags) reset() {
	cf.pressing_counts = make(map[glfw.MouseButton]int)
	cf.releasing_counts = make(map[glfw.MouseButton]int)
	cf.pressing_flags = make(map[glfw.MouseButton]bool)
}
