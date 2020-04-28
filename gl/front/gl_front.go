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
		"simple_2d",
		"./Data/Shader/330/simple_2d/vshader.glsl",
		"./Data/Shader/330/simple_2d/fshader.glsl")

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
	texture, _ := shader.NewShaderProgram("texture")
	color, _ := shader.NewShaderProgram("color")

	AddProgramToCamera(texture)
	AddProgramToCamera(color)
	AddProgramToFog(texture)
	AddProgramToFog(color)
	AddProgramToLighting(texture)
}
