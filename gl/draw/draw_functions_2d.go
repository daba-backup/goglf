package draw

import (
	"unsafe"

	"github.com/dabasan/goglf/gl/common"

	"github.com/go-gl/gl/all-core/gl"

	"github.com/dabasan/goglf/gl/shader"
	"github.com/dabasan/goglf/gl/wrapper"
)

var simple_2d_program *shader.ShaderProgram
var texture_drawer_program *shader.ShaderProgram

var window_width int
var window_height int

func init() {
	window_width = common.WINDOW_DEFAULT_WIDTH
	window_height = common.WINDOW_DEFAULT_HEIGHT
}

func InitializeDrawFunctions2D() {
	simple_2d_program, _ := shader.NewShaderProgram("simple_2d")
	texture_drawer_program, _ := shader.NewShaderProgram("texture_drawer")
}

func SetWindowSize(width int, height int) {
	window_width = width
	window_height = height
}

func DrawTexture(
	texture_handle int, x int, y int, width int, height int,
	bottom_left_u float32, bottom_left_v float32,
	bottom_right_u float32, bottom_right_v float32,
	top_right_u float32, top_right_v float32,
	top_left_u float32, top_left_v float32) int {
	indices := make([]int32, 6)
	pos_buffer := make([]float32, 8)
	uv_buffer := make([]float32, 8)

	indices[0] = 0
	indices[1] = 1
	indices[2] = 2
	indices[3] = 2
	indices[4] = 3
	indices[5] = 0

	normalized_x := 2.0*float32(x)/float32(window_width) - 1.0
	normalized_y := 2.0*float32(y)/float32(window_height) - 1.0
	normalized_width := float32(width) / float32(window_width) * 2.0
	normalized_height := float32(height) / float32(window_height) * 2.0

	//Bottom left
	pos_buffer[0] = normalized_x
	pos_buffer[1] = normalized_y
	uv_buffer[0] = bottom_left_u
	uv_buffer[1] = bottom_left_v
	//Bottom right
	pos_buffer[2] = normalized_x + normalized_width
	pos_buffer[3] = normalized_y
	uv_buffer[2] = bottom_right_u
	uv_buffer[3] = bottom_right_v
	//Top right
	pos_buffer[4] = normalized_x + normalized_width
	pos_buffer[5] = normalized_y + normalized_height
	uv_buffer[4] = top_right_u
	uv_buffer[5] = top_right_v
	//Top left
	pos_buffer[6] = normalized_x
	pos_buffer[7] = normalized_y + normalized_height
	uv_buffer[6] = top_left_u
	uv_buffer[7] = top_left_v

	var indices_vbo uint32
	var pos_vbo uint32
	var uv_vbo uint32
	var vao uint32

	wrapper.GenBuffers(1, &indices_vbo)
	wrapper.GenBuffers(1, &pos_vbo)
	wrapper.GenBuffers(1, &uv_vbo)
	wrapper.GenVertexArrays(1, &vao)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, pos_vbo)
	wrapper.BufferData(gl.ARRAY_BUFFER,
		wrapper.SIZEOF_FLOAT*len(pos_buffer), unsafe.Pointer(&pos_buffer[0]), gl.STATIC_DRAW)
	wrapper.BindBuffer(gl.ARRAY_BUFFER, uv_vbo)
	wrapper.BufferData(gl.ARRAY_BUFFER,
		wrapper.SIZEOF_FLOAT*len(uv_buffer), unsafe.Pointer(&uv_buffer[0]), gl.STATIC_DRAW)

	wrapper.BindVertexArray(vao)

	wrapper.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indices_vbo)
	wrapper.BufferData(gl.ELEMENT_ARRAY_BUFFER,
		wrapper.SIZEOF_INT*len(indices), unsafe.Pointer(&indices[0]), gl.STATIC_DRAW)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, pos_vbo)
	wrapper.EnableVertexAttribArray(0)
	wrapper.VertexAttribPointer(0, 2, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*2, nil)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, uv_vbo)
	wrapper.EnableVertexAttribArray(1)
	wrapper.VertexAttribPointer(1, 2, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*2, nil)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, 0)

	wrapper.Enable(gl.BLEND)
	texture_program.Enable()
	texture_program.SetTexture("texture_sampler", 0, texture_handle)
	wrapper.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
	texture_program.Disable()
	wrapper.Disable(gl.BLEND)

	wrapper.BindVertexArray(0)

	wrapper.DeleteBuffers(1, &indices_vbo)
	wrapper.DeleteBuffers(1, &pos_vbo)
	wrapper.DeleteBuffers(1, &uv_vbo)
	wrapper.DeleteVertexArrays(1, &vao)

	return 0
}
func DrawTexture_Simple(texture_handle int, x int, y int, width int, height int) int {
	ret := DrawTexture(texture_handle, x, y, width, height, 0.0, 0.0, 1.0, 0.0, 1.0, 1.0, 0.0, 1.0)
	return ret
}
