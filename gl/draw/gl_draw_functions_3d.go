package draw

import (
	"unsafe"

	"github.com/go-gl/gl/all-core/gl"

	"github.com/dabasan/go-dh3dbasis/coloru8"
	"github.com/dabasan/go-dh3dbasis/vector"
	"github.com/dabasan/goglf/gl/shader"
	"github.com/dabasan/goglf/gl/wrapper"
)

func DrawLine3D(
	line_pos_1 vector.Vector, line_pos_2 vector.Vector,
	color_1 coloru8.ColorU8, color_2 coloru8.ColorU8) {
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

	shader.UseProgram("color")

	wrapper.GenBuffers(1, &pos_vbo)
	wrapper.GenBuffers(1, &color_vbo)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, pos_vbo)
	wrapper.BufferData(gl.ARRAY_BUFFER,
		wrapper.SIZEOF_FLOAT*len(pos_buffer), unsafe.Pointer(&pos_buffer[0]), gl.STATIC_DRAW)
	wrapper.BindBuffer(gl.ARRAY_BUFFER, color_vbo)
	wrapper.BufferData(gl.ARRAY_BUFFER,
		wrapper.SIZEOF_FLOAT*len(color_buffer), unsafe.Pointer(&color_buffer[0]), gl.STATIC_DRAW)

	wrapper.GenVertexArrays(1, &vao)
	wrapper.BindVertexArray(vao)

	//Position attribute
	wrapper.BindBuffer(gl.ARRAY_BUFFER, pos_vbo)
	wrapper.EnableVertexAttribArray(0)
	wrapper.VertexAttribPointer(0, 3, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*3, nil)

	//Color attribute
	wrapper.BindBuffer(gl.ARRAY_BUFFER, color_vbo)
	wrapper.EnableVertexAttribArray(1)
	wrapper.VertexAttribPointer(1, 4, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*4, nil)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, 0)
	wrapper.BindVertexArray(0)

	//Draw
	wrapper.BindVertexArray(vao)
	wrapper.Enable(gl.BLEND)
	wrapper.DrawArrays(gl.LINES, 0, 2)
	wrapper.Disable(gl.BLEND)
	wrapper.BindVertexArray(0)

	//Delete buffers
	wrapper.DeleteBuffers(1, &pos_vbo)
	wrapper.DeleteBuffers(1, &color_vbo)
	wrapper.DeleteVertexArrays(1, &vao)
}
func DrawLine3D_SingleColor(line_pos_1 vector.Vector, line_pos_2 vector.Vector, color coloru8.ColorU8) {
	DrawLine3D(line_pos_1, line_pos_2, color, color)
}

func DrawAxes(length float32) {
	DrawLine3D_SingleColor(
		vector.VGet(-length, 0.0, 0.0),
		vector.VGet(length, 0.0, 0.0),
		coloru8.GetColorU8FromFloat32Components(1.0, 0.0, 0.0, 1.0))
	DrawLine3D_SingleColor(
		vector.VGet(0.0, -length, 0.0),
		vector.VGet(0.0, length, 0.0),
		coloru8.GetColorU8FromFloat32Components(0.0, 1.0, 0.0, 1.0))
	DrawLine3D_SingleColor(
		vector.VGet(0.0, 0.0, -length),
		vector.VGet(0.0, 0.0, length),
		coloru8.GetColorU8FromFloat32Components(0.0, 0.0, 1.0, 1.0))
}
func DrawAxes_Positive(length float32) {
	DrawLine3D_SingleColor(
		vector.VGet(0.0, 0.0, 0.0),
		vector.VGet(length, 0.0, 0.0),
		coloru8.GetColorU8FromFloat32Components(1.0, 0.0, 0.0, 1.0))
	DrawLine3D_SingleColor(
		vector.VGet(0.0, 0.0, 0.0),
		vector.VGet(0.0, length, 0.0),
		coloru8.GetColorU8FromFloat32Components(0.0, 1.0, 0.0, 1.0))
	DrawLine3D_SingleColor(
		vector.VGet(0.0, 0.0, 0.0),
		vector.VGet(0.0, 0.0, length),
		coloru8.GetColorU8FromFloat32Components(0.0, 0.0, 1.0, 1.0))
}
func DrawAxes_Negative(length float32) {
	DrawLine3D_SingleColor(
		vector.VGet(-length, 0.0, 0.0),
		vector.VGet(0.0, 0.0, 0.0),
		coloru8.GetColorU8FromFloat32Components(1.0, 0.0, 0.0, 1.0))
	DrawLine3D_SingleColor(
		vector.VGet(0.0, -length, 0.0),
		vector.VGet(0.0, 0.0, 0.0),
		coloru8.GetColorU8FromFloat32Components(0.0, 1.0, 0.0, 1.0))
	DrawLine3D_SingleColor(
		vector.VGet(0.0, 0.0, -length),
		vector.VGet(0.0, 0.0, 0.0),
		coloru8.GetColorU8FromFloat32Components(0.0, 0.0, 1.0, 1.0))
}
