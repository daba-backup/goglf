package front

import (
	"github.com/dabasan/go-dh3dbasis/coloru8"
	"github.com/dabasan/goglf/gl/fog"
	"github.com/dabasan/goglf/gl/shader"
)

var f *fog.Fog

func init() {
	f = fog.NewFog()
}

func AddProgram(program *shader.ShaderProgram) {
	f.AddProgram(program)
}
func RemoveAllPrograms() {
	f.RemoveAllPrograms()
}

func SetFogColor(color coloru8.ColorU8) {
	f.SetFogColor(color)
}
func SetFogStartEnd(start float32, end float32) {
	f.SetFogStartEnd(start, end)
}

func Update() {
	f.Update()
}
