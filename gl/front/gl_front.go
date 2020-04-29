package front

import (
	"errors"
	"log"

	"github.com/dabasan/goglf/gl/shader"
	"github.com/dabasan/goglf/gl/wrapper"

	"github.com/go-gl/gl/all-core/gl"
)

func Initialize() error {
	err := loadDefaultShaders()
	if err != nil {
		return err
	}

	setDefaultProperties()
	addProgramsToFront()

	return nil
}
func loadDefaultShaders() error {
	r1 := shader.CreateProgram(
		"texture",
		"./Data/Shader/330/texture/gouraud/vshader.glsl",
		"./Data/Shader/330/texture/gouraud/fshader.glsl")
	r2 := shader.CreateProgram(
		"color",
		"./Data/Shader/330/color/vshader.glsl",
		"./Data/Shader/330/color/fshader.glsl")
	r3 := shader.CreateProgram(
		"texture_drawer",
		"./Data/Shader/330/texture_drawer/vshader.glsl",
		"./Data/Shader/330/texture_drawer/fshader.glsl")
	r4 := shader.CreateProgram(
		"simple_2d",
		"./Data/Shader/330/simple_2d/vshader.glsl",
		"./Data/Shader/330/simple_2d/fshader.glsl")

	if r1 != 0 || r2 != 0 || r3 != 0 || r4 != 0 {
		log.Printf("error: Failed to create a program(s).")
		return errors.New("Failed to create a program(s).")
	}

	log.Printf("info: Default shaders loaded.")
	return nil
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
