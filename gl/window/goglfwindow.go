package window

import (
	"log"

	"github.com/dabasan/go-dh3dbasis/coloru8"
	"github.com/dabasan/go-dh3dbasis/vector"
	"github.com/dabasan/goglf/gl/draw"
	"github.com/dabasan/goglf/gl/wrapper"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/dabasan/goglf/gl/front"
)

type InitFunc func(gw *GOGLFWindow)
type UpdateFunc func(gw *GOGLFWindow)
type DrawFunc func(gw *GOGLFWindow)
type DisposeFunc func(gw *GOGLFWindow)

type GOGLFWindow struct {
	window    *glfw.Window
	is_hidden bool

	keyboard_caf      *keyboardCountsAndFlags
	mouse_button_caf  *mouseButtonCountsAndFlags
	last_cursor_pos_x float64
	last_cursor_pos_y float64
	cursor_diff_x     float64
	cursor_diff_y     float64
	scroll_x          float64
	scroll_y          float64

	update_func  UpdateFunc
	draw_func    DrawFunc
	dispose_func DisposeFunc

	background_color coloru8.ColorU8

	user_data interface{}
}

func NewGOGLFWindow(width int, height int, title string, init_func InitFunc) (*GOGLFWindow, error) {
	Lock()
	defer Unlock()

	gw := new(GOGLFWindow)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return nil, err
	}

	window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		return nil, err
	}
	log.Printf("info: OpenGL version=%v", gl.GoStr(gl.GetString(gl.VERSION)))

	front.Initialize()

	window.SetKeyCallback(gw.keyCallback)
	window.SetMouseButtonCallback(gw.mouseButtonCallback)
	window.SetScrollCallback(gw.scrollCallback)
	window.SetCloseCallback(gw.closeCallback)
	gw.window = window
	gw.is_hidden = false

	gw.keyboard_caf = newkeyboardCountsAndFlags()
	gw.mouse_button_caf = newMouseButtonCountsAndFlags()
	gw.last_cursor_pos_x, gw.last_cursor_pos_y = window.GetCursorPos()
	gw.cursor_diff_x = 0.0
	gw.cursor_diff_y = 0.0
	gw.scroll_x = 0.0
	gw.scroll_y = 0.0

	gw.dispose_func = Dispose
	gw.update_func = Update
	gw.draw_func = Draw

	gw.background_color = coloru8.GetColorU8FromFloat32Components(0.0, 0.0, 0.0, 1.0)

	if init_func != nil {
		init_func(gw)
	}

	return gw, nil
}

func (gw *GOGLFWindow) keyCallback(w *glfw.Window, k glfw.Key, s int, a glfw.Action, mk glfw.ModifierKey) {
	switch a {
	case glfw.Press:
		gw.keyboard_caf.pressing_flags[k] = true
	case glfw.Release:
		gw.keyboard_caf.pressing_flags[k] = false
	}
}
func (gw *GOGLFWindow) mouseButtonCallback(w *glfw.Window, b glfw.MouseButton, a glfw.Action, mk glfw.ModifierKey) {
	switch a {
	case glfw.Press:
		gw.mouse_button_caf.pressing_flags[b] = true
	case glfw.Release:
		gw.mouse_button_caf.pressing_flags[b] = false
	}
}
func (gw *GOGLFWindow) scrollCallback(w *glfw.Window, xoff float64, yoff float64) {
	gw.scroll_x = xoff
	gw.scroll_y = yoff
}
func (gw *GOGLFWindow) closeCallback(w *glfw.Window) {
	Lock()
	defer Unlock()

	gw.window.MakeContextCurrent()
	gw.dispose_func(gw)
}

