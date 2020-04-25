package matrixtool

import (
	"math"

	"github.com/dabasan/go-dh3dbasis/matrix"
	"github.com/dabasan/go-dh3dbasis/vector"
)

func GetOrthogonalMatrix(
	left float32, right float32,
	bottom float32, top float32,
	near float32, far float32) matrix.Matrix {
	vec_translate := vector.VGet(-(right+left)/2.0, -(bottom+top)/2.0, (far+near)/2.0)
	vec_scale := vector.VGet(2.0/(right-left), 2.0/(top-bottom), -2.0/(far-near))

	mat_translate := matrix.MGetTranslate(vec_translate)
	mat_scale := matrix.MGetScale(vec_scale)

	orthogonal_matrix := matrix.MMult(mat_scale, mat_translate)

	return orthogonal_matrix
}
func GetPerspectiveMatrix(
	left float32, right float32,
	bottom float32, top float32,
	near float32, far float32) matrix.Matrix {
	var perspective_matrix matrix.Matrix

	perspective_matrix.M[0][0] = 2.0 * near / (right - left)
	perspective_matrix.M[0][1] = 0.0
	perspective_matrix.M[0][2] = (right + left) / (right - left)
	perspective_matrix.M[0][3] = 0.0
	perspective_matrix.M[1][0] = 0.0
	perspective_matrix.M[1][1] = 2.0 * near / (top - bottom)
	perspective_matrix.M[1][2] = (top * bottom) / (top - bottom)
	perspective_matrix.M[1][3] = 0.0
	perspective_matrix.M[2][0] = 0.0
	perspective_matrix.M[2][1] = 0.0
	perspective_matrix.M[2][2] = -(far + near) / (far - near)
	perspective_matrix.M[2][3] = -2.0 * far * near / (far - near)
	perspective_matrix.M[3][0] = 0.0
	perspective_matrix.M[3][1] = 0.0
	perspective_matrix.M[3][2] = -1.0
	perspective_matrix.M[3][3] = 0.0

	return perspective_matrix
}
func GetPerspectiveMatrix_FovAndAspect(fov float32, aspect float32, near float32, far float32) matrix.Matrix {
	var perspective_matrix matrix.Matrix

	f := float32(1.0 / math.Tan(float64(fov/2.0)))

	perspective_matrix.M[0][0] = f / aspect
	perspective_matrix.M[0][1] = 0.0
	perspective_matrix.M[0][2] = 0.0
	perspective_matrix.M[0][3] = 0.0
	perspective_matrix.M[1][0] = 0.0
	perspective_matrix.M[1][1] = f
	perspective_matrix.M[1][2] = 0.0
	perspective_matrix.M[1][3] = 0.0
	perspective_matrix.M[2][0] = 0.0
	perspective_matrix.M[2][1] = 0.0
	perspective_matrix.M[2][2] = -(far + near) / (far - near)
	perspective_matrix.M[2][3] = -2.0 * far * near / (far - near)
	perspective_matrix.M[3][0] = 0.0
	perspective_matrix.M[3][1] = 0.0
	perspective_matrix.M[3][2] = -1.0
	perspective_matrix.M[3][3] = 0.0

	return perspective_matrix
}
