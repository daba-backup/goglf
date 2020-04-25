package dhmatrixtool

import (
	"github.com/dabasan/go-dh3dbasis/dhmatrix"
	"github.com/dabasan/go-dh3dbasis/dhvector"
)

func GetViewTransformationMatrix(pos dhvector.Vector, target dhvector.Vector, up dhvector.Vector) dhmatrix.Matrix {
	var view_transformation_matrix dhmatrix.Matrix

	vec_translate := dhvector.VScale(pos, -1.0)
	t := dhmatrix.MGetTranslate(vec_translate)

	view_coord_z := dhvector.VSub(pos, target)
	view_coord_z = dhvector.VNorm(view_coord_z)
	view_coord_x := dhvector.VCross(up, view_coord_z)
	view_coord_x = dhvector.VNorm(view_coord_x)
	view_coord_y := dhvector.VCross(view_coord_z, view_coord_x)

	var r dhmatrix.Matrix
	r.M[0][0] = view_coord_x.X
	r.M[0][1] = view_coord_x.Y
	r.M[0][2] = view_coord_x.Z
	r.M[0][3] = 0.0
	r.M[1][0] = view_coord_y.X
	r.M[1][1] = view_coord_y.Y
	r.M[1][2] = view_coord_y.Z
	r.M[1][3] = 0.0
	r.M[2][0] = view_coord_z.X
	r.M[2][1] = view_coord_z.Y
	r.M[2][2] = view_coord_z.Z
	r.M[2][3] = 0.0
	r.M[3][0] = 0.0
	r.M[3][1] = 0.0
	r.M[3][2] = 0.0
	r.M[3][3] = 1.0

	view_transformation_matrix = dhmatrix.MMult(r, t)

	return view_transformation_matrix
}
