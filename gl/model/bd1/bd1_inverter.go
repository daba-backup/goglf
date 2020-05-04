package bd1

func invertZ(blocks []*bd1Block) {
	for _, block := range blocks {
		vertex_positions := &block.Vertex_positions
		for i := 0; i < 8; i++ {
			vertex_positions[i].Z = vertex_positions[i].Z * (-1.0)
		}
	}

	for _, block := range blocks {
		vertex_positions := &block.Vertex_positions
		us := &block.Us
		vs := &block.Vs
		texture_ids := &block.Texture_ids

		vertex_positions_orig := block.Vertex_positions
		for i := 0; i < 4; i++ {
			vertex_positions[i] = vertex_positions_orig[3-i]
		}
		for i := 0; i < 4; i++ {
			vertex_positions[i+4] = vertex_positions_orig[7-i]
		}

		us_orig := block.Us
		vs_orig := block.Vs

		for i := 0; i < 6; i++ {
			var uv_indices []int

			if i == 2 {
				uv_indices = getFaceCorrespondingUVIndices(4)
			} else if i == 4 {
				uv_indices = getFaceCorrespondingUVIndices(2)
			} else {
				uv_indices = getFaceCorrespondingUVIndices(i)
			}

			for j := 0; j < 4; j++ {
				index := i*4 + j

				us[index] = us_orig[uv_indices[j]]
				vs[index] = vs_orig[uv_indices[j]]
			}
		}

		texture_ids_orig := block.Texture_ids
		texture_ids[2] = texture_ids_orig[4]
		texture_ids[4] = texture_ids_orig[2]
	}
}
