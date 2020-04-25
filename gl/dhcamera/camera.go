package dhcamera

import (
	"github.com/dabasan/go-dh3dbasis/dhmatrix"
	"github.com/dabasan/go-dh3dbasis/dhvector"
	"github.com/dabasan/go-dhtool/dhmath"
	"github.com/dabasan/goglf/gl/dhmatrixtool"
	"github.com/dabasan/goglf/gl/dhshader"
	"github.com/dabasan/goglf/gl/dhwindow"
)

const PERSPECTIVE = iota
const ORTHOGRAPHIC = iota

type Camera struct {
	near float32
	far  float32

	camera_mode int
	fov         float32
	size        float32

	aspect float32

	position dhvector.Vector
	target   dhvector.Vector
	up       dhvector.Vector

	projection_matrix          dhmatrix.Matrix
	view_transformation_matrix dhmatrix.Matrix

	programs []*dhshader.ShaderProgram
}

func NewCamera() *Camera {
	camera := new(Camera)

	camera.near = 1.0
	camera.far = 1000.0

	camera.camera_mode = PERSPECTIVE
	camera.fov = dhmath.DegToRad(60.0)
	camera.size = 10.0

	camera.aspect = dhwindow.DEFAULT_WIDTH / dhwindow.DEFAULT_HEIGHT

	camera.position = dhvector.VGet(-50.0, 50.0, -50.0)
	camera.target = dhvector.VGet(0.0, 0.0, 0.0)
	camera.up = dhvector.VGet(0.0, 1.0, 0.0)

	camera.programs = make([]*dhshader.ShaderProgram, 0)

	return camera
}

func (c *Camera) AddProgram(program *dhshader.ShaderProgram) {
	c.programs = append(c.programs, program)
}
func (c *Camera) RemoveAllPrograms() {
	c.programs = make([]*dhshader.ShaderProgram, 0)
}

func (c *Camera) SetCameraNearFar(near float32, far float32) {
	c.near = near
	c.far = far
}
func (c *Camera) SetCameraPositionAndTarget(position dhvector.Vector, target dhvector.Vector) {
	c.position = position
	c.target = target
}
func (c *Camera) SetCameraUpVector(up dhvector.Vector) {
	c.up = up
}
func (c *Camera) GetCameraPosition() dhvector.Vector {
	return c.position
}
func (c *Camera) GetCameraTarget() dhvector.Vector {
	return c.target
}
func (c *Camera) GetCameraFrontVector() dhvector.Vector {
	front := dhvector.VSub(c.target, c.position)
	front = dhvector.VNorm(front)

	return front
}
func (c *Camera) GetCameraUpVector() dhvector.Vector {
	return c.up
}

func (c *Camera) SetupCamera_Perspective(fov float32) {
	c.projection_matrix = dhmatrixtool.GetPerspectiveMatrix_FovAndAspect(fov, c.aspect, c.near, c.far)

	c.camera_mode = PERSPECTIVE
	c.fov = fov
}
func (c *Camera) SetupCamera_Ortho(size float32) {
	c.projection_matrix = dhmatrixtool.GetOrthogonalMatrix(-size, size, -size, size, c.near, c.far)

	c.camera_mode = ORTHOGRAPHIC
	c.size = size
}

func (c *Camera) SetCameraViewMatrix(m dhmatrix.Matrix) {
	c.view_transformation_matrix = m
}

func (c *Camera) GetProjectionMatrix() dhmatrix.Matrix {
	var ret dhmatrix.Matrix

	if c.camera_mode == PERSPECTIVE {
		ret = dhmatrixtool.GetPerspectiveMatrix_FovAndAspect(c.fov, c.aspect, c.near, c.far)
	} else {
		ret = dhmatrixtool.GetOrthogonalMatrix(-c.size, c.size, -c.size, c.size, c.near, c.far)
	}

	return ret
}
func (c *Camera) GetViewTransformationMatrix() dhmatrix.Matrix {
	ret := dhmatrixtool.GetViewTransformationMatrix(c.position, c.target, c.up)
	return ret
}

func (c *Camera) UpdateAspect(width int, height int) {
	c.aspect = float32(width) / float32(height)
}
func (c *Camera) Update() {
	if c.camera_mode == PERSPECTIVE {
		c.projection_matrix = dhmatrixtool.GetPerspectiveMatrix_FovAndAspect(c.fov, c.aspect, c.near, c.far)
	}

	if c.view_transformation_matrix == nil {
		c.view_transformation_matrix = dhmatrixtool.GetViewTransformationMatrix(c.position, c.target, c.up)
	}

	for _, program := range c.programs {
		program.Enable()
		program.SetUniformVector("camera_position", c.position)
		program.SetUniformVector("camera_target", c.target)
		program.SetUniformMatrix("projection", true, c.projection_matrix)
		program.SetUniformMatrix("view_transformation", true, c.view_transformation_matrix)
		program.SetUniform1f("near", true, c.near)
		program.SetUniform1f("far", true, c.far)
		program.Disable()
	}

	c.view_transformation_matrix = nil
}
