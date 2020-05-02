package model

import (
	"github.com/dabasan/go-dh3dbasis/coloru8"
)

type BufferedVertices struct {
	texture_handle int
	indices_count  int32

	indices_buffer []uint32
	pos_buffer     []float32
	uv_buffer      []float32
	norm_buffer    []float32

	ambient_color       coloru8.ColorU8
	diffuse_color       coloru8.ColorU8
	specular_color      coloru8.ColorU8
	specular_exponent   float32
	dissolve            float32
	diffuse_texture_map string
}

func NewBufferedVertices() *BufferedVertices {
	bv := new(BufferedVertices)

	bv.texture_handle = -1
	bv.indices_count = 0

	bv.ambient_color = coloru8.GetColorU8FromFloat32Components(1.0, 1.0, 1.0, 1.0)
	bv.diffuse_color = coloru8.GetColorU8FromFloat32Components(1.0, 1.0, 1.0, 1.0)
	bv.specular_color = coloru8.GetColorU8FromFloat32Components(0.0, 0.0, 0.0, 0.0)
	bv.specular_exponent = 10.0
	bv.dissolve = 1.0
	bv.diffuse_texture_map = ""

	return bv
}

func (bv *BufferedVertices) Copy() *BufferedVertices {
	copied := NewBufferedVertices()

	copied.texture_handle = bv.texture_handle
	copied.indices_count = bv.indices_count

	copied.indices_buffer = copyUint32Slice(bv.indices_buffer)
	copied.pos_buffer = copyFloat32Slice(bv.pos_buffer)
	copied.uv_buffer = copyFloat32Slice(bv.uv_buffer)
	copied.norm_buffer = copyFloat32Slice(bv.norm_buffer)

	copied.ambient_color = bv.ambient_color
	copied.diffuse_color = bv.diffuse_color
	copied.specular_color = bv.specular_color
	copied.specular_exponent = bv.specular_exponent
	copied.dissolve = bv.dissolve
	copied.diffuse_texture_map = bv.diffuse_texture_map

	return copied
}
func copyUint32Slice(s []uint32) []uint32 {
	copied := make([]uint32, len(s))
	copy(copied, s)

	return copied
}
func copyFloat32Slice(s []float32) []float32 {
	copied := make([]float32, len(s))
	copy(copied, s)

	return copied
}

func InterpolateBufferedVertices(bv1 *BufferedVertices, bv2 *BufferedVertices, blend_ratio float32) *BufferedVertices {
	interpolated := NewBufferedVertices()

	interpolated.ambient_color = bv1.ambient_color
	interpolated.diffuse_color = bv1.diffuse_color
	interpolated.specular_color = bv1.specular_color
	interpolated.specular_exponent = bv1.specular_exponent
	interpolated.dissolve = bv1.dissolve
	interpolated.diffuse_texture_map = bv1.diffuse_texture_map

	texture_handle := bv1.texture_handle
	indices_buffer := bv1.indices_buffer
	uv_buffer := bv1.uv_buffer

	//Interpolate positions and normals.
	pos_buffer_1 := bv1.pos_buffer
	pos_buffer_2 := bv2.pos_buffer
	norm_buffer_1 := bv1.norm_buffer
	norm_buffer_2 := bv2.norm_buffer

	pos_buffer_len := len(pos_buffer_1)
	norm_buffer_len := len(norm_buffer_1)

	interpolated_pos_buffer := make([]float32, pos_buffer_len)
	interpolated_norm_buffer := make([]float32, norm_buffer_len)

	for i := 0; i < pos_buffer_len; i++ {
		interpolated_pos_buffer[i] = pos_buffer_1[i]*(1.0-blend_ratio) + pos_buffer_2[i]*blend_ratio
	}
	for i := 0; i < norm_buffer_len; i++ {
		interpolated_norm_buffer[i] = norm_buffer_1[i]*(1.0-blend_ratio) + norm_buffer_2[i]*blend_ratio
	}

	interpolated.texture_handle = texture_handle
	interpolated.indices_buffer = indices_buffer
	interpolated.uv_buffer = uv_buffer
	interpolated.pos_buffer = interpolated_pos_buffer
	interpolated.norm_buffer = interpolated_norm_buffer

	return interpolated
}

func (bv *BufferedVertices) SetTextureHandle(texture_handle int) {
	bv.texture_handle = texture_handle
}
func (bv *BufferedVertices) SetIndicesBuffer(indices_buffer []uint32) {
	bv.indices_buffer = indices_buffer
	bv.indices_count = int32(len(indices_buffer))
}
func (bv *BufferedVertices) SetPosBuffer(pos_buffer []float32) {
	bv.pos_buffer = pos_buffer
}
func (bv *BufferedVertices) SetUVBuffer(uv_buffer []float32) {
	bv.uv_buffer = uv_buffer
}
func (bv *BufferedVertices) SetNormBuffer(norm_buffer []float32) {
	bv.norm_buffer = norm_buffer
}
func (bv *BufferedVertices) SetAmbientColor(ambient_color coloru8.ColorU8) {
	bv.ambient_color = ambient_color
}
func (bv *BufferedVertices) SetDiffuseColor(diffuse_color coloru8.ColorU8) {
	bv.diffuse_color = diffuse_color
}
func (bv *BufferedVertices) SetSpecularColor(specular_color coloru8.ColorU8) {
	bv.specular_color = specular_color
}
func (bv *BufferedVertices) SetDissolve(dissolve float32) {
	bv.dissolve = dissolve
}
func (bv *BufferedVertices) SetDiffuseTextureMap(diffuse_texture_map string) {
	bv.diffuse_texture_map = diffuse_texture_map
}

func (bv *BufferedVertices) GetTextureHandle() int {
	return bv.texture_handle
}
func (bv *BufferedVertices) GetIndicesCount() int32 {
	return bv.indices_count
}
func (bv *BufferedVertices) GetIndicesBuffer() []uint32 {
	return bv.indices_buffer
}
func (bv *BufferedVertices) GetPosBuffer() []float32 {
	return bv.pos_buffer
}
func (bv *BufferedVertices) GetUVBuffer() []float32 {
	return bv.uv_buffer
}
func (bv *BufferedVertices) GetNormBuffer() []float32 {
	return bv.norm_buffer
}
func (bv *BufferedVertices) GetAmbientColor() coloru8.ColorU8 {
	return bv.ambient_color
}
func (bv *BufferedVertices) GetDiffuseColor() coloru8.ColorU8 {
	return bv.diffuse_color
}
func (bv *BufferedVertices) GetSpecularColor() coloru8.ColorU8 {
	return bv.specular_color
}
func (bv *BufferedVertices) GetSpecularExponent() float32 {
	return bv.specular_exponent
}
func (bv *BufferedVertices) GetDissolve() float32 {
	return bv.dissolve
}
func (bv *BufferedVertices) GetDiffuseTextureMap() string {
	return bv.diffuse_texture_map
}
