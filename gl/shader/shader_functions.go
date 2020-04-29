package shader

import (
	"log"
	"strings"

	"github.com/dabasan/go-dhtool/file"
	"github.com/dabasan/goglf/gl/wrapper"
	"github.com/go-gl/gl/all-core/gl"
)

var program_ids_map map[string]uint32

func init() {
	program_ids_map = make(map[string]uint32)
}

func CreateProgram(program_name string, vertex_shader_filename string, fragment_shader_filename string) int {
	log.Printf("info: Start creating a program. program_name=%v", program_name)
	log.Printf("info: vertex_shader_filename=%v fragment_shader_filename=%v", vertex_shader_filename, fragment_shader_filename)

	vertex_shader_id := wrapper.CreateShader(gl.VERTEX_SHADER)
	fragment_shader_id := wrapper.CreateShader(gl.FRAGMENT_SHADER)

	//Load the code of shaders.
	vertex_shader_code, err := file.GetFileAllLines_Concat(vertex_shader_filename)
	if err != nil {
		log.Printf("error: %v", err.Error())
		return -1
	}
	fragment_shader_code, err := file.GetFileAllLines_Concat(fragment_shader_filename)
	if err != nil {
		log.Printf("error: %v", err.Error())
		return -1
	}

	vertex_shader_srcs, v_free_fn := gl.Strs(vertex_shader_code + "\x00")
	fragment_shader_srcs, f_free_fn := gl.Strs(fragment_shader_code + "\x00")
	defer v_free_fn()
	defer f_free_fn()

	var info_log_length int32
	var result int32
	var error_message string

	//Compile vertex shader.
	wrapper.ShaderSource(vertex_shader_id, 1, vertex_shader_srcs, nil)
	wrapper.CompileShader(vertex_shader_id)

	//Check vertex shader.
	wrapper.GetShaderiv(vertex_shader_id, gl.COMPILE_STATUS, &result)
	if result == gl.FALSE {
		wrapper.GetShaderiv(vertex_shader_id, gl.INFO_LOG_LENGTH, &info_log_length)

		error_message = strings.Repeat("\x00", int(info_log_length+1))
		wrapper.GetShaderInfoLog(vertex_shader_id, info_log_length, nil, gl.Str(error_message))

		log.Printf("error: Vertex shader compilation failed.")
		log.Printf("error: %v", error_message)

		return -1
	}

	//Compile fragment shader.
	wrapper.ShaderSource(fragment_shader_id, 1, fragment_shader_srcs, nil)
	wrapper.CompileShader(fragment_shader_id)

	//Check fragment shader.
	wrapper.GetShaderiv(fragment_shader_id, gl.COMPILE_STATUS, &result)
	if result == gl.FALSE {
		wrapper.GetShaderiv(fragment_shader_id, gl.INFO_LOG_LENGTH, &info_log_length)

		error_message = strings.Repeat("\x00", int(info_log_length+1))
		wrapper.GetShaderInfoLog(fragment_shader_id, info_log_length, nil, gl.Str(error_message))

		log.Printf("error: Fragment shader compilation failed.")
		log.Printf("error: %v", error_message)

		return -1
	}

	//Link program.
	program_id := wrapper.CreateProgram()
	wrapper.AttachShader(program_id, vertex_shader_id)
	wrapper.AttachShader(program_id, fragment_shader_id)

	wrapper.LinkProgram(program_id)

	//Check program.
	wrapper.GetProgramiv(program_id, gl.LINK_STATUS, &result)
	if result == gl.FALSE {
		wrapper.GetProgramiv(program_id, gl.INFO_LOG_LENGTH, &info_log_length)

		error_message = strings.Repeat("\x00", int(info_log_length+1))
		wrapper.GetProgramInfoLog(program_id, info_log_length, nil, gl.Str(error_message))

		log.Printf("error: Program link failed.")
		log.Printf("error: %v", error_message)

		return -1
	}

	wrapper.DeleteShader(vertex_shader_id)
	wrapper.DeleteShader(fragment_shader_id)

	program_ids_map[program_name] = program_id
	log.Printf("info: Successfully created a program. program_name=%v program_id=%v", program_name, program_id)

	return 0
}

func GetProgramID(program_name string) (uint32, bool) {
	val, ok := program_ids_map[program_name]
	return val, ok
}
func UseProgram(program_name string) int {
	val, ok := program_ids_map[program_name]
	if ok == false {
		log.Printf("trace: No such program. program_name=%v", program_name)
		return -1
	}

	wrapper.UseProgram(val)

	return 0
}
