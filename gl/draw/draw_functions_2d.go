package draw

import (
	"math"
	"unsafe"

	"github.com/dabasan/go-dh3dbasis/coloru8"

	"github.com/dabasan/goglf/gl/common"
	"github.com/dabasan/goglf/gl/coordinatetool"
	"github.com/dabasan/goglf/gl/shader"
	"github.com/dabasan/goglf/gl/wrapper"

	"github.com/go-gl/gl/all-core/gl"
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
	simple_2d_program, _ = shader.NewShaderProgram("simple_2d")
	simple_2d_program.Enable()
	simple_2d_program.SetUniform1f("z", -1.0)
	simple_2d_program.Disable()

	texture_drawer_program, _ = shader.NewShaderProgram("texture_drawer")
	texture_drawer_program.Enable()
	texture_drawer_program.SetUniform1f("z", -1.0)
	texture_drawer_program.Disable()
}

func SetSimple2DZ(z float32) {
	simple_2d_program.Enable()
	simple_2d_program.SetUniform1f("z", z)
	simple_2d_program.Disable()
}
func SetTextureDrawerZ(z float32) {
	texture_drawer_program.Enable()
	texture_drawer_program.SetUniform1f("z", z)
	texture_drawer_program.Disable()
}

func SetWindowSize(width int, height int) {
	window_width = width
	window_height = height
}

func DrawLine2D(x1 int, y1 int, x2 int, y2 int, color coloru8.ColorU8) {
	var pos_vbo uint32
	var color_vbo uint32
	var vao uint32

	pos_buffer := make([]float32, 4)
	color_buffer := make([]float32, 8)

	normalized_x1 := coordinatetool.NormalizeCoordinate_Int(x1, window_width)
	normalized_y1 := coordinatetool.NormalizeCoordinate_Int(y1, window_height)
	normalized_x2 := coordinatetool.NormalizeCoordinate_Int(x2, window_width)
	normalized_y2 := coordinatetool.NormalizeCoordinate_Int(y2, window_height)

	pos_buffer[0] = normalized_x1
	pos_buffer[1] = normalized_y1
	pos_buffer[2] = normalized_x2
	pos_buffer[3] = normalized_y2

	color_buffer[0] = color.R
	color_buffer[1] = color.G
	color_buffer[2] = color.B
	color_buffer[3] = color.A
	color_buffer[4] = color.R
	color_buffer[5] = color.G
	color_buffer[6] = color.B
	color_buffer[7] = color.A

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
	wrapper.VertexAttribPointer(0, 2, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*2, nil)

	//Color attribute
	wrapper.BindBuffer(gl.ARRAY_BUFFER, color_vbo)
	wrapper.EnableVertexAttribArray(1)
	wrapper.VertexAttribPointer(1, 4, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*4, nil)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, 0)
	wrapper.BindVertexArray(0)

	//Draw
	wrapper.BindVertexArray(vao)
	wrapper.Enable(gl.BLEND)
	simple_2d_program.Enable()
	wrapper.DrawArrays(gl.LINES, 0, 2)
	simple_2d_program.Disable()
	wrapper.Disable(gl.BLEND)
	wrapper.BindVertexArray(0)

	//Delete buffers
	wrapper.DeleteBuffers(1, &pos_vbo)
	wrapper.DeleteBuffers(1, &color_vbo)
	wrapper.DeleteVertexArrays(1, &vao)
}

