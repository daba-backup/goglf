package model

import (
	"log"
	"unsafe"

	"github.com/dabasan/go-dh3dbasis/matrix"
	"github.com/dabasan/go-dh3dbasis/vector"

	"github.com/dabasan/goglf/gl/model/buffer"
	"github.com/dabasan/goglf/gl/shader"
	"github.com/dabasan/goglf/gl/shape"
	"github.com/dabasan/goglf/gl/texture"
	"github.com/dabasan/goglf/gl/wrapper"

	"github.com/go-gl/gl/all-core/gl"
)

type ModelMgr struct {
	buffered_vertices_list []*buffer.BufferedVertices

	property_updated_flag bool

	indices_vbo []uint32
	pos_vbo     []uint32
	uv_vbo      []uint32
	norm_vbo    []uint32
	vao         []uint32

	programs []*shader.ShaderProgram
}

func NewModelMgr(buffered_vertices_list []*buffer.BufferedVertices) *ModelMgr {
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

	interpolated_bv_list := make([]*buffer.BufferedVertices, element_num)

	for i := 0; i < element_num; i++ {
		frame1_bv := frame1_bv_list[i]
		frame2_bv := frame2_bv_list[i]

		interpolated_bv := buffer.Interpolate(frame1_bv, frame2_bv, blend_ratio)
		interpolated_bv_list[i] = interpolated_bv
	}

	m.buffered_vertices_list = interpolated_bv_list
}

