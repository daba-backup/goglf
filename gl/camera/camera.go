package camera

import (
	"github.com/dabasan/go-dh3dbasis/matrix"
	"github.com/dabasan/go-dh3dbasis/vector"
	"github.com/dabasan/go-dhtool/math"
	"github.com/dabasan/goglf/gl/common"
	"github.com/dabasan/goglf/gl/matrixtool"
	"github.com/dabasan/goglf/gl/shader"
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

	position vector.Vector
	target   vector.Vector
	up       vector.Vector

	projection_matrix          matrix.Matrix
	view_transformation_matrix matrix.Matrix
	matrix_set_flag            bool

	programs []*shader.ShaderProgram
}

func NewCamera() *Camera {
	camera := new(Camera)

	camera.near = 1.0
	camera.far = 1000.0

	camera.camera_mode = PERSPECTIVE
	camera.fov = math.DegToRad(60.0)
	camera.size = 10.0

	camera.aspect = common.WINDOW_DEFAULT_WIDTH / common.WINDOW_DEFAULT_HEIGHT

	camera.position = vector.VGet(-50.0, 50.0, -50.0)
	camera.target = vector.VGet(0.0, 0.0, 0.0)
	camera.up = vector.VGet(0.0, 1.0, 0.0)

	camera.matrix_set_flag = false

	camera.programs = make([]*shader.ShaderProgram, 0)

	return camera
}

func (c *Camera) AddProgram(program *shader.ShaderProgram) {
	c.programs = append(c.programs, program)
}
func (c *Camera) RemoveAllPrograms() {
	c.programs = make([]*shader.ShaderProgram, 0)
}

func (c *Camera) SetCameraNearFar(near float32, far float32) {
	c.near = near
	c.far = far
}
func (c *Camera) SetCameraPositionAndTarget(position vector.Vector, target vector.Vector) {
	c.position = position
	c.target = target
}
func (c *Camera) SetCameraUpVector(up vector.Vector) {
	c.up = up
}
func (c *Camera) GetCameraPosition() vector.Vector {
	return c.position
}
func (c *Camera) GetCameraTarget() vector.Vector {
	return c.target
}
func (c *Camera) GetCameraFrontVector() vector.Vector {
	front := vector.VSub(c.target, c.position)
	front = vector.VNorm(front)

	return front
}
func (c *Camera) GetCameraUpVector() vector.Vector {
	return c.up
}

func (c *Camera) SetupCamera_Perspective(fov float32) {
	c.projection_matrix = matrixtool.GetPerspectiveMatrix_FovAndAspect(fov, c.aspect, c.near, c.far)

	c.camera_mode = PERSPECTIVE
	c.fov = fov
}
func (c *Camera) SetupCamera_Ortho(size float32) {
	c.projection_matrix = matrixtool.GetOrthogonalMatrix(-size, size, -size, size, c.near, c.far)

	c.camera_mode = ORTHOGRAPHIC
	c.size = size
}

func (c *Camera) SetCameraViewMatrix(m matrix.Matrix) {
	c.view_transformation_matrix = m
	c.matrix_set_flag = true
}

func (c *Camera) GetProjectionMatrix() matrix.Matrix {
	var ret matrix.Matrix

	if c.camera_mode == PERSPECTIVE {
		ret = matrixtool.GetPerspectiveMatrix_FovAndAspect(c.fov, c.aspect, c.near, c.far)
	} else {
		ret = matrixtool.GetOrthogonalMatrix(-c.size, c.size, -c.size, c.size, c.near, c.far)
	}

	return ret
}
func (c *Camera) GetViewTransformationMatrix() matrix.Matrix {
	ret := matrixtool.GetViewTransformationMatrix(c.position, c.target, c.up)
	return ret
}

func (c *Camera) UpdateAspect(width int, height int) {
	c.aspect = float32(width) / float32(height)
}
func (c *Camera) Update() {
	if c.camera_mode == PERSPECTIVE {
		c.projection_matrix = matrixtool.GetPerspectiveMatrix_FovAndAspect(c.fov, c.aspect, c.near, c.far)
	}

	if c.matrix_set_flag == false {
		c.view_transformation_matrix = matrixtool.GetViewTransformationMatrix(c.position, c.target, c.up)
	}

	for _, program := range c.programs {
		program.Enable()
		program.SetUniformVector("camera.position", c.position)
		program.SetUniformVector("camera.target", c.target)
		program.SetUniformMatrix("camera.projection", true, c.projection_matrix)
		program.SetUniformMatrix("camera.view_transformation", true, c.view_transformation_matrix)
		program.SetUniform1f("camera.near", c.near)
		program.SetUniform1f("camera.far", c.far)
		program.Disable()
	}

	c.matrix_set_flag = false
}
