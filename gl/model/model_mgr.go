package model

import (
	"unsafe"

	"github.com/dabasan/goglf/gl/texture"

	"github.com/dabasan/goglf/gl/shader"
	"github.com/dabasan/goglf/gl/wrapper"
	"github.com/go-gl/gl/all-core/gl"
)

type ModelMgr struct {
	buffered_vertices_list []*BufferedVertices

	property_updated_flag bool

	indices_vbo []uint32
	pos_vbo     []uint32
	uv_vbo      []uint32
	norm_vbo    []uint32
	vao         []uint32

	programs []*shader.ShaderProgram
}

func NewModelMgr(buffered_vertices_list []*BufferedVertices) *ModelMgr {
	model := new(ModelMgr)

	model.buffered_vertices_list = buffered_vertices_list

	model.property_updated_flag = false

	model.programs = make([]*shader.ShaderProgram, 1)
	program, _ := shader.NewShaderProgram("texture")
	model.programs[0] = program

	model.generateBuffers()

	return model
}

func (m *ModelMgr) AddProgram(program *shader.ShaderProgram) {
	m.programs = append(m.programs, program)
}
func (m *ModelMgr) RemoveAllPrograms() {
	m.programs = make([]*shader.ShaderProgram, 0)
}

func (m *ModelMgr) Interpolate(frame1 *ModelMgr, frame2 *ModelMgr, blend_ratio float32) {
	frame1_bv_list := frame1.buffered_vertices_list
	frame2_bv_list := frame2.buffered_vertices_list
	element_num := len(frame1_bv_list)

	interpolated_bv_list := make([]*BufferedVertices, element_num)

	for i := 0; i < element_num; i++ {
		frame1_bv := frame1_bv_list[i]
		frame2_bv := frame2_bv_list[i]

		interpolated_bv := InterpolateBufferedVertices(frame1_bv, frame2_bv, blend_ratio)
		interpolated_bv_list[i] = interpolated_bv
	}

	m.buffered_vertices_list = interpolated_bv_list
}

func (m *ModelMgr) Copy() *ModelMgr {
	copied_buffered_vertices_list := make([]*BufferedVertices, 0)

	for _, buffered_vertices := range m.buffered_vertices_list {
		copied_buffered_vertices := buffered_vertices.Copy()
		copied_buffered_vertices_list = append(copied_buffered_vertices_list, copied_buffered_vertices)
	}

	copied_model := NewModelMgr(copied_buffered_vertices_list)

	return copied_model
}