func DrawRectangle2D(x1 int, y1 int, x2 int, y2 int, color coloru8.ColorU8) {
	var pos_vbo uint32
	var color_vbo uint32
	var vao uint32

	pos_buffer := make([]float32, 2*4)
	color_buffer := make([]float32, 4*4)

	normalized_x1 := coordinatetool.NormalizeCoordinate_Int(x1, window_width)
	normalized_y1 := coordinatetool.NormalizeCoordinate_Int(y1, window_height)
	normalized_x2 := coordinatetool.NormalizeCoordinate_Int(x2, window_width)
	normalized_y2 := coordinatetool.NormalizeCoordinate_Int(y2, window_height)

	pos_buffer[0] = normalized_x1
	pos_buffer[1] = normalized_y1
	pos_buffer[2] = normalized_x2
	pos_buffer[3] = normalized_y1
	pos_buffer[4] = normalized_x2
	pos_buffer[5] = normalized_y2
	pos_buffer[6] = normalized_x1
	pos_buffer[7] = normalized_y2

	for i := 0; i < 4; i++ {
		color_buffer[i*4] = color.R
		color_buffer[i*4+1] = color.G
		color_buffer[i*4+2] = color.B
		color_buffer[i*4+3] = color.A
	}

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
	wrapper.VertexAttribPointer(0, 2, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*2, nil)

	//Color attribute
	wrapper.BindBuffer(gl.ARRAY_BUFFER, color_vbo)
	wrapper.EnableVertexAttribArray(1)
	wrapper.VertexAttribPointer(1, 4, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*4, nil)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, 0)
	wrapper.BindVertexArray(0)

	//Draw
	wrapper.BindVertexArray(vao)
	wrapper.Enable(gl.BLEND)
	simple_2d_program.Enable()
	wrapper.DrawArrays(gl.LINE_LOOP, 0, 4)
	simple_2d_program.Disable()
	wrapper.Disable(gl.BLEND)
	wrapper.BindVertexArray(0)

	//Delete buffers
	wrapper.DeleteBuffers(1, &pos_vbo)
	wrapper.DeleteBuffers(1, &color_vbo)
	wrapper.DeleteVertexArrays(1, &vao)
}

func DrawFilledRectangle2D(x1 int, y1 int, x2 int, y2 int, color coloru8.ColorU8) {
	var indices_vbo uint32
	var pos_vbo uint32
	var color_vbo uint32
	var vao uint32

	pos_buffer := make([]float32, 2*4)
	color_buffer := make([]float32, 4*4)

	normalized_x1 := coordinatetool.NormalizeCoordinate_Int(x1, window_width)
	normalized_y1 := coordinatetool.NormalizeCoordinate_Int(y1, window_height)
	normalized_x2 := coordinatetool.NormalizeCoordinate_Int(x2, window_width)
	normalized_y2 := coordinatetool.NormalizeCoordinate_Int(y2, window_height)

	//Bottom left
	pos_buffer[0] = normalized_x1
	pos_buffer[1] = normalized_y1
	//Bottom right
	pos_buffer[2] = normalized_x2
	pos_buffer[3] = normalized_y1
	//Top right
	pos_buffer[4] = normalized_x2
	pos_buffer[5] = normalized_y2
	//Top left
	pos_buffer[6] = normalized_x1
	pos_buffer[7] = normalized_y2

	for i := 0; i < 4; i++ {
		color_buffer[i*4] = color.R
		color_buffer[i*4+1] = color.G
		color_buffer[i*4+2] = color.B
		color_buffer[i*4+3] = color.A
	}

	indices_buffer := make([]uint32, 3*2)
	indices_buffer[0] = 0
	indices_buffer[1] = 1
	indices_buffer[2] = 2
	indices_buffer[3] = 2
	indices_buffer[4] = 3
	indices_buffer[5] = 0

	wrapper.GenBuffers(1, &indices_vbo)
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

	//Indices
	wrapper.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indices_vbo)
	wrapper.BufferData(gl.ELEMENT_ARRAY_BUFFER,
		wrapper.SIZEOF_INT*len(indices_buffer), unsafe.Pointer(&indices_buffer[0]), gl.STATIC_DRAW)

	//Position attribute
	wrapper.BindBuffer(gl.ARRAY_BUFFER, pos_vbo)
	wrapper.EnableVertexAttribArray(0)
	wrapper.VertexAttribPointer(0, 2, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*2, nil)

	//Color attribute
	wrapper.BindBuffer(gl.ARRAY_BUFFER, color_vbo)
	wrapper.EnableVertexAttribArray(1)
	wrapper.VertexAttribPointer(1, 4, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*4, nil)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, 0)
	wrapper.BindVertexArray(0)

	//Draw
	wrapper.BindVertexArray(vao)
	wrapper.Enable(gl.BLEND)
	simple_2d_program.Enable()
	wrapper.DrawElements(gl.TRIANGLES, int32(len(indices_buffer)), gl.UNSIGNED_INT, nil)
	simple_2d_program.Disable()
	wrapper.Disable(gl.BLEND)
	wrapper.BindVertexArray(0)

	//Delete buffers
	wrapper.DeleteBuffers(1, &indices_vbo)
	wrapper.DeleteBuffers(1, &pos_vbo)
	wrapper.DeleteBuffers(1, &color_vbo)
	wrapper.DeleteVertexArrays(1, &vao)
}

