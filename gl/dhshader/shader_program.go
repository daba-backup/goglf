package dhshader

import (
	"errors"
	"log"

	"github.com/dabasan/go-dh3dbasis/dhcoloru8"
	"github.com/dabasan/go-dh3dbasis/dhmatrix"
	"github.com/dabasan/go-dh3dbasis/dhvector"

	"github.com/dabasan/goglf/gl/dhwrapper"
	"github.com/go-gl/gl/all-core/gl"
)

type ShaderProgram struct {
	program_name string
	program_id   uint32
}

func NewShaderProgram(program_name string) (*ShaderProgram, error) {
	program_id, ok := GetProgramID(program_name)
	if ok == false {
		log.Printf("warn: This program is invalid. program_name=%v", program_name)
		return nil, errors.New("This program is invalid.")
	}

	program := new(ShaderProgram)
	program.program_name = program_name
	program.program_id = program_id

	return program, nil
}

func (p *ShaderProgram) EnableProgram() {
	dhwrapper.UseProgram(p.program_id)
}
func (p *ShaderProgram) DisableProgram() {
	dhwrapper.UseProgram(0)
}

func (p *ShaderProgram) SetUniform1i(name string, value int32) int {
	location := dhwrapper.GetUniformLocation(p.program_id, gl.Str(name))
	if location < 0 {
		log.Printf("debug: (%v) Invalid uniform name. name=%v", p.program_name, name)
		return -1
	}

	dhwrapper.Uniform1i(location, value)

	return 0
}
func (p *ShaderProgram) SetUniform2i(name string, value0 int32, value1 int32) int {
	location := dhwrapper.GetUniformLocation(p.program_id, gl.Str(name))
	if location < 0 {
		log.Printf("debug: (%v) Invalid uniform name. name=%v", p.program_name, name)
		return -1
	}

	dhwrapper.Uniform2i(location, value0, value1)

	return 0
}
func (p *ShaderProgram) SetUniform3i(name string, value0 int32, value1 int32, value2 int32) int {
	location := dhwrapper.GetUniformLocation(p.program_id, gl.Str(name))
	if location < 0 {
		log.Printf("debug: (%v) Invalid uniform name. name=%v", p.program_name, name)
		return -1
	}

	dhwrapper.Uniform3i(location, value0, value1, value2)

	return 0
}
func (p *ShaderProgram) SetUniform4i(name string, value0 int32, value1 int32, value2 int32, value3 int32) int {
	location := dhwrapper.GetUniformLocation(p.program_id, gl.Str(name))
	if location < 0 {
		log.Printf("debug: (%v) Invalid uniform name. name=%v", p.program_name, name)
		return -1
	}

	dhwrapper.Uniform4i(location, value0, value1, value2, value3)

	return 0
}
func (p *ShaderProgram) SetUniform1f(name string, value float32) int {
	location := dhwrapper.GetUniformLocation(p.program_id, gl.Str(name))
	if location < 0 {
		log.Printf("debug: (%v) Invalid uniform name. name=%v", p.program_name, name)
		return -1
	}

	dhwrapper.Uniform1f(location, value)

	return 0
}
func (p *ShaderProgram) SetUniform2f(name string, value0 float32, value1 float32) int {
	location := dhwrapper.GetUniformLocation(p.program_id, gl.Str(name))
	if location < 0 {
		log.Printf("debug: (%v) Invalid uniform name. name=%v", p.program_name, name)
		return -1
	}

	dhwrapper.Uniform2f(location, value0, value1)

	return 0
}
func (p *ShaderProgram) SetUniform3f(name string, value0 float32, value1 float32, value2 float32) int {
	location := dhwrapper.GetUniformLocation(p.program_id, gl.Str(name))
	if location < 0 {
		log.Printf("debug: (%v) Invalid uniform name. name=%v", p.program_name, name)
		return -1
	}

	dhwrapper.Uniform3f(location, value0, value1, value2)

	return 0
}
func (p *ShaderProgram) SetUniform4f(name string, value0 float32, value1 float32, value2 float32, value3 float32) int {
	location := dhwrapper.GetUniformLocation(p.program_id, gl.Str(name))
	if location < 0 {
		log.Printf("debug: (%v) Invalid uniform name. name=%v", p.program_name, name)
		return -1
	}

	dhwrapper.Uniform4f(location, value0, value1, value2, value3)

	return 0
}
func (p *ShaderProgram) SetUniformVector(name string, value dhvector.Vector) int {
	location := dhwrapper.GetUniformLocation(p.program_id, gl.Str(name))
	if location < 0 {
		log.Printf("debug: (%v) Invalid uniform name. name=%v", p.program_name, name)
		return -1
	}

	dhwrapper.Uniform3f(location, value.X, value.Y, value.Z)

	return 0
}
func (p *ShaderProgram) SetUniformColorU8(name string,value dhcoloru8.ColorU8) int{
	location := dhwrapper.GetUniformLocation(p.program_id, gl.Str(name))
	if location < 0 {
		log.Printf("debug: (%v) Invalid uniform name. name=%v", p.program_name, name)
		return -1
	}

	dhwrapper.Uniform4f(location, value.R, value.G, value.B,value.A)

	return 0
}
func(p *ShaderProgram) SetUniformMatrix(name string,transpose bool,value dhmatrix.Matrix) int{
	location := dhwrapper.GetUniformLocation(p.program_id, gl.Str(name))
	if location < 0 {
		log.Printf("debug: (%v) Invalid uniform name. name=%v", p.program_name, name)
		return -1
	}

	dhwrapper.UniformMatrix4fv(location,1,transpose,&value.M)

	return 0
}
