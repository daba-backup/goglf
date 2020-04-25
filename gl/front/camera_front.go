package front

import (
	"math"

	"github.com/dabasan/go-dh3dbasis/matrix"
	"github.com/dabasan/go-dh3dbasis/vector"
	dhmath "github.com/dabasan/go-dhtool/math"
	"github.com/dabasan/goglf/gl/camera"
	"github.com/dabasan/goglf/gl/common"
	"github.com/dabasan/goglf/gl/coordinatetool"
	"github.com/dabasan/goglf/gl/shader"
)

var c *camera.Camera
var window_width int
var window_height int

func init() {
	c = camera.NewCamera()
	window_width = common.WINDOW_DEFAULT_WIDTH
	window_height = common.WINDOW_DEFAULT_HEIGHT
}

func AddProgramToCamera(program *shader.ShaderProgram) {
	c.AddProgram(program)
}
func RemoveAllProgramsFromCamera() {
	c.RemoveAllPrograms()
}

func SetCameraNearFar(near float32, far float32) {
	c.SetCameraNearFar(near, far)
}
func SetCameraPositionAndTarget_UpVecY(position vector.Vector, target vector.Vector) {
	c.SetCameraPositionAndTarget(position, target)
	c.SetCameraUpVector(vector.VGet(0.0, 1.0, 0.0))
}
func SetCameraPositionAndTargetAndUpVec(position vector.Vector, target vector.Vector, up vector.Vector) {
	c.SetCameraPositionAndTarget(position, target)
	c.SetCameraUpVector(up)
}
func SetCameraViewMatrix(m matrix.Matrix) {
	c.SetCameraViewMatrix(m)
}
func SetCameraPositionAndAngle(position vector.Vector, v_rotate float32, h_rotate float32, t_rotate float32) {
	direction := vector.VGet(0.0, 0.0, 0.0)
	direction.X = float32(math.Cos(float64(h_rotate)))
	direction.Y = float32(math.Sin(float64(v_rotate)))
	direction.Z = float32(math.Sin(-float64(h_rotate)))
	direction = vector.VNorm(direction)

	target := vector.VAdd(position, direction)

	rot_direction := matrix.MGetRotAxis(direction, t_rotate)
	up := matrix.VTransform(vector.VGet(0.0, 1.0, 0.0), rot_direction)

	c.SetCameraPositionAndTarget(position, target)
	c.SetCameraUpVector(up)
}

func SetupCamera_Perspective(fov float32) {
	c.SetupCamera_Perspective(fov)
}
func SetupCamera_Ortho(size float32) {
	c.SetupCamera_Ortho(size)
}

func ConvertWorldPosToScreenPos(world_pos vector.Vector) vector.Vector {
	projection := c.GetProjectionMatrix()
	view_transformation := c.GetViewTransformationMatrix()

	camera_matrix := matrix.MMult(projection, view_transformation)

	world_pos_matrix := matrix.MGetZero()
	world_pos_matrix.M[0][0] = world_pos.X
	world_pos_matrix.M[1][0] = world_pos.Y
	world_pos_matrix.M[2][0] = world_pos.Z
	world_pos_matrix.M[3][0] = 1.0
	clip_space_matrix := matrix.MMult(camera_matrix, world_pos_matrix)
	w := clip_space_matrix.M[3][0]

	ret := vector.VGet(clip_space_matrix.M[0][0], clip_space_matrix.M[1][0], clip_space_matrix.M[2][0])
	ret = vector.VScale(ret, 1.0/w)

	x := ret.X
	y := ret.Y

	x = float32(coordinatetool.ExpandNormalizedCoordinate(x, window_width))
	y = float32(coordinatetool.ExpandNormalizedCoordinate(y, window_height))
	ret.X = x
	ret.Y = y

	return ret
}
func ConvertScreenPosToWorldPos(screen_pos vector.Vector) vector.Vector {
	x := coordinatetool.NormalizeCoordinate_Float32(screen_pos.X, float32(window_width))
	y := coordinatetool.NormalizeCoordinate_Float32(screen_pos.Y, float32(window_height))
	z := dhmath.Clamp_float32(screen_pos.Z, -1.0, 1.0)
	normalized_screen_pos := vector.VGet(x, y, z)

	projection := c.GetProjectionMatrix()
	view_transformation := c.GetViewTransformationMatrix()

	camera_matrix := matrix.MMult(projection, view_transformation)
	inv_camera_matrix := matrix.MInverse(camera_matrix)

	clip_space_matrix := matrix.MGetZero()
	clip_space_matrix.M[0][0] = normalized_screen_pos.X
	clip_space_matrix.M[1][0] = normalized_screen_pos.Y
	clip_space_matrix.M[2][0] = normalized_screen_pos.Z
	clip_space_matrix.M[3][0] = 1.0
	world_pos_matrix := matrix.MMult(inv_camera_matrix, clip_space_matrix)
	w := world_pos_matrix.M[3][0]

	ret := vector.VGet(world_pos_matrix.M[0][0], world_pos_matrix.M[1][0], world_pos_matrix.M[2][0])
	ret = vector.VScale(ret, 1.0/w)

	return ret
}

func UpdateAspect(width int, height int) {
	c.UpdateAspect(width, height)
	window_width = width
	window_height = height
}
func UpdateCamera() {
	c.Update()
}
