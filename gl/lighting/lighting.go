package lighting

import (
	"github.com/dabasan/go-dh3dbasis/coloru8"
	"github.com/dabasan/go-dh3dbasis/vector"
	"github.com/dabasan/goglf/gl/shader"
)

type Lighting struct {
	direction     vector.Vector
	ambient_color coloru8.ColorU8

	diffuse_power  float32
	specular_power float32

	programs []*shader.ShaderProgram
}

func NewLighting() *Lighting {
	lighting := new(Lighting)

	lighting.direction = vector.VNorm(vector.VGet(1.0, -1.0, 1.0))
	lighting.ambient_color = coloru8.GetColorU8FromFloat32Components(0.6, 0.6, 0.6, 0.6)

	lighting.diffuse_power = 0.3
	lighting.specular_power = 0.1

	lighting.programs = make([]*shader.ShaderProgram, 0)

	return lighting
}

func (l *Lighting) AddProgram(program *shader.ShaderProgram) {
	l.programs = append(l.programs, program)
}
func (l *Lighting) RemoveAllPrograms() {
	l.programs = make([]*shader.ShaderProgram, 0)
}

func (l *Lighting) SetLightDirection(direction vector.Vector) {
	l.direction = direction
}
func (l *Lighting) SetLightDirection_PositionAndTarget(position vector.Vector, target vector.Vector) {
	l.direction = vector.VSub(target, position)
	l.direction = vector.VNorm(l.direction)
}
func (l *Lighting) SetAmbientColor(ambient_color coloru8.ColorU8) {
	l.ambient_color = ambient_color
}
func (l *Lighting) SetDiffusePower(diffuse_power float32) {
	l.diffuse_power = diffuse_power
}
func (l *Lighting) SetSpecularPower(specular_power float32) {
	l.specular_power = specular_power
}

func (l *Lighting) Update() {
	for _, program := range l.programs {
		program.Enable()
		program.SetUniformVector("light_direction", l.direction)
		program.SetUniformColorU8("ambient_color", l.ambient_color)
		program.SetUniform1f("diffuse_power", l.diffuse_power)
		program.SetUniform1f("specular_power", l.specular_power)
		program.Disable()
	}
}
