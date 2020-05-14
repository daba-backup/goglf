package util

import (
	"math"

	"github.com/dabasan/go-dh3dbasis/vector"
	dhmath "github.com/dabasan/go-dhtool/math"
	"github.com/dabasan/goglf/gl/front"
)

var camera_vrot_min float32
var camera_vrot_max float32

func init() {
	camera_vrot_min = dhmath.DegToRad(-80.0)
	camera_vrot_max = dhmath.DegToRad(80.0)
}

type FreeCamera struct {
	position vector.Vector
	v_rotate float32
	h_rotate float32

	translate_speed float32
	rotate_speed    float32
}

func NewFreeCamera() *FreeCamera {
	camera := new(FreeCamera)
	camera.position = vector.VGet(50.0, 50.0, 50.0)
	camera.v_rotate = 0.0
	camera.h_rotate = 0.0
	camera.translate_speed = 0.3
	camera.rotate_speed = 0.01

	return camera
}

func (c *FreeCamera) GetPosition() vector.Vector {
	return c.position
}
func (c *FreeCamera) GetRotation() (float32, float32) {
	return c.v_rotate, c.h_rotate
}

func (c *FreeCamera) SetPosition(position vector.Vector) {
	c.position = position
}
func (c *FreeCamera) SetRotation(v_rotate float32, h_rotate float32) {
	c.v_rotate = v_rotate
	c.h_rotate = h_rotate
}
func (c *FreeCamera) SetTranslateSpeed(translate_speed float32) {
	c.translate_speed = translate_speed
}
func (c *FreeCamera) SetRotateSpeed(rotate_speed float32) {
	c.rotate_speed = rotate_speed
}

func (c *FreeCamera) Translate(front int, back int, right int, left int) {
	translate := vector.VGet(0.0, 0.0, 0.0)

	front_vec := vector.VGetFromAngles(c.v_rotate, c.h_rotate)
	right_vec := vector.VCross(front_vec, vector.VGet(0.0, 1.0, 0.0))
	right_vec = vector.VNorm(right_vec)

	if front > 0 {
		translate = vector.VAdd(translate, front_vec)
	}
	if back > 0 {
		translate = vector.VAdd(translate, vector.VScale(front_vec, -1.0))
	}
	if right > 0 {
		translate = vector.VAdd(translate, right_vec)
	}
	if left > 0 {
		translate = vector.VAdd(translate, vector.VScale(right_vec, -1.0))
	}

	if vector.VSize(translate) > 1.0E-8 {
		translate = vector.VScale(translate, c.translate_speed)
		c.position = vector.VAdd(c.position, translate)
	}
}
func (c *FreeCamera) Rotate(diff_x float64, diff_y float64) {
	c.h_rotate += c.rotate_speed * float32(-diff_x)
	c.v_rotate += c.rotate_speed * float32(-diff_y)

	if c.h_rotate > math.Pi {
		c.h_rotate -= 2.0 * math.Pi
	} else if c.h_rotate < (-math.Pi) {
		c.h_rotate += 2.0 * math.Pi
	}

	if c.v_rotate < camera_vrot_min {
		c.v_rotate = camera_vrot_min
	} else if c.v_rotate > camera_vrot_max {
		c.v_rotate = camera_vrot_max
	}
}

func (c *FreeCamera) Update() {
	front.SetCameraPositionAndAngle(c.position, c.v_rotate, c.h_rotate, 0.0)
}
