package shape

import (
	"github.com/dabasan/go-dh3dbasis/coloru8"
	"github.com/dabasan/go-dh3dbasis/vector"
)

type Vertex3D struct {
	Pos  vector.Vector
	Norm vector.Vector
	Dif  coloru8.ColorU8
	Spc  coloru8.ColorU8
	U    float32
	V    float32
}
