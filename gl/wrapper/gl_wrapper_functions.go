package wrapper

import (
	"log"
	"unsafe"

	"github.com/go-gl/gl/all-core/gl"
)

func ActiveTexture(arg0 uint32) {
	gl.ActiveTexture(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func AttachShader(arg0 uint32, arg1 uint32) {
	gl.AttachShader(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func BindBuffer(arg0 uint32, arg1 uint32) {
	gl.BindBuffer(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func BindFramebuffer(arg0 uint32, arg1 uint32) {
	gl.BindFramebuffer(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func BindRenderbuffer(arg0 uint32, arg1 uint32) {
	gl.BindRenderbuffer(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func BindSampler(arg0 uint32, arg1 uint32) {
	gl.BindSampler(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func BindTexture(arg0 uint32, arg1 uint32) {
	gl.BindTexture(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func BindVertexArray(arg0 uint32) {
	gl.BindVertexArray(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func BlendFunc(arg0 uint32, arg1 uint32) {
	gl.BlendFunc(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func BufferData(arg0 uint32, arg1 int, arg2 unsafe.Pointer, arg3 uint32) {
	gl.BufferData(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func CheckFramebufferStatus(arg0 uint32) uint32 {
	ret := gl.CheckFramebufferStatus(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}

	return ret
}
func Clear(arg0 uint32) {
	gl.Clear(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func ClearColor(arg0 float32, arg1 float32, arg2 float32, arg3 float32) {
	gl.ClearColor(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func ClearDepth(arg0 float64) {
	gl.ClearDepth(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func ClearDepthf(arg0 float32) {
	gl.ClearDepthf(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func ClearStencil(arg0 int32) {
	gl.ClearStencil(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func ColorMask(arg0 bool, arg1 bool, arg2 bool, arg3 bool) {
	gl.ColorMask(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func CompileShader(arg0 uint32) {
	gl.CompileShader(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func CopyTexImage2D(arg0 uint32, arg1 int32, arg2 uint32, arg3 int32, arg4 int32, arg5 int32, arg6 int32, arg7 int32) {
	gl.CopyTexImage2D(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func CreateProgram() uint32 {
	ret := gl.CreateProgram()

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}

	return ret
}
func CreateShader(arg0 uint32) uint32 {
	ret := gl.CreateShader(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}

	return ret
}
func CullFace(arg0 uint32) {
	gl.CullFace(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func DeleteBuffers(arg0 int32, arg1 *uint32) {
	gl.DeleteBuffers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func DeleteFramebuffers(arg0 int32, arg1 *uint32) {
	gl.DeleteFramebuffers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func DeleteRenderbuffers(arg0 int32, arg1 *uint32) {
	gl.DeleteRenderbuffers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func DeleteShader(arg0 uint32) {
	gl.DeleteShader(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func DeleteTextures(arg0 int32, arg1 *uint32) {
	gl.DeleteTextures(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func DeleteVertexArrays(arg0 int32, arg1 *uint32) {
	gl.DeleteVertexArrays(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func DepthFunc(arg0 uint32) {
	gl.DepthFunc(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func DepthMask(arg0 bool) {
	gl.DepthMask(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Disable(arg0 uint32) {
	gl.Disable(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func DisableVertexAttribArray(arg0 uint32) {
	gl.DisableVertexAttribArray(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func DrawArrays(arg0 uint32, arg1 int32, arg2 int32) {
	gl.DrawArrays(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func DrawBuffer(arg0 uint32) {
	gl.DrawBuffer(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func DrawBuffers(arg0 int32, arg1 *uint32) {
	gl.DrawBuffers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func DrawElements(arg0 uint32, arg1 int32, arg2 uint32, arg3 unsafe.Pointer) {
	gl.DrawElements(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Enable(arg0 uint32) {
	gl.Enable(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func EnableVertexAttribArray(arg0 uint32) {
	gl.EnableVertexAttribArray(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Flush() {
	gl.Flush()

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func FramebufferRenderbuffer(arg0 uint32, arg1 uint32, arg2 uint32, arg3 uint32) {
	gl.FramebufferRenderbuffer(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func FramebufferTexture(arg0 uint32, arg1 uint32, arg2 uint32, arg3 int32) {
	gl.FramebufferTexture(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func FramebufferTexture2D(arg0 uint32, arg1 uint32, arg2 uint32, arg3 uint32, arg4 int32) {
	gl.FramebufferTexture2D(arg0, arg1, arg2, arg3, arg4)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func GenBuffers(arg0 int32, arg1 *uint32) {
	gl.GenBuffers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func GenerateMipmap(arg0 uint32) {
	gl.GenerateMipmap(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func GenFramebuffers(arg0 int32, arg1 *uint32) {
	gl.GenFramebuffers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func GenRenderbuffers(arg0 int32, arg1 *uint32) {
	gl.GenRenderbuffers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func GenSamplers(arg0 int32, arg1 *uint32) {
	gl.GenSamplers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func GenTextures(arg0 int32, arg1 *uint32) {
	gl.GenTextures(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func GenVertexArrays(arg0 int32, arg1 *uint32) {
	gl.GenVertexArrays(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func GetIntegerv(arg0 uint32, arg1 *int32) {
	gl.GetIntegerv(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func GetProgramInfoLog(arg0 uint32, arg1 int32, arg2 *int32, arg3 *uint8) {
	gl.GetProgramInfoLog(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func GetProgramiv(arg0 uint32, arg1 uint32, arg2 *int32) {
	gl.GetProgramiv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func GetShaderInfoLog(arg0 uint32, arg1 int32, arg2 *int32, arg3 *uint8) {
	gl.GetShaderInfoLog(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func GetShaderiv(arg0 uint32, arg1 uint32, arg2 *int32) {
	gl.GetShaderiv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func GetTexImage(arg0 uint32, arg1 int32, arg2 uint32, arg3 uint32, arg4 unsafe.Pointer) {
	gl.GetTexImage(arg0, arg1, arg2, arg3, arg4)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func GetUniformLocation(arg0 uint32, arg1 *uint8) int32 {
	ret := gl.GetUniformLocation(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}

	return ret
}
func LinkProgram(arg0 uint32) {
	gl.LinkProgram(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func MapBuffer(arg0 uint32, arg1 uint32) unsafe.Pointer {
	ret := gl.MapBuffer(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}

	return ret
}
func ReadBuffer(arg0 uint32) {
	gl.ReadBuffer(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func ReadPixels(arg0 int32, arg1 int32, arg2 int32, arg3 int32, arg4 uint32, arg5 uint32, arg6 unsafe.Pointer) {
	gl.ReadPixels(arg0, arg1, arg2, arg3, arg4, arg5, arg6)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func RenderbufferStorage(arg0 uint32, arg1 uint32, arg2 int32, arg3 int32) {
	gl.RenderbufferStorage(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func SamplerParameteri(arg0 uint32, arg1 uint32, arg2 int32) {
	gl.SamplerParameteri(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func ShaderSource(arg0 uint32, arg1 int32, arg2 **uint8, arg3 *int32) {
	gl.ShaderSource(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func StencilFunc(arg0 uint32, arg1 int32, arg2 uint32) {
	gl.StencilFunc(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func StencilOp(arg0 uint32, arg1 uint32, arg2 uint32) {
	gl.StencilOp(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func TexImage2D(arg0 uint32, arg1 int32, arg2 int32, arg3 int32, arg4 int32, arg5 int32, arg6 uint32, arg7 uint32, arg8 unsafe.Pointer) {
	gl.TexImage2D(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func TexParameterf(arg0 uint32, arg1 uint32, arg2 float32) {
	gl.TexParameterf(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func TexParameterfv(arg0 uint32, arg1 uint32, arg2 *float32) {
	gl.TexParameterfv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func TexParameteri(arg0 uint32, arg1 uint32, arg2 int32) {
	gl.TexParameteri(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform1f(arg0 int32, arg1 float32) {
	gl.Uniform1f(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform1fv(arg0 int32, arg1 int32, arg2 *float32) {
	gl.Uniform1fv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform1i(arg0 int32, arg1 int32) {
	gl.Uniform1i(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform1iv(arg0 int32, arg1 int32, arg2 *int32) {
	gl.Uniform1iv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform2f(arg0 int32, arg1 float32, arg2 float32) {
	gl.Uniform2f(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform2fv(arg0 int32, arg1 int32, arg2 *float32) {
	gl.Uniform2fv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform2i(arg0 int32, arg1 int32, arg2 int32) {
	gl.Uniform2i(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform2iv(arg0 int32, arg1 int32, arg2 *int32) {
	gl.Uniform2iv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform3f(arg0 int32, arg1 float32, arg2 float32, arg3 float32) {
	gl.Uniform3f(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform3fv(arg0 int32, arg1 int32, arg2 *float32) {
	gl.Uniform3fv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform3i(arg0 int32, arg1 int32, arg2 int32, arg3 int32) {
	gl.Uniform3i(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform3iv(arg0 int32, arg1 int32, arg2 *int32) {
	gl.Uniform3iv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform4f(arg0 int32, arg1 float32, arg2 float32, arg3 float32, arg4 float32) {
	gl.Uniform4f(arg0, arg1, arg2, arg3, arg4)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform4fv(arg0 int32, arg1 int32, arg2 *float32) {
	gl.Uniform4fv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform4i(arg0 int32, arg1 int32, arg2 int32, arg3 int32, arg4 int32) {
	gl.Uniform4i(arg0, arg1, arg2, arg3, arg4)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Uniform4iv(arg0 int32, arg1 int32, arg2 *int32) {
	gl.Uniform4iv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func UniformMatrix2fv(arg0 int32, arg1 int32, arg2 bool, arg3 *float32) {
	gl.UniformMatrix2fv(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func UniformMatrix3fv(arg0 int32, arg1 int32, arg2 bool, arg3 *float32) {
	gl.UniformMatrix3fv(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func UniformMatrix4fv(arg0 int32, arg1 int32, arg2 bool, arg3 *float32) {
	gl.UniformMatrix4fv(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func UnmapBuffer(arg0 uint32) bool {
	ret := gl.UnmapBuffer(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}

	return ret
}
func UseProgram(arg0 uint32) {
	gl.UseProgram(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func VertexAttribPointer(arg0 uint32, arg1 int32, arg2 uint32, arg3 bool, arg4 int32, arg5 unsafe.Pointer) {
	gl.VertexAttribPointer(arg0, arg1, arg2, arg3, arg4, arg5)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
func Viewport(arg0 int32, arg1 int32, arg2 int32, arg3 int32) {
	gl.Viewport(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace: code=%v", code)
	}
}
