package dhdraw

import (
	"unsafe"

	"github.com/go-gl/gl/all-core/gl"

	"github.com/dabasan/go-dh3dbasis/dhcoloru8"
	"github.com/dabasan/go-dh3dbasis/dhvector"
	"github.com/dabasan/goglf/gl/dhshader"
	"github.com/dabasan/goglf/gl/dhwrapper"
)

func DrawLine3D(
	line_pos_1 dhvector.Vector, line_pos_2 dhvector.Vector,
	color_1 dhcoloru8.ColorU8, color_2 dhcoloru8.ColorU8) {
	var pos_vbo uint32
	var color_vbo uint32
	var vao uint32

	pos_buffer := make([]float32, 6)
	color_buffer := make([]float32, 8)
	pos_buffer[0] = line_pos_1.X
	pos_buffer[1] = line_pos_1.Y
	pos_buffer[2] = line_pos_1.Z
	pos_buffer[3] = line_pos_2.X
	pos_buffer[4] = line_pos_2.Y
	pos_buffer[5] = line_pos_2.Z
	color_buffer[0] = color_1.R
	color_buffer[1] = color_1.G
	color_buffer[2] = color_1.B
	color_buffer[3] = color_1.A
	color_buffer[4] = color_2.R
	color_buffer[5] = color_2.G
	color_buffer[6] = color_2.B
	color_buffer[7] = color_2.A

	dhshader.UseProgram("color")

	dhwrapper.GenBuffers(1, &pos_vbo)
	dhwrapper.GenBuffers(1, &color_vbo)

	dhwrapper.BindBuffer(gl.ARRAY_BUFFER, pos_vbo)
	dhwrapper.BufferData(gl.ARRAY_BUFFER,
		unsafe.Sizeof(float32)*len(pos_buffer), &pos_buffer, gl.STATIC_DRAW)
	dhwrapper.BindBuffer(gl.ARRAY_BUFFER, color_vbo)
	dhwrapper.BufferData(gl.ARRAY_BUFFER,
		unsafe.Sizeof(float32)*len(color_buffer), &color_buffer, gl.STATIC_DRAW)

	dhwrapper.GenVertexArrays(1, &vao)
	dhwrapper.BindVertexArray(vao)

	//Position attribute
	dhwrapper.BindBuffer(gl.ARRAY_BUFFER, pos_vbo)
	dhwrapper.EnableVertexAttribArray(0)
	dhwrapper.VertexAttribPointer(0, 3, gl.FLOAT, false, unsafe.Sizeof(float32)*3, 0)

	//Color attribute
	dhwrapper.BindBuffer(gl.ARRAY_BUFFER, color_vbo)
	dhwrapper.EnableVertexAttribArray(1)
	dhwrapper.VertexAttribPointer(1, 4, gl.FLOAT, false, unsafe.Sizeof(float32)*4, 0)

	dhwrapper.BindBuffer(gl.ARRAY_BUFFER, 0)
	dhwrapper.BindVertexArray(0)

	//Draw
	dhwrapper.BindVertexArray(vao)
	dhwrapper.Enable(gl.BLEND)
	dhwrapper.DrawArrays(gl.LINES, 0, 2)
	dhwrapper.Disable(gl.BLEND)
	dhwrapper.BindVertexArray(0)

	//Delete buffers
	dhwrapper.DeleteBuffers(1, &pos_vbo)
	dhwrapper.DeleteBuffers(1, &color_vbo)
	dhwrapper.DeleteVertexArrays(1, &vao)
}
func DrawLine3D_SingleColor(line_pos_1 dhvector.Vector, line_pos_2 dhvector.Vector, color dhcoloru8.ColorU8) {
	DrawLine3D(line_pos_1, line_pos_2, color, color)
}

func DrawAxes(length float32) {
	DrawLine3D_SingleColor(
		dhvector.VGet(-length, 0.0, 0.0),
		dhvector.Vector(length, 0.0, 0.0),
		dhcoloru8.GetColorU8FromFloat32Components(1.0, 0.0, 0.0, 1.0))
	DrawLine3D_SingleColor(
		dhvector.VGet(0.0, -length, 0.0),
		dhvector.Vector(0.0, length, 0.0),
		dhcoloru8.GetColorU8FromFloat32Components(0.0, 1.0, 0.0, 1.0))
	DrawLine3D_SingleColor(
		dhvector.VGet(0.0, 0.0, -length),
		dhvector.Vector(0.0, 0.0, length),
		dhcoloru8.GetColorU8FromFloat32Components(0.0, 0.0, 1.0, 1.0))
}
func DrawAxes_Positive(length float32) {
	DrawLine3D_SingleColor(
		dhvector.VGet(0.0, 0.0, 0.0),
		dhvector.Vector(length, 0.0, 0.0),
		dhcoloru8.GetColorU8FromFloat32Components(1.0, 0.0, 0.0, 1.0))
	DrawLine3D_SingleColor(
		dhvector.VGet(0.0, 0.0, 0.0),
		dhvector.Vector(0.0, length, 0.0),
		dhcoloru8.GetColorU8FromFloat32Components(0.0, 1.0, 0.0, 1.0))
	DrawLine3D_SingleColor(
		dhvector.VGet(0.0, 0.0, 0.0),
		dhvector.Vector(0.0, 0.0, length),
		dhcoloru8.GetColorU8FromFloat32Components(0.0, 0.0, 1.0, 1.0))
}
func DrawAxes_Negative(length float32) {
	DrawLine3D_SingleColor(
		dhvector.VGet(-length, 0.0, 0.0),
		dhvector.Vector(0.0, 0.0, 0.0),
		dhcoloru8.GetColorU8FromFloat32Components(1.0, 0.0, 0.0, 1.0))
	DrawLine3D_SingleColor(
		dhvector.VGet(0.0, -length, 0.0),
		dhvector.Vector(0.0, 0.0, 0.0),
		dhcoloru8.GetColorU8FromFloat32Components(0.0, 1.0, 0.0, 1.0))
	DrawLine3D_SingleColor(
		dhvector.VGet(0.0, 0.0, -length),
		dhvector.Vector(0.0, 0.0, 0.0),
		dhcoloru8.GetColorU8FromFloat32Components(0.0, 0.0, 1.0, 1.0))
}
