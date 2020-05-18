package assimp

import (
	"fmt"
	"log"

	assimpgo "github.com/dabasan/assimp-go/assimp"
	"github.com/dabasan/goglf/gl/model/buffer"
)

func ShowVersion() {
	major := assimpgo.GetVersionMajor()
	minor := assimpgo.GetVersionMinor()
	revision := assimpgo.GetVersionRevision()
	fmt.Printf("Version: %v.%v.%v\n", major, minor, revision)
}
func ShowLegalString() {
	fmt.Println(assimpgo.GetLegalString())
}

//LoadModelWithAssimp loads a model with Assimp and generates BufferedVertices.
//This function does not concern textures, so load textures on your own.
func LoadModelWithAssimp(model_filename string) ([]*buffer.BufferedVertices, error) {
	log.Printf("info: Start loading a model with Assimp. model_filename=%v", model_filename)

	ret := make([]*buffer.BufferedVertices, 0)

	meshes, err := assimpgo.ParseFile(model_filename)
	if err != nil {
		return nil, err
	}

	for _, mesh := range meshes {
		vertex_num := int(mesh.VertexCount)

		buffered_vertices := buffer.NewBufferedVertices()
		indices_buffer := make([]uint32, vertex_num)
		pos_buffer := make([]float32, vertex_num*3)
		uv_buffer := make([]float32, vertex_num*2)
		norm_buffer := make([]float32, vertex_num*3)

		vec_count := 0
		uv_count := 0
		for i := 0; i < vertex_num; i++ {
			indices_buffer[i] = uint32(i)

			pos_buffer[vec_count] = mesh.Vertices[i].X()
			pos_buffer[vec_count+1] = mesh.Vertices[i].Y()
			pos_buffer[vec_count+2] = mesh.Vertices[i].Z()
			uv_buffer[uv_count] = mesh.UVChannels[0][i].X()
			uv_buffer[uv_count+1] = mesh.UVChannels[0][i].Y()
			norm_buffer[vec_count] = mesh.Normals[i].X()
			norm_buffer[vec_count+1] = mesh.Normals[i].Y()
			norm_buffer[vec_count+2] = mesh.Normals[i].Z()

			vec_count += 3
			uv_count += 2
		}

		buffered_vertices.SetIndicesBuffer(indices_buffer)
		buffered_vertices.SetPosBuffer(pos_buffer)
		buffered_vertices.SetUVBuffer(uv_buffer)
		buffered_vertices.SetNormBuffer(norm_buffer)

		ret = append(ret, buffered_vertices)
	}

	return ret, nil
}
