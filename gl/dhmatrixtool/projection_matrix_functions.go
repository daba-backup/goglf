package dhmatrixtool

import (
	"github.com/dabasan/go-dh3dbasis/dhmatrix"
	"github.com/dabasan/go-dh3dbasis/dhvector"
)

func GetOrthogonalMatrix(
	left float32, right float32,
	bottom float32, top float32,
	near float32, far float32) dhmatrix.Matrix {
	vec_translate := dhvector.VGet(-(right+left)/2.0, -(bottom+top)/2.0, (far+near)/2.0)
	vec_scale := dhvector.VGet(2.0/(right-left), 2.0/(top-bottom), -2.0/(far-near))

	mat_translate := dhmatrix.MGetTranslate(vec_translate)
	mat_scale := dhmatrix.MGetScale(vec_scale)

	orthogonal_matrix := dhmatrix.MMult(mat_scale, mat_translate)
	return orthogonal_matrix
}
