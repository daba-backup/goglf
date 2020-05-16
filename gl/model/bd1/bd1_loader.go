package bd1

import (
	"github.com/dabasan/go-dhtool/filename"
	"github.com/dabasan/goglf/gl/model/buffer"
	"github.com/dabasan/goglf/gl/texture"
)

func LoadBD1(bd1_filename string) ([]*buffer.BufferedVertices, error) {
	ret := make([]*buffer.BufferedVertices, 0)

	reader := newbd1Reader()
	err := reader.read(bd1_filename)
	if err != nil {
		return nil, err
	}

	bd1_directory := filename.GetFileDirectory(bd1_filename)
	texture_filenames := reader.getTextureFilenames()
	texture_handles_map := make(map[int]int)
	for texture_id, texture_filename := range texture_filenames {
		if texture_filename == "" {
			texture_handles_map[texture_id] = -1
			continue
		}

		texture_filename = bd1_directory + "/" + texture_filename

		texture_handle := texture.LoadTexture(texture_filename)
		texture_handles_map[texture_id] = texture_handle
	}

	blocks := reader.getBlocks()
	invertZ(blocks)

	triangulator := newBD1Triangulator()
	triangulator.triangulateBlocks(blocks)

	triangles_map := triangulator.getTrianglesMap()
	for texture_id, triangles := range triangles_map {
		texture_handle, ok := texture_handles_map[texture_id]
		if !ok {
			texture_handle = -1
		}

		triangle_num := len(triangles)
		vertex_num := triangle_num * 3

		buffered_vertices := buffer.NewBufferedVertices()
		indices_buffer := make([]uint32, vertex_num)
		pos_buffer := make([]float32, vertex_num*3)
		uv_buffer := make([]float32, vertex_num*2)
		norm_buffer := make([]float32, vertex_num*3)

		indices_count := 0
		vec_count := 0
		uv_count := 0
		for i := 0; i < triangle_num; i++ {
			triangle := triangles[i]

			for j := 0; j < 3; j++ {
				vertex := &triangle.Triangle.Vertices[j]

				position := &vertex.Pos
				u := vertex.U
				v := vertex.V
				normal := &vertex.Norm

				indices_buffer[indices_count] = uint32(i*3 + j)
				pos_buffer[vec_count] = position.X
				pos_buffer[vec_count+1] = position.Y
				pos_buffer[vec_count+2] = position.Z
				uv_buffer[uv_count] = u
				uv_buffer[uv_count+1] = v
				norm_buffer[vec_count] = normal.X
				norm_buffer[vec_count+1] = normal.Y
				norm_buffer[vec_count+2] = normal.Z

				indices_count++
				vec_count += 3
				uv_count += 2
			}
		}

		buffered_vertices.SetIndicesBuffer(indices_buffer)
		buffered_vertices.SetPosBuffer(pos_buffer)
		buffered_vertices.SetUVBuffer(uv_buffer)
		buffered_vertices.SetNormBuffer(norm_buffer)
		buffered_vertices.SetTextureHandle(texture_handle)

		ret = append(ret, buffered_vertices)
	}

	return ret, nil
}
func LoadBD1_KeepOrder(bd1_filename string) ([]*buffer.BufferedVertices, error) {
	ret := make([]*buffer.BufferedVertices, 0)

	reader := newbd1Reader()
	err := reader.read(bd1_filename)
	if err != nil {
		return ret, err
	}

	bd1_directory := filename.GetFileDirectory(bd1_filename)
	texture_filenames := reader.getTextureFilenames()
	texture_handles_map := make(map[int]int)
	for texture_id, texture_filename := range texture_filenames {
		if texture_filename == "" {
			texture_handles_map[texture_id] = -1
			continue
		}

		texture_filename = bd1_directory + "/" + texture_filename

		texture_handle := texture.LoadTexture(texture_filename)
		texture_handles_map[texture_id] = texture_handle
	}

	blocks := reader.getBlocks()
	invertZ(blocks)

	triangulator := newBD1Triangulator()
	triangulator.triangulateBlocks(blocks)

	triangles_list := triangulator.getTriangleList()
	for _, triangle := range triangles_list {
		texture_id := triangle.Texture_id
		texture_handle, ok := texture_handles_map[texture_id]
		if !ok {
			texture_handle = -1
		}

		vertex_num := 3

		buffered_vertices := buffer.NewBufferedVertices()
		indices_buffer := make([]uint32, vertex_num)
		pos_buffer := make([]float32, vertex_num*3)
		uv_buffer := make([]float32, vertex_num*2)
		norm_buffer := make([]float32, vertex_num*3)

		indices_count := 0
		vec_count := 0
		uv_count := 0
		for i := 0; i < 3; i++ {
			vertex := &triangle.Triangle.Vertices[i]

			position := &vertex.Pos
			u := vertex.U
			v := vertex.V
			normal := &vertex.Norm

			indices_buffer[indices_count] = uint32(i)
			pos_buffer[vec_count] = position.X
			pos_buffer[vec_count+1] = position.Y
			pos_buffer[vec_count+2] = position.Z
			uv_buffer[uv_count] = u
			uv_buffer[uv_count] = v
			norm_buffer[vec_count] = normal.X
			norm_buffer[vec_count+1] = normal.Y
			norm_buffer[vec_count+2] = normal.Z

			indices_count++
			vec_count += 3
			uv_count += 2
		}

		buffered_vertices.SetIndicesBuffer(indices_buffer)
		buffered_vertices.SetPosBuffer(pos_buffer)
		buffered_vertices.SetUVBuffer(uv_buffer)
		buffered_vertices.SetNormBuffer(norm_buffer)
		buffered_vertices.SetTextureHandle(texture_handle)

		ret = append(ret, buffered_vertices)
	}

	return ret, nil
}
