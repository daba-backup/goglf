package transferrer

import (
	"unsafe"

	"github.com/dabasan/goglf/gl/wrapper"
	"github.com/go-gl/gl/all-core/gl"
)

type FullscreenQuadTransferrerWithUV struct {
	fullscreenQuadTransferrerBase
	uv_vbo uint32
}

func NewFullscreenQuadTransferrerWithUV(flip_v_flag bool) *FullscreenQuadTransferrerWithUV {
	t := new(FullscreenQuadTransferrerWithUV)

	wrapper.GenBuffers(1, &t.indices_vbo)
	wrapper.GenBuffers(1, &t.pos_vbo)
	wrapper.GenBuffers(1, &t.uv_vbo)
	wrapper.GenVertexArrays(1, &t.vao)

	indices_buffer := make([]uint32, 6)
	pos_buffer := make([]float32, 2*4)
	uv_buffer := make([]float32, 2*4)

	//First triangle
	indices_buffer[0] = 0
	indices_buffer[1] = 1
	indices_buffer[2] = 2
	//Second triangle
	indices_buffer[3] = 2
	indices_buffer[4] = 3
	indices_buffer[5] = 0

	//Bottom left
	pos_buffer[0] = -1.0
	pos_buffer[1] = -1.0
	//Bottom right
	pos_buffer[2] = 1.0
	pos_buffer[3] = -1.0
	//Top right
	pos_buffer[4] = 1.0
	pos_buffer[5] = 1.0
	//Top left
	pos_buffer[6] = -1.0
	pos_buffer[7] = 1.0

	if flip_v_flag == true {
		//Bottom left
		uv_buffer[0] = 0.0
		uv_buffer[1] = 0.0
		//Bottom right
		uv_buffer[2] = 1.0
		uv_buffer[3] = 0.0
		//Top right
		uv_buffer[4] = 1.0
		uv_buffer[5] = 1.0
		//Top left
		uv_buffer[6] = 0.0
		uv_buffer[7] = 1.0
	} else {
		//Bottom left
		uv_buffer[0] = 0.0
		uv_buffer[1] = 1.0
		//Bottom right
		uv_buffer[2] = 1.0
		uv_buffer[3] = 1.0
		//Top right
		uv_buffer[4] = 1.0
		uv_buffer[5] = 0.0
		//Top left
		uv_buffer[6] = 0.0
		uv_buffer[7] = 0.0
	}

	wrapper.BindBuffer(gl.ARRAY_BUFFER, t.pos_vbo)
	wrapper.BufferData(gl.ARRAY_BUFFER,
		wrapper.SIZEOF_FLOAT*len(pos_buffer), unsafe.Pointer(&pos_buffer[0]), gl.STATIC_DRAW)
	wrapper.BindBuffer(gl.ARRAY_BUFFER, t.uv_vbo)
	wrapper.BufferData(gl.ARRAY_BUFFER,
		wrapper.SIZEOF_FLOAT*len(uv_buffer), unsafe.Pointer(&uv_buffer[0]), gl.STATIC_DRAW)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, 0)

	wrapper.BindVertexArray(t.vao)

	wrapper.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, t.indices_vbo)
	wrapper.BufferData(gl.ELEMENT_ARRAY_BUFFER,
		wrapper.SIZEOF_INT*len(indices_buffer), unsafe.Pointer(&indices_buffer[0]), gl.STATIC_DRAW)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, t.pos_vbo)
	wrapper.EnableVertexAttribArray(0)
	wrapper.VertexAttribPointer(0, 2, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*2, nil)
	wrapper.BindBuffer(gl.ARRAY_BUFFER, t.uv_vbo)
	wrapper.EnableVertexAttribArray(1)
	wrapper.VertexAttribPointer(1, 2, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*2, nil)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, 0)
	wrapper.BindVertexArray(0)

	return t
}

func (t *FullscreenQuadTransferrerWithUV) DeleteBuffers() {
	wrapper.DeleteBuffers(1, &t.indices_vbo)
	wrapper.DeleteBuffers(1, &t.pos_vbo)
	wrapper.DeleteBuffers(1, &t.uv_vbo)
	wrapper.DeleteVertexArrays(1, &t.vao)
}

func (t *FullscreenQuadTransferrerWithUV) Transfer() {
	wrapper.BindVertexArray(t.vao)
	wrapper.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
	wrapper.BindVertexArray(0)
}
