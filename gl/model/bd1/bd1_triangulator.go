package bd1

import (
	"log"

	"github.com/dabasan/go-dh3dbasis/coloru8"
	"github.com/dabasan/go-dh3dbasis/vector"
	"github.com/dabasan/goglf/gl/shape"
)

type bd1Triangulator struct {
	triangles_list []*bd1Triangle
}

func newBD1Triangulator() *bd1Triangulator {
	t := new(bd1Triangulator)
	t.triangles_list = make([]*bd1Triangle, 0)

	return t
}

func (t *bd1Triangulator) triangulateBlock(block *bd1Block) {
	if block == nil {
		log.Printf("warn: Nil argument where non-nil required.")
		return
	}

	positions := &block.Vertex_positions
	us := &block.Us
	vs := &block.Vs
	texture_ids := &block.Texture_ids

	var vertices [24]shape.Vertex3D

	for i := 0; i < 6; i++ {
		vertex_indices := getFaceCorrespondingVertexIndices(i)
		uv_indices := getFaceCorrespondingUVIndices(i)

		v1 := vector.VSub(positions[vertex_indices[3]], positions[vertex_indices[0]])
		v2 := vector.VSub(positions[vertex_indices[1]], positions[vertex_indices[0]])

		face_normal := vector.VCross(v1, v2)
		face_normal = vector.VNorm(face_normal)

		for j := 0; j < 4; j++ {
			array_index := i*4 + j
			vertex_index := vertex_indices[j]
			uv_index := uv_indices[j]

			vertices[array_index].Pos = positions[vertex_index]
			vertices[array_index].Dif = coloru8.GetColorU8FromFloat32Components(1.0, 1.0, 1.0, 1.0)
			vertices[array_index].Spc = coloru8.GetColorU8FromFloat32Components(0.0, 0.0, 0.0, 0.0)
			vertices[array_index].U = us[uv_index]
			vertices[array_index].V = vs[uv_index] * (-1.0)

			vertices[array_index].Norm = face_normal
		}
	}

	var inverted_vertices [24]*shape.Vertex3D
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			inverted_vertices[i*4+j] = &vertices[(i+1)*4-1-j]
		}
	}

	var triangles [12]bd1Triangle
	for i := 0; i < 12; i += 2 {
		var vertex_array_index int
		for j := 0; j < 3; j++ {
			vertex_array_index = (i/2)*4 + j
			triangles[i].Triangle.Vertices[j] = *inverted_vertices[vertex_array_index]
		}
		for j := 0; j < 3; j++ {
			vertex_array_index = (i/2)*4 + (j+2)%4
			triangles[i+1].Triangle.Vertices[j] = *inverted_vertices[vertex_array_index]
		}
	}

	for i := 0; i < 12; i += 2 {
		texture_id := texture_ids[i/2]

		triangles[i].Texture_id = texture_id
		triangles[i+1].Texture_id = texture_id

		t.triangles_list = append(t.triangles_list, &triangles[i])
		t.triangles_list = append(t.triangles_list, &triangles[i+1])
	}
}
func (t *bd1Triangulator) triangulateBlocks(blocks []*bd1Block) {
	if blocks == nil {
		log.Printf("warn: Nil argument where non-nil required.")
		return
	}

	for _, block := range blocks {
		t.triangulateBlock(block)
	}
}

func (t *bd1Triangulator) getTriangleList() []*bd1Triangle {
	return t.triangles_list
}
func (t *bd1Triangulator) getTrianglesMap() map[int][]*bd1Triangle {
	ret := make(map[int][]*bd1Triangle)

	for _, triangle := range t.triangles_list {
		texture_id := triangle.Texture_id

		if _, ok := ret[texture_id]; !ok {
			list_temp := make([]*bd1Triangle, 0)
			ret[texture_id] = list_temp
		}
		triangles := ret[texture_id]
		triangles = append(triangles, triangle)
	}

	return ret
}
