package shader

import (
	"log"

	"github.com/dabasan/goglf/gl/texture"

	"github.com/dabasan/go-dh3dbasis/coloru8"
	"github.com/dabasan/go-dh3dbasis/matrix"
	"github.com/dabasan/go-dh3dbasis/vector"

	"github.com/dabasan/goglf/gl/wrapper"
	"github.com/go-gl/gl/all-core/gl"
)

type ShaderProgram struct {
	program_name         string
	program_id           uint32
	logging_enabled_flag bool
}

func NewShaderProgram(program_name string) (*ShaderProgram, bool) {
	program := new(ShaderProgram)
	program.program_name = program_name
	program.logging_enabled_flag = false

	program_id, ok := GetProgramID(program_name)
	if ok == false {
		log.Printf("warn: This program is invalid. program_name=%v", program_name)
		return program, false
	}
	program.program_id = program_id

	return program, true
}

func (p *ShaderProgram) EnableLogging(flag bool) {
	p.logging_enabled_flag = flag
}

func (p *ShaderProgram) Enable() {
	wrapper.UseProgram(p.program_id)
}
func (p *ShaderProgram) Disable() {
	wrapper.UseProgram(0)
}

func (p *ShaderProgram) SetUniform1i(name string, value int32) int {
	location := wrapper.GetUniformLocation(p.program_id, gl.Str(name+"\x00"))
	if location < 0 {
		if p.logging_enabled_flag == true {
			log.Printf("trace: (%v) Invalid uniform name. name=%v", p.program_name, name)
		}
		return -1
	}

	wrapper.Uniform1i(location, value)

	return 0
}
func (p *ShaderProgram) SetUniform2i(name string, value0 int32, value1 int32) int {
	location := wrapper.GetUniformLocation(p.program_id, gl.Str(name+"\x00"))
	if location < 0 {
		if p.logging_enabled_flag == true {
			log.Printf("trace: (%v) Invalid uniform name. name=%v", p.program_name, name)
		}
		return -1
	}

	wrapper.Uniform2i(location, value0, value1)

	return 0
}
func (p *ShaderProgram) SetUniform3i(name string, value0 int32, value1 int32, value2 int32) int {
	location := wrapper.GetUniformLocation(p.program_id, gl.Str(name+"\x00"))
	if location < 0 {
		if p.logging_enabled_flag == true {
			log.Printf("trace: (%v) Invalid uniform name. name=%v", p.program_name, name)
		}
		return -1
	}

	wrapper.Uniform3i(location, value0, value1, value2)

	return 0
}
func (p *ShaderProgram) SetUniform4i(name string, value0 int32, value1 int32, value2 int32, value3 int32) int {
	location := wrapper.GetUniformLocation(p.program_id, gl.Str(name+"\x00"))
	if location < 0 {
		if p.logging_enabled_flag == true {
			log.Printf("trace: (%v) Invalid uniform name. name=%v", p.program_name, name)
		}
		return -1
	}

	wrapper.Uniform4i(location, value0, value1, value2, value3)

	return 0
}
func (p *ShaderProgram) SetUniform1f(name string, value float32) int {
	location := wrapper.GetUniformLocation(p.program_id, gl.Str(name+"\x00"))
	if location < 0 {
		if p.logging_enabled_flag == true {
			log.Printf("trace: (%v) Invalid uniform name. name=%v", p.program_name, name)
		}
		return -1
	}

	wrapper.Uniform1f(location, value)

	return 0
}
func (p *ShaderProgram) SetUniform2f(name string, value0 float32, value1 float32) int {
	location := wrapper.GetUniformLocation(p.program_id, gl.Str(name+"\x00"))
	if location < 0 {
		if p.logging_enabled_flag == true {
			log.Printf("trace: (%v) Invalid uniform name. name=%v", p.program_name, name)
		}
		return -1
	}

	wrapper.Uniform2f(location, value0, value1)

	return 0
}
func (p *ShaderProgram) SetUniform3f(name string, value0 float32, value1 float32, value2 float32) int {
	location := wrapper.GetUniformLocation(p.program_id, gl.Str(name+"\x00"))
	if location < 0 {
		if p.logging_enabled_flag == true {
			log.Printf("trace: (%v) Invalid uniform name. name=%v", p.program_name, name)
		}
		return -1
	}

	wrapper.Uniform3f(location, value0, value1, value2)

	return 0
}
func (p *ShaderProgram) SetUniform4f(name string, value0 float32, value1 float32, value2 float32, value3 float32) int {
	location := wrapper.GetUniformLocation(p.program_id, gl.Str(name+"\x00"))
	if location < 0 {
		if p.logging_enabled_flag == true {
			log.Printf("trace: (%v) Invalid uniform name. name=%v", p.program_name, name)
		}
		return -1
	}

	wrapper.Uniform4f(location, value0, value1, value2, value3)

	return 0
}
func (p *ShaderProgram) SetUniformVector(name string, value vector.Vector) int {
	location := wrapper.GetUniformLocation(p.program_id, gl.Str(name+"\x00"))
	if location < 0 {
		if p.logging_enabled_flag == true {
			log.Printf("trace: (%v) Invalid uniform name. name=%v", p.program_name, name)
		}
		return -1
	}

	wrapper.Uniform3f(location, value.X, value.Y, value.Z)

	return 0
}
func (p *ShaderProgram) SetUniformColorU8(name string, value coloru8.ColorU8) int {
	location := wrapper.GetUniformLocation(p.program_id, gl.Str(name+"\x00"))
	if location < 0 {
		if p.logging_enabled_flag == true {
			log.Printf("trace: (%v) Invalid uniform name. name=%v", p.program_name, name)
		}
		return -1
	}

	wrapper.Uniform4f(location, value.R, value.G, value.B, value.A)

	return 0
}
func (p *ShaderProgram) SetUniformMatrix(name string, transpose bool, value matrix.Matrix) int {
	location := wrapper.GetUniformLocation(p.program_id, gl.Str(name+"\x00"))
	if location < 0 {
		if p.logging_enabled_flag == true {
			log.Printf("trace: (%v) Invalid uniform name. name=%v", p.program_name, name)
		}
		return -1
	}

	wrapper.UniformMatrix4fv(location, 1, transpose, &value.M[0][0])

	return 0
}

func (p *ShaderProgram) SetTexture(name string, texture_unit int, texture_handle int) int {
	location := wrapper.GetUniformLocation(p.program_id, gl.Str(name+"\x00"))
	if location < 0 {
		if p.logging_enabled_flag == true {
			log.Printf("trace: (%v) Invalid uniform name. name=%v", p.program_name, name)
		}
		return -1
	}

	wrapper.ActiveTexture(gl.TEXTURE0 + texture_unit)
	texture.BindTexture(texture_handle)
	wrapper.Uniform1i(location, texture_unit)
	texture.UnbindTexture()

	return 0
}
