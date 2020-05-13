package front

import (
	"log"

	"github.com/dabasan/goglf/gl/draw"
	"github.com/dabasan/goglf/gl/shader"
	"github.com/dabasan/goglf/gl/texture"
	"github.com/dabasan/goglf/gl/wrapper"

	"github.com/go-gl/gl/all-core/gl"
)

func Initialize() {
	loadDefaultShaders()
	setDefaultProperties()
	addProgramsToFront()

	draw.InitializeDrawFunctions2D()
	draw.InitializeDrawFunctions3D()
	texture.Initialize()
}
func loadDefaultShaders() {
	shader.CreateProgram(
		"texture",
		"./Data/Shader/330/default/texture/gouraud/vshader.glsl",
		"./Data/Shader/330/default/texture/gouraud/fshader.glsl")
	shader.CreateProgram(
		"texture2",
		"./Data/Shader/330/default/texture/phong/vshader.glsl",
		"./Data/Shader/330/default/texture/phong/fshader.glsl")
	shader.CreateProgram(
		"color",
		"./Data/Shader/330/default/color/vshader.glsl",
		"./Data/Shader/330/default/color/fshader.glsl")
	shader.CreateProgram(
		"texture_drawer",
		"./Data/Shader/330/default/texture_drawer/vshader.glsl",
		"./Data/Shader/330/default/texture_drawer/fshader.glsl")
	shader.CreateProgram(
		"simple_2d",
		"./Data/Shader/330/default/simple_2d/vshader.glsl",
		"./Data/Shader/330/default/simple_2d/fshader.glsl")
	shader.CreateProgram(
		"simple_3d",
		"./Data/Shader/330/default/simple_3d/vshader.glsl",
		"./Data/Shader/330/default/simple_3d/fshader.glsl")

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
	texture2, _ := shader.NewShaderProgram("texture2")
	color, _ := shader.NewShaderProgram("color")
	simple_3d, _ := shader.NewShaderProgram("simple_3d")

	AddProgramToCamera(texture)
	AddProgramToCamera(texture2)
	AddProgramToCamera(color)
	AddProgramToCamera(simple_3d)
	AddProgramToFog(texture)
	AddProgramToFog(texture2)
	AddProgramToFog(color)
	AddProgramToLighting(texture)
	AddProgramToLighting(texture2)
}
