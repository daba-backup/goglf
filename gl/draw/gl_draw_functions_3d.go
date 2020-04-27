package draw

import (
	"unsafe"

	"github.com/go-gl/gl/all-core/gl"

	"github.com/dabasan/go-dh3dbasis/coloru8"
	"github.com/dabasan/go-dh3dbasis/vector"
	"github.com/dabasan/goglf/gl/shader"
	"github.com/dabasan/goglf/gl/shape"
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

func DrawTriangle3D(triangle *shape.Triangle) {
	var pos_vbo uint32
	var color_vbo uint32
	var vao uint32

	pos_buffer := make([]float32, 9)
	color_buffer := make([]float32, 12)
	for i := 0; i < 3; i++ {
		pos := triangle.Vertices[i].Pos
		dif := triangle.Vertices[i].Dif

		pos_buffer[i*3] = pos.X
		pos_buffer[i*3+1] = pos.Y
		pos_buffer[i*3+2] = pos.Z
		color_buffer[i*4] = dif.R
		color_buffer[i*4+1] = dif.G
		color_buffer[i*4+2] = dif.B
		color_buffer[i*4+3] = dif.A
	}

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
	wrapper.DrawArrays(gl.LINE_LOOP, 0, 3)
	wrapper.Disable(gl.BLEND)
	wrapper.BindVertexArray(0)

	//Delete buffers
	wrapper.DeleteBuffers(1, &pos_vbo)
	wrapper.DeleteBuffers(1, &color_vbo)
	wrapper.DeleteVertexArrays(1, &vao)
}
func DrawTriangle3D_Vertices(pos_1 vector.Vector, pos_2 vector.Vector, pos_3 vector.Vector, color coloru8.ColorU8) {
	var triangle shape.Triangle

	triangle.Vertices[0].Pos = pos_1
	triangle.Vertices[1].Pos = pos_2
	triangle.Vertices[2].Pos = pos_3
	for i := 0; i < 3; i++ {
		triangle.Vertices[i].Dif = color
	}

	DrawTriangle3D(&triangle)
}
func DrawQuadrangle3D(quadrangle *shape.Quadrangle) {
	var pos_vbo uint32
	var color_vbo uint32
	var vao uint32

	pos_buffer := make([]float32, 12)
	color_buffer := make([]float32, 16)
	for i := 0; i < 4; i++ {
		pos := quadrangle.Vertices[i].Pos
		dif := quadrangle.Vertices[i].Dif

		pos_buffer[i*3] = pos.X
		pos_buffer[i*3+1] = pos.Y
		pos_buffer[i*3+2] = pos.Z
		color_buffer[i*4] = dif.R
		color_buffer[i*4+1] = dif.G
		color_buffer[i*4+2] = dif.B
		color_buffer[i*4+3] = dif.A
	}

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
	wrapper.DrawArrays(gl.LINE_LOOP, 0, 4)
	wrapper.Disable(gl.BLEND)
	wrapper.BindVertexArray(0)

	//Delete buffers
	wrapper.DeleteBuffers(1, &pos_vbo)
	wrapper.DeleteBuffers(1, &color_vbo)
	wrapper.DeleteVertexArrays(1, &vao)
}
func DrawQuadrangle3D_Vertices(
	pos_1 vector.Vector, pos_2 vector.Vector,
	pos_3 vector.Vector, pos_4 vector.Vector, color coloru8.ColorU8) {
	var quadrangle shape.Quadrangle

	quadrangle.Vertices[0].Pos = pos_1
	quadrangle.Vertices[1].Pos = pos_2
	quadrangle.Vertices[2].Pos = pos_3
	quadrangle.Vertices[3].Pos = pos_4

	for i := 0; i < 4; i++ {
		quadrangle.Vertices[i].Dif = color
	}

	DrawQuadrangle3D(&quadrangle)
}