func (gw *GOGLFWindow) clearDrawScreen() {
	wrapper.ClearColor(gw.background_color.R, gw.background_color.G, gw.background_color.B, gw.background_color.A)
	wrapper.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (gw *GOGLFWindow) ResetKeyboardInputState() {
	gw.keyboard_caf.reset()
}
func (gw *GOGLFWindow) ResetMouseInputState() {
	gw.mouse_button_caf.reset()
}

func (gw *GOGLFWindow) GetGLFWWinow() *glfw.Window {
	return gw.window
}

func (gw *GOGLFWindow) IsHidden() bool {
	return gw.is_hidden
}
func (gw *GOGLFWindow) ShowWindow() {
	gw.window.Show()
	gw.is_hidden = false
}
func (gw *GOGLFWindow) HideWindow() {
	gw.window.Hide()
	gw.is_hidden = true
}

func (gw *GOGLFWindow) SetBackgroundColor(color coloru8.ColorU8) {
	gw.background_color = color
}
func (gw *GOGLFWindow) GetBackgroundColor() coloru8.ColorU8 {
	return gw.background_color
}

func (gw *GOGLFWindow) SetUserData(user_data interface{}) {
	gw.user_data = user_data
}
func (gw *GOGLFWindow) GetUserData() interface{} {
	return gw.user_data
}

func (gw *GOGLFWindow) display() {
	//Update input.
	gw.updateKeyCounts()
	gw.updateMouseButtonCounts()
	gw.updateCursorProperties()

	//Default updates==========
	gw.clearDrawScreen()
	gw.updateAspect()
	front.UpdateLighting()
	front.UpdateFog()
	//====================
	gw.update_func(gw) //User update
	gw.resetScrollVols()
	front.UpdateCamera()
	gw.draw_func(gw) //User draw
}
func (gw *GOGLFWindow) updateKeyCounts() {
	for key, val := range gw.keyboard_caf.pressing_flags {
		if _, ok := gw.keyboard_caf.pressing_counts[key]; !ok {
			gw.keyboard_caf.pressing_counts[key] = 0
		}
		if _, ok := gw.keyboard_caf.releasing_counts[key]; !ok {
			gw.keyboard_caf.releasing_counts[key] = 0
		}

		if val == true {
			gw.keyboard_caf.pressing_counts[key]++
			gw.keyboard_caf.releasing_counts[key] = 0
		} else {
			gw.keyboard_caf.releasing_counts[key]++
			gw.keyboard_caf.pressing_counts[key] = 0
		}
	}
}
func (gw *GOGLFWindow) updateMouseButtonCounts() {
	for key, val := range gw.mouse_button_caf.pressing_flags {
		if _, ok := gw.mouse_button_caf.pressing_counts[key]; !ok {
			gw.mouse_button_caf.pressing_counts[key] = 0
		}
		if _, ok := gw.mouse_button_caf.releasing_counts[key]; !ok {
			gw.mouse_button_caf.releasing_counts[key] = 0
		}

		if val == true {
			gw.mouse_button_caf.pressing_counts[key]++
			gw.mouse_button_caf.releasing_counts[key] = 0
		} else {
			gw.mouse_button_caf.releasing_counts[key]++
			gw.mouse_button_caf.pressing_counts[key] = 0
		}
	}
}
func (gw *GOGLFWindow) updateCursorProperties() {
	cursor_pos_x, cursor_pos_y := gw.window.GetCursorPos()
	gw.cursor_diff_x = cursor_pos_x - gw.last_cursor_pos_x
	gw.cursor_diff_y = cursor_pos_y - gw.last_cursor_pos_y
	gw.last_cursor_pos_x = cursor_pos_x
	gw.last_cursor_pos_y = cursor_pos_y
}
func (gw *GOGLFWindow) updateAspect() {
	width, height := gw.window.GetSize()

	wrapper.Viewport(0, 0, int32(width), int32(height))

	front.UpdateCameraAspect(width, height)
	draw.SetWindowSize(width, height)
}
func (gw *GOGLFWindow) resetScrollVols() {
	gw.scroll_x = 0.0
	gw.scroll_y = 0.0
}

func (gw *GOGLFWindow) InLoop() {
	Lock()
	defer Unlock()

	gw.window.MakeContextCurrent()
	gw.display()
	gw.window.SwapBuffers()
}

func (gw *GOGLFWindow) GetKeyboardPressingCount(k glfw.Key) int {
	val, ok := gw.keyboard_caf.pressing_counts[k]
	if ok == false {
		return 0
	} else {
		return val
	}
}
func (gw *GOGLFWindow) GetkeyboardReleasingCount(k glfw.Key) int {
	val, ok := gw.keyboard_caf.releasing_counts[k]
	if ok == false {
		return 0
	} else {
		return val
	}
}
func (gw *GOGLFWindow) GetMousePressingCount(b glfw.MouseButton) int {
	val, ok := gw.mouse_button_caf.pressing_counts[b]
	if ok == false {
		return 0
	} else {
		return val
	}
}
func (gw *GOGLFWindow) GetMouseReleasingCount(b glfw.MouseButton) int {
	val, ok := gw.mouse_button_caf.releasing_counts[b]
	if ok == false {
		return 0
	} else {
		return val
	}
}
func (gw *GOGLFWindow) GetCursorDiff() (float64, float64) {
	return gw.cursor_diff_x, gw.cursor_diff_y
}
func (gw *GOGLFWindow) GetScrollVols() (float64, float64) {
	return gw.scroll_x, gw.scroll_y
}

func Init(gw *GOGLFWindow) {

}
func Update(gw *GOGLFWindow) {
	front.SetCameraPositionAndTarget_UpVecY(
		vector.VGet(50.0, 50.0, 50.0), vector.VGet(0.0, 0.0, 0.0))
}
func Draw(gw *GOGLFWindow) {
	draw.DrawAxes(100.0)
}
func Dispose(gw *GOGLFWindow) {

}

func (gw *GOGLFWindow) SetUpdateFunc(f UpdateFunc) {
	gw.update_func = f
}
func (gw *GOGLFWindow) SetDrawFunc(f DrawFunc) {
	gw.draw_func = f
}
func (gw *GOGLFWindow) SetDisposeFunc(f DisposeFunc) {
	gw.dispose_func = f
}
