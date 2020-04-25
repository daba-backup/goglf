package front

import (
	"log"

	"github.com/dabasan/goglf/gl/shader"
	"github.com/dabasan/goglf/gl/wrapper"

	"github.com/go-gl/gl/all-core/gl"
)

func Initialize() {
	loadDefaultShaders()
	setDefaultProperties()
	addProgramsToFront()
}
func loadDefaultShaders() {
	shader.CreateProgram(
		"texture",
		"./Data/Shader/330/texture/gouraud/vshader.glsl",
		"./Data/Shader/330/texture/gouraud/fshader.glsl")
	shader.CreateProgram(
		"color",
		"./Data/Shader/330/color/vshader.glsl",
		"./Data/Shader/330/color/fshader.glsl")
	shader.CreateProgram(
		"texture_drawer",
		"./Data/Shader/330/texture_drawer/vshader.glsl",
		"./Data/Shader/330/texture_drawer/fshader.glsl")
	shader.CreateProgram(
		"line_drawer",
		"./Data/Shader/330/line_drawer/vshader.glsl",
		"./Data/Shader/330/line_drawer/fshader.glsl")

	log.Printf("info: Default shaders loaded.")
}
func setDefaultProperties() {
	wrapper.Enable(gl.DEPTH_TEST)
	wrapper.DepthFunc(gl.LESS)

	wrapper.Enable(gl.CULL_FACE)
	wrapper.CullFace(gl.BACK)

	wrapper.Enable(gl.BLEND)
	wrapper.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	log.Printf("info: Default properties set.")
}
func addProgramsToFront() {

}
