package bd1

import (
	"github.com/dabasan/go-dh3dbasis/vector"
)

type bd1Block struct {
	Vertex_positions [8]vector.Vector
	Us               [24]float32
	Vs               [24]float32
	Texture_ids      [6]int
	Enabled_flag     bool
}
