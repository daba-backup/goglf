package front

import (
	"github.com/dabasan/go-dh3dbasis/coloru8"
	"github.com/dabasan/go-dh3dbasis/vector"
	"github.com/dabasan/goglf/gl/lighting"
	"github.com/dabasan/goglf/gl/shader"
)

var l *lighting.Lighting

func init() {
	l = lighting.NewLighting()
}

func AddProgram(program *shader.ShaderProgram) {
	l.AddProgram(program)
}
func RemoveAllPrograms() {
	l.RemoveAllPrograms()
}

func SetAmbientColor(color coloru8.ColorU8) {
	l.SetAmbientColor(color)
}
func SetLightDirection(direction vector.Vector) {
	l.SetLightDirection(direction)
}
func SetLightDirection_PositionAndTarget(position vector.Vector, target vector.Vector) {
	l.SetLightDirection_PositionAndTarget(position, target)
}
func SetDiffusePower(diffuse_power float32) {
	l.SetDiffusePower(diffuse_power)
}
func SetSpecularPower(specular_power float32) {
	l.SetSpecularPower(specular_power)
}

func Update() {
	l.Update()
}
