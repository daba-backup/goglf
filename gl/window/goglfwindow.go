package window

import (
	"log"

	"github.com/dabasan/go-dh3dbasis/coloru8"
	"github.com/dabasan/go-dh3dbasis/vector"
	"github.com/dabasan/goglf/gl/wrapper"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/dabasan/goglf/gl/front"
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

type ReshapeFunc func(gw *GOGLFWindow, width int, height int)
type UpdateFunc func(gw *GOGLFWindow)
type DrawFunc func(gw *GOGLFWindow)

type GOGLFWindow struct {
	window *glfw.Window

	key_counts          keyCountsAndFlags
	mouse_button_counts mouseButtonCountsAndFlags

	reshape_func ReshapeFunc
	update_func  UpdateFunc
	draw_func    DrawFunc

	background_color coloru8.ColorU8
}

func NewGOGLFWindow(width int, height int, title string) (*GOGLFWindow, error) {
	w := new(GOGLFWindow)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return nil, err
	}
	log.Printf("info: OpenGL version=%v", gl.GoStr(gl.GetString(gl.VERSION)))

	window.SetKeyCallback(w.keyCallback)
	window.SetMouseButtonCallback(w.mouseButtonCallback)
	window.SetFramebufferSizeCallback(w.framebufferSizeCallback)
	w.window = window

	w.reshape_func = Reshape
	w.update_func = Update
	w.draw_func = Draw

	w.background_color = coloru8.GetColorU8FromFloat32Components(0.0, 0.0, 0.0, 1.0)

	return w, nil
}

func (gw *GOGLFWindow) keyCallback(w *glfw.Window, k glfw.Key, s int, a glfw.Action, mk glfw.ModifierKey) {
	switch a {
	case glfw.Press:
		gw.key_counts.pressing_flags[k] = true
	case glfw.Release:
		gw.key_counts.pressing_flags[k] = false
	}
}
func (gw *GOGLFWindow) mouseButtonCallback(w *glfw.Window, b glfw.MouseButton, a glfw.Action, mk glfw.ModifierKey) {
	switch a {
	case glfw.Press:
		gw.mouse_button_counts.pressing_flags[b] = true
	case glfw.Release:
		gw.mouse_button_counts.pressing_flags[b] = false
	}
}
func (gw *GOGLFWindow) framebufferSizeCallback(w *glfw.Window, width int, height int) {
	front.UpdateAspect(width, height)
	Reshape(gw, width, height)
}

func (w *GOGLFWindow) ClearDrawScreen() {
	wrapper.ClearColor(w.background_color.R, w.background_color.G, w.background_color.B, w.background_color.A)
	wrapper.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (w *GOGLFWindow) display() {
	w.updateKeyCounts()
	w.updateMouseButtonCounts()

	w.ClearDrawScreen()
	front.UpdateLighting()
	front.UpdateFog()

	Update(w)
	front.UpdateCamera()
	Draw(w)
}
func (w *GOGLFWindow) updateKeyCounts() {
	for key, val := range w.key_counts.pressing_flags {
		if _, ok := w.key_counts.pressing_counts[key]; !ok {
			w.key_counts.pressing_counts[key] = 0
		}
		if _, ok := w.key_counts.releasing_counts[key]; !ok {
			w.key_counts.releasing_counts[key] = 0
		}

		if val == true {
			w.key_counts.pressing_counts[key]++
			w.key_counts.releasing_counts[key] = 0
		} else {
			w.key_counts.releasing_counts[key]++
			w.key_counts.pressing_counts[key] = 0
		}
	}
}
func (w *GOGLFWindow) updateMouseButtonCounts() {
	for key, val := range w.mouse_button_counts.pressing_flags {
		if _, ok := w.mouse_button_counts.pressing_counts[key]; !ok {
			w.mouse_button_counts.pressing_counts[key] = 0
		}
		if _, ok := w.mouse_button_counts.releasing_counts[key]; !ok {
			w.mouse_button_counts.releasing_counts[key] = 0
		}

		if val == true {
			w.mouse_button_counts.pressing_counts[key]++
			w.mouse_button_counts.releasing_counts[key] = 0
		} else {
			w.mouse_button_counts.releasing_counts[key]++
			w.mouse_button_counts.pressing_counts[key] = 0
		}
	}
}

func (w *GOGLFWindow) Loop() {
	for !w.window.ShouldClose() {
		w.display()
		w.window.SwapBuffers()

		glfw.PollEvents()
	}
}

func Reshape(gw *GOGLFWindow, width int, height int) {

}
func Update(gw *GOGLFWindow) {
	front.SetCameraPositionAndTarget_UpVecY(
		vector.VGet(50.0, 50.0, 50.0), vector.VGet(0.0, 0.0, 0.0))
}
func Draw(gw *GOGLFWindow) {

}

func (w *GOGLFWindow) GetKeyPressingCount(k glfw.Key) int {
	val, ok := w.key_counts.pressing_counts[k]
	if ok == false {
		return -1
	} else {
		return val
	}
}
