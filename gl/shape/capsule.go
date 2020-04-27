package shape

import (
	"github.com/dabasan/go-dh3dbasis/coloru8"
	"github.com/dabasan/go-dh3dbasis/vector"
)

type Capsule struct {
	Pos_1     vector.Vector
	Pos_2     vector.Vector
	Radius    float32
	Slice_num int
	Stack_num int
	Color     coloru8.ColorU8
}