func (m *ModelMgr) generateBuffers() {
	element_num := len(m.buffered_vertices_list)
	indices_vbo := make([]uint32, element_num)
	pos_vbo := make([]uint32, element_num)
	uv_vbo := make([]uint32, element_num)
	norm_vbo := make([]uint32, element_num)
	vao := make([]uint32, element_num)

	m.indices_vbo = indices_vbo
	m.pos_vbo = pos_vbo
	m.uv_vbo = uv_vbo
	m.norm_vbo = norm_vbo
	m.vao = vao

	element_num_32 := int32(element_num)
	wrapper.GenBuffers(element_num_32, &indices_vbo[0])
	wrapper.GenBuffers(element_num_32, &pos_vbo[0])
	wrapper.GenBuffers(element_num_32, &uv_vbo[0])
	wrapper.GenBuffers(element_num_32, &norm_vbo[0])
	wrapper.GenVertexArrays(element_num_32, &vao[0])

	for i := 0; i < element_num; i++ {
		buffered_vertices := m.buffered_vertices_list[i]

		pos_buffer := buffered_vertices.GetPosBuffer()
		uv_buffer := buffered_vertices.GetUVBuffer()
		norm_buffer := buffered_vertices.GetNormBuffer()

		texture_handle := buffered_vertices.GetTextureHandle()

		//Flip UVs.
		uv_buffer_len := len(uv_buffer)

		for j := 0; j < uv_buffer_len; j += 2 {
			v := uv_buffer[j+1]
			v *= (-1.0)
			uv_buffer[j+1] = v
		}

		wrapper.BindBuffer(gl.ARRAY_BUFFER, pos_vbo[i])
		wrapper.BufferData(gl.ARRAY_BUFFER,
			wrapper.SIZEOF_FLOAT*len(pos_buffer), unsafe.Pointer(&pos_buffer[0]), gl.DYNAMIC_DRAW)
		wrapper.BindBuffer(gl.ARRAY_BUFFER, uv_vbo[i])
		wrapper.BufferData(gl.ARRAY_BUFFER,
			wrapper.SIZEOF_FLOAT*len(uv_buffer), unsafe.Pointer(&uv_buffer[0]), gl.STATIC_DRAW)
		wrapper.BindBuffer(gl.ARRAY_BUFFER, norm_vbo[i])
		wrapper.BufferData(gl.ARRAY_BUFFER,
			wrapper.SIZEOF_FLOAT*len(norm_buffer), unsafe.Pointer(&norm_buffer[0]), gl.DYNAMIC_DRAW)
		wrapper.BindBuffer(gl.ARRAY_BUFFER, 0)
	}
	for i := 0; i < element_num; i++ {
		buffered_vertices := m.buffered_vertices_list[i]
		indices_buffer := buffered_vertices.GetIndicesBuffer()

		wrapper.BindVertexArray(vao[i])

		wrapper.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indices_vbo[i])
		wrapper.BufferData(gl.ELEMENT_ARRAY_BUFFER,
			wrapper.SIZEOF_INT*len(indices_buffer), unsafe.Pointer(&indices_buffer[0]), gl.STATIC_DRAW)

		wrapper.BindBuffer(gl.ARRAY_BUFFER, pos_vbo[i])
		wrapper.EnableVertexAttribArray(0)
		wrapper.VertexAttribPointer(0, 3, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*3, nil)

		wrapper.BindBuffer(gl.ARRAY_BUFFER, uv_vbo[i])
		wrapper.EnableVertexAttribArray(1)
		wrapper.VertexAttribPointer(1, 2, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*2, nil)

		wrapper.BindBuffer(gl.ARRAY_BUFFER, norm_vbo[i])
		wrapper.EnableVertexAttribArray(2)
		wrapper.VertexAttribPointer(2, 3, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*3, nil)

		wrapper.BindBuffer(gl.ARRAY_BUFFER, 0)
		wrapper.BindVertexArray(0)
	}
}
func (m *ModelMgr) updateBuffers() {
	element_num := len(m.buffered_vertices_list)

	for i := 0; i < element_num; i++ {
		buffered_vertices := m.buffered_vertices_list[i]

		pos_buffer := buffered_vertices.GetPosBuffer()
		norm_buffer := buffered_vertices.GetNormBuffer()

		wrapper.BindVertexArray(m.vao[i])

		wrapper.BindBuffer(gl.ARRAY_BUFFER, m.pos_vbo[i])
		wrapper.BufferData(gl.ARRAY_BUFFER,
			wrapper.SIZEOF_FLOAT*len(pos_buffer), unsafe.Pointer(&pos_buffer[0]), gl.DYNAMIC_DRAW)
		wrapper.BindBuffer(gl.ARRAY_BUFFER, m.norm_vbo[i])
		wrapper.BufferData(gl.ARRAY_BUFFER,
			wrapper.SIZEOF_FLOAT*len(norm_buffer), unsafe.Pointer(&norm_buffer[0]), gl.DYNAMIC_DRAW)

		wrapper.BindBuffer(gl.ARRAY_BUFFER, 0)
		wrapper.BindVertexArray(0)
	}

	m.property_updated_flag = false
}

func (m *ModelMgr) DeleteBuffers() {
	element_num_32 := int32(len(m.buffered_vertices_list))
	indices_vbo := m.indices_vbo
	pos_vbo := m.pos_vbo
	uv_vbo := m.uv_vbo
	norm_vbo := m.norm_vbo
	vao := m.vao

	wrapper.DeleteBuffers(element_num_32, &indices_vbo[0])
	wrapper.DeleteBuffers(element_num_32, &pos_vbo[0])
	wrapper.DeleteBuffers(element_num_32, &uv_vbo[0])
	wrapper.DeleteBuffers(element_num_32, &norm_vbo[0])
	wrapper.DeleteBuffers(element_num_32, &vao[0])

	for _, buffered_vertices := range m.buffered_vertices_list {
		texture_handle := buffered_vertices.GetTextureHandle()
		texture.DeleteTexture(texture_handle)
	}
}

func (m *ModelMgr) DrawWithProgram(program *shader.ShaderProgram, texture_unit int, sampler_name string) {
	if m.property_updated_flag == true {
		m.updateBuffers()
	}

	element_num_32 := len(m.buffered_vertices_list)

	program.Enable()

	for i := 0; i < element_num_32; i++ {
		buffered_vertices := m.buffered_vertices_list[i]
		texture_handle := buffered_vertices.GetTextureHandle()
		indices_count := buffered_vertices.GetIndicesCount()

		wrapper.BindVertexArray(m.vao[i])

		program.SetTexture(sampler_name, texture_unit, texture_handle)

		wrapper.Enable(gl.BLEND)
		wrapper.DrawElements(gl.TRIANGLES, indices_count, gl.UNSIGNED_INT, nil)
		wrapper.Disable(gl.BLEND)

		wrapper.BindVertexArray(0)
	}

	program.Disable()
}
