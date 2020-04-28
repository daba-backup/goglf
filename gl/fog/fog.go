package fog

import (
	"github.com/dabasan/go-dh3dbasis/coloru8"
	"github.com/dabasan/goglf/gl/shader"
)

type Fog struct {
	start float32
	end   float32
	color coloru8.ColorU8

	programs []*shader.ShaderProgram
}

func NewFog() *Fog {
	fog := new(Fog)

	fog.start = 100.0
	fog.end = 200.0
	fog.color = coloru8.GetColorU8FromFloat32Components(0.0, 0.0, 0.0, 1.0)

	fog.programs = make([]*shader.ShaderProgram, 0)

	return fog
}

func (f *Fog) AddProgram(program *shader.ShaderProgram) {
	f.programs = append(f.programs, program)
}
func (f *Fog) RemoveAllPrograms() {
	f.programs = make([]*shader.ShaderProgram, 0)
}

func (f *Fog) SetFogStartEnd(start float32, end float32) {
	f.start = start
	f.end = end
}
func (f *Fog) SetFogColor(color coloru8.ColorU8) {
	f.color = color
}

func (f *Fog) Update() {
	for _, program := range f.programs {
		program.Enable()
		program.SetUniform1f("fog.start", f.start)
		program.SetUniform1f("fog.end", f.end)
		program.SetUniformColorU8("fog.color", f.color)
		program.Disable()
	}
}
