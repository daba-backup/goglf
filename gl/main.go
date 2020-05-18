package main

import (
	"log"
	"runtime"

	"github.com/dabasan/go-dh3dbasis/coloru8"

	"github.com/comail/colog"
	"github.com/dabasan/goglf/gl/draw"
	"github.com/dabasan/goglf/gl/util"
	"github.com/dabasan/goglf/gl/window"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type WindowFields struct {
	Camera *util.FreeCamera
}

func init() {
	runtime.LockOSThread()

	colog.SetMinLevel(colog.LTrace)
	colog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	colog.Register()
}

func main() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	//glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	w, err := window.NewGOGLFWindow(640, 480, "Window", inittestfunc)
	if err != nil {
		log.Printf("error: %v", err.Error())
		return
	}
	w.SetUpdateFunc(updatetestfunc)
	w.SetDrawFunc(drawtestfunc)
	w.SetDisposeFunc(disposetestfunc)

	for !w.GetGLFWWinow().ShouldClose() {
		glfw.PollEvents()
		w.InLoop()
		util.Sleep()
	}
}

func inittestfunc(gw *window.GOGLFWindow) {
	var fields WindowFields
	fields.Camera = util.NewFreeCamera()

	gw.SetUserData(&fields)
}
func updatetestfunc(gw *window.GOGLFWindow) {
	fields := gw.GetUserData().(*WindowFields)

	front_key := gw.GetKeyboardPressingCount(glfw.KeyW)
	back_key := gw.GetKeyboardPressingCount(glfw.KeyS)
	right_key := gw.GetKeyboardPressingCount(glfw.KeyD)
	left_key := gw.GetKeyboardPressingCount(glfw.KeyA)

	var diff_x float64
	var diff_y float64
	if gw.GetMousePressingCount(glfw.MouseButtonMiddle) > 0 {
		diff_x, diff_y = gw.GetCursorDiff()
	} else {
		diff_x, diff_y = 0.0, 0.0
	}

	fields.Camera.Translate(front_key, back_key, right_key, left_key)
	fields.Camera.Rotate(diff_x, diff_y)
	fields.Camera.Update()
}
func drawtestfunc(gw *window.GOGLFWindow) {
	draw.DrawFilledRectangle2D(5, 5, 400, 400, coloru8.GetColorU8FromFloat32Components(1.0, 0.0, 1.0, 1.0))
}
func disposetestfunc(gw *window.GOGLFWindow) {

}