func DrawCircle2D(center_x int, center_y int, radius int, div_num int, color coloru8.ColorU8) {
	var pos_vbo uint32
	var color_vbo uint32
	var vao uint32

	pos_buffer := make([]float32, 2*div_num)
	color_buffer := make([]float32, 4*div_num)

	count := 0
	for i := 0; i < div_num; i++ {
		th := math.Pi * 2.0 / float64(div_num) * float64(i)

		x := float64(radius)*math.Cos(th) + float64(center_x)
		y := float64(radius)*math.Sin(th) + float64(center_y)

		normalized_x := coordinatetool.NormalizeCoordinate_Float32(float32(x), float32(window_width))
		normalized_y := coordinatetool.NormalizeCoordinate_Float32(float32(y), float32(window_height))

		pos_buffer[count] = normalized_x
		pos_buffer[count+1] = normalized_y
		count += 2
	}

	for i := 0; i < div_num; i++ {
		color_buffer[i*4] = color.R
		color_buffer[i*4+1] = color.G
		color_buffer[i*4+2] = color.B
		color_buffer[i*4+3] = color.A
	}

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
	wrapper.VertexAttribPointer(0, 2, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*2, nil)

	//Color attribute
	wrapper.BindBuffer(gl.ARRAY_BUFFER, color_vbo)
	wrapper.EnableVertexAttribArray(1)
	wrapper.VertexAttribPointer(1, 4, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*4, nil)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, 0)
	wrapper.BindVertexArray(0)

	//Draw
	wrapper.BindVertexArray(vao)
	wrapper.Enable(gl.BLEND)
	simple_2d_program.Enable()
	wrapper.DrawArrays(gl.LINE_LOOP, 0, int32(div_num))
	simple_2d_program.Disable()
	wrapper.Disable(gl.BLEND)
	wrapper.BindVertexArray(0)

	//Delete buffers
	wrapper.DeleteBuffers(1, &pos_vbo)
	wrapper.DeleteBuffers(1, &color_vbo)
	wrapper.DeleteVertexArrays(1, &vao)
}

func DrawTexture(
	texture_handle int, x int, y int, width int, height int,
	bottom_left_u float32, bottom_left_v float32,
	bottom_right_u float32, bottom_right_v float32,
	top_right_u float32, top_right_v float32,
	top_left_u float32, top_left_v float32) int {
	indices_buffer := make([]uint32, 6)
	pos_buffer := make([]float32, 8)
	uv_buffer := make([]float32, 8)

	indices_buffer[0] = 0
	indices_buffer[1] = 1
	indices_buffer[2] = 2
	indices_buffer[3] = 2
	indices_buffer[4] = 3
	indices_buffer[5] = 0

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
		wrapper.SIZEOF_INT*len(indices_buffer), unsafe.Pointer(&indices_buffer[0]), gl.STATIC_DRAW)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, pos_vbo)
	wrapper.EnableVertexAttribArray(0)
	wrapper.VertexAttribPointer(0, 2, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*2, nil)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, uv_vbo)
	wrapper.EnableVertexAttribArray(1)
	wrapper.VertexAttribPointer(1, 2, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*2, nil)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, 0)

	wrapper.Enable(gl.BLEND)
	texture_drawer_program.Enable()
	texture_drawer_program.SetTexture("texture_sampler", 0, texture_handle)
	wrapper.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
	texture_drawer_program.Disable()
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
