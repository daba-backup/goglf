package dhbuffertool

import (
	"github.com/dabasan/go-dh3dbasis/dhcoloru8"
	"github.com/dabasan/go-dh3dbasis/dhmatrix"
	"github.com/dabasan/go-dh3dbasis/dhvector"
)

func MakeFloat32SliceFromColorU8(c dhcoloru8.ColorU8) []float32 {
	ret := make([]float32, 4)

	ret[0] = c.R
	ret[1] = c.G
	ret[2] = c.B
	ret[3] = c.A

	return ret
}
func MakeFloat32SliceFromVector(v dhvector.Vector) []float32 {
	ret := make([]float32, 3)

	ret[0] = v.X
	ret[1] = v.Y
	ret[2] = v.Z

	return ret
}
func MakeFloat32SliceFromMatrix(m dhmatrix.Matrix) []float32 {
	ret := make([]float32, 16)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			ret[i*4+j] = m.M[i][j]
		}
	}

	return ret
}
