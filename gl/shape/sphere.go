package shape

import (
	"github.com/dabasan/go-dh3dbasis/coloru8"
	"github.com/dabasan/go-dh3dbasis/vector"
)

type Sphere struct {
	Center    vector.Vector
	Radius    float32
	Slice_Num int
	Stack_Num int
	Color     coloru8.ColorU8
}