func (m *ModelMgr) Copy() *ModelMgr {
	copied_buffered_vertices_list := make([]*buffer.BufferedVertices, 0)

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
	for i, buffered_vertices := range m.buffered_vertices_list {
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

	wrapper.DeleteBuffers(element_num_32, &m.indices_vbo[0])
	wrapper.DeleteBuffers(element_num_32, &m.pos_vbo[0])
	wrapper.DeleteBuffers(element_num_32, &m.uv_vbo[0])
	wrapper.DeleteBuffers(element_num_32, &m.norm_vbo[0])
	wrapper.DeleteBuffers(element_num_32, &m.vao[0])

	for _, buffered_vertices := range m.buffered_vertices_list {
		texture_handle := buffered_vertices.GetTextureHandle()
		texture.DeleteTexture(texture_handle)
	}
}

func (m *ModelMgr) DrawWithProgram(program *shader.ShaderProgram, sampler_name string, texture_unit int) {
	if m.property_updated_flag == true {
		m.updateBuffers()
	}

	program.Enable()

	for i, buffered_vertices := range m.buffered_vertices_list {
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
func (m *ModelMgr) Draw(sampler_name string, texture_unit int) {
	for _, program := range m.programs {
		m.DrawWithProgram(program, sampler_name, texture_unit)
	}
}
func (m *ModelMgr) Draw_Simple() {
	m.Draw("texture_sampler", 0)
}

func (m *ModelMgr) Transfer() {
	if m.property_updated_flag == false {
		m.updateBuffers()
	}

	for i, buffered_vertices := range m.buffered_vertices_list {
		indices_count := buffered_vertices.GetIndicesCount()

		wrapper.BindVertexArray(m.vao[i])

		wrapper.Enable(gl.BLEND)
		wrapper.DrawElements(gl.TRIANGLES, indices_count, gl.UNSIGNED_INT, nil)
		wrapper.Disable(gl.BLEND)

		wrapper.BindVertexArray(0)
	}
}

func (m *ModelMgr) DrawElements(sampler_name string, texture_unit int, bound int) {
	if m.property_updated_flag == false {
		m.updateBuffers()
	}

	element_num := len(m.buffered_vertices_list)

	var clamped_bound int
	if bound < 0 {
		clamped_bound = 0
	} else if bound < element_num {
		clamped_bound = bound
	} else {
		clamped_bound = element_num
	}

	for _, program := range m.programs {
		program.Enable()

		for i := 0; i < clamped_bound; i++ {
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
}
func (m *ModelMgr) DrawElements_Simple(bound int) {
	m.DrawElements("texture_sampler", 0, bound)
}

func (m *ModelMgr) GetElementNum() int {
	return len(m.buffered_vertices_list)
}

func (m *ModelMgr) SetMatrix(mat matrix.Matrix) {
	for _, buffered_vertices := range m.buffered_vertices_list {
		pos_buffer := buffered_vertices.GetPosBuffer()
		norm_buffer := buffered_vertices.GetNormBuffer()

		length := len(pos_buffer)
		for i := 0; i < length; i += 3 {
			//pos buffer
			var pos vector.Vector
			pos.X = pos_buffer[i]
			pos.Y = pos_buffer[i+1]
			pos.Z = pos_buffer[i+2]

			pos = matrix.VTransform(pos, mat)

			pos_buffer[i] = pos.X
			pos_buffer[i+1] = pos.Y
			pos_buffer[i+2] = pos.Z

			//norm buffer
			var norm vector.Vector
			norm.X = norm_buffer[i]
			norm.Y = norm_buffer[i+1]
			norm.Z = norm_buffer[i+2]

			norm = matrix.VTransformSR(norm, mat)
			norm = vector.VNorm(norm)

			norm_buffer[i] = norm.X
			norm_buffer[i+1] = norm.Y
			norm_buffer[i+2] = norm.Z
		}

		buffered_vertices.SetPosBuffer(pos_buffer)
		buffered_vertices.SetNormBuffer(norm_buffer)
	}

	m.property_updated_flag = true
}

func (m *ModelMgr) ChangeTexture(material_index int, new_texture_handle int) int {
	if !(0 <= material_index && material_index < len(m.buffered_vertices_list)) {
		log.Printf("warn: Index out of bounds. material_index=%v", material_index)
		return -1
	}

	buffered_vertices := m.buffered_vertices_list[material_index]
	buffered_vertices.SetTextureHandle(new_texture_handle)

	return 0
}

func (m *ModelMgr) GetFaces() []*shape.Triangle {
	total_triangle_num := 0
	for _, buffered_vertices := range m.buffered_vertices_list {
		indices_count := buffered_vertices.GetIndicesCount()
		triangle_num := indices_count / 9

		total_triangle_num += int(triangle_num)
	}

	ret := make([]*shape.Triangle, total_triangle_num)
	count := 0

	for _, buffered_vertices := range m.buffered_vertices_list {
		indices_buffer := buffered_vertices.GetIndicesBuffer()
		pos_buffer := buffered_vertices.GetPosBuffer()
		norm_buffer := buffered_vertices.GetNormBuffer()
		uv_buffer := buffered_vertices.GetUVBuffer()

		length := len(pos_buffer)
		triangle_num := length / 9

		for i := 0; i < triangle_num; i++ {
			var triangle shape.Triangle

			for j := 0; j < 3; j++ {
				index := indices_buffer[i*3+j]

				vec_base_index := index * 3
				uv_base_index := index * 2

				//pos buffer
				var pos vector.Vector
				pos.X = pos_buffer[vec_base_index]
				pos.Y = pos_buffer[vec_base_index+1]
				pos.Z = pos_buffer[vec_base_index+2]

				//norm buffer
				var norm vector.Vector
				norm.X = norm_buffer[vec_base_index]
				norm.Y = norm_buffer[vec_base_index+1]
				norm.Z = norm_buffer[vec_base_index+2]

				//uv buffer
				u := uv_buffer[uv_base_index]
				v := uv_buffer[uv_base_index+1]

				triangle.Vertices[j].Pos = pos
				triangle.Vertices[j].Norm = norm
				triangle.Vertices[j].U = u
				triangle.Vertices[j].V = v
			}

			ret[count] = &triangle
			count++
		}
	}

	return ret
}
