package dhwrapper

import (
	"log"
	"strconv"
	"unsafe"

	"github.com/go-gl/gl/all-core/gl"
)

func ActiveTexture(arg0 uint32) {
	gl.ActiveTexture(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[ActiveTexture] code:" + strconv.Itoa(int(code)))
	}
}
func AttachShader(arg0 uint32, arg1 uint32) {
	gl.AttachShader(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[AttachShader] code:" + strconv.Itoa(int(code)))
	}
}
func BindBuffer(arg0 uint32, arg1 uint32) {
	gl.BindBuffer(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[BindBuffer] code:" + strconv.Itoa(int(code)))
	}
}
func BindFramebuffer(arg0 uint32, arg1 uint32) {
	gl.BindFramebuffer(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[BindFramebuffer] code:" + strconv.Itoa(int(code)))
	}
}
func BindRenderbuffer(arg0 uint32, arg1 uint32) {
	gl.BindRenderbuffer(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[BindRenderbuffer] code:" + strconv.Itoa(int(code)))
	}
}
func BindSampler(arg0 uint32, arg1 uint32) {
	gl.BindSampler(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[BindSampler] code:" + strconv.Itoa(int(code)))
	}
}
func BindTexture(arg0 uint32, arg1 uint32) {
	gl.BindTexture(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[BindTexture] code:" + strconv.Itoa(int(code)))
	}
}
func BindVertexArray(arg0 uint32) {
	gl.BindVertexArray(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[BindVertexArray] code:" + strconv.Itoa(int(code)))
	}
}
func BlendFunc(arg0 uint32, arg1 uint32) {
	gl.BlendFunc(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[BlendFunc] code:" + strconv.Itoa(int(code)))
	}
}
func BufferData(arg0 uint32, arg1 int, arg2 unsafe.Pointer, arg3 uint32) {
	gl.BufferData(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[BufferData] code:" + strconv.Itoa(int(code)))
	}
}
func CheckFramebufferStatus(arg0 uint32) uint32 {
	ret := gl.CheckFramebufferStatus(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[CheckFramebufferStatus] code:" + strconv.Itoa(int(code)))
	}

	return ret
}
func Clear(arg0 uint32) {
	gl.Clear(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Clear] code:" + strconv.Itoa(int(code)))
	}
}
func ClearColor(arg0 float32, arg1 float32, arg2 float32, arg3 float32) {
	gl.ClearColor(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[ClearColor] code:" + strconv.Itoa(int(code)))
	}
}
func ClearDepth(arg0 float64) {
	gl.ClearDepth(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[ClearDepth] code:" + strconv.Itoa(int(code)))
	}
}
func ClearDepthf(arg0 float32) {
	gl.ClearDepthf(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[ClearDepthf] code:" + strconv.Itoa(int(code)))
	}
}
func ClearStencil(arg0 int32) {
	gl.ClearStencil(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[ClearStencil] code:" + strconv.Itoa(int(code)))
	}
}
func ColorMask(arg0 bool, arg1 bool, arg2 bool, arg3 bool) {
	gl.ColorMask(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[ColorMask] code:" + strconv.Itoa(int(code)))
	}
}
func CompileShader(arg0 uint32) {
	gl.CompileShader(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[CompileShader] code:" + strconv.Itoa(int(code)))
	}
}
func CopyTexImage2D(arg0 uint32, arg1 int32, arg2 uint32, arg3 int32, arg4 int32, arg5 int32, arg6 int32, arg7 int32) {
	gl.CopyTexImage2D(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[CopyTexImage2D] code:" + strconv.Itoa(int(code)))
	}
}
func CreateProgram() uint32 {
	ret := gl.CreateProgram()

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[CreateProgram] code:" + strconv.Itoa(int(code)))
	}

	return ret
}
func CreateShader(arg0 uint32) uint32 {
	ret := gl.CreateShader(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[CreateShader] code:" + strconv.Itoa(int(code)))
	}

	return ret
}
func CullFace(arg0 uint32) {
	gl.CullFace(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[CullFace] code:" + strconv.Itoa(int(code)))
	}
}
func DeleteBuffers(arg0 int32, arg1 *uint32) {
	gl.DeleteBuffers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[DeleteBuffers] code:" + strconv.Itoa(int(code)))
	}
}
func DeleteFramebuffers(arg0 int32, arg1 *uint32) {
	gl.DeleteFramebuffers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[DeleteFramebuffers] code:" + strconv.Itoa(int(code)))
	}
}
func DeleteRenderbuffers(arg0 int32, arg1 *uint32) {
	gl.DeleteRenderbuffers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[DeleteRenderbuffers] code:" + strconv.Itoa(int(code)))
	}
}
func DeleteShader(arg0 uint32) {
	gl.DeleteShader(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[DeleteShader] code:" + strconv.Itoa(int(code)))
	}
}
func DeleteTextures(arg0 int32, arg1 *uint32) {
	gl.DeleteTextures(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[DeleteTextures] code:" + strconv.Itoa(int(code)))
	}
}
func DeleteVertexArrays(arg0 int32, arg1 *uint32) {
	gl.DeleteVertexArrays(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[DeleteVertexArrays] code:" + strconv.Itoa(int(code)))
	}
}
func DepthFunc(arg0 uint32) {
	gl.DepthFunc(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[DepthFunc] code:" + strconv.Itoa(int(code)))
	}
}
func DepthMask(arg0 bool) {
	gl.DepthMask(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[DepthMask] code:" + strconv.Itoa(int(code)))
	}
}
func Disable(arg0 uint32) {
	gl.Disable(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Disable] code:" + strconv.Itoa(int(code)))
	}
}
func DisableVertexAttribArray(arg0 uint32) {
	gl.DisableVertexAttribArray(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[DisableVertexAttribArray] code:" + strconv.Itoa(int(code)))
	}
}
func DrawArrays(arg0 uint32, arg1 int32, arg2 int32) {
	gl.DrawArrays(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[DrawArrays] code:" + strconv.Itoa(int(code)))
	}
}
func DrawBuffer(arg0 uint32) {
	gl.DrawBuffer(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[DrawBuffer] code:" + strconv.Itoa(int(code)))
	}
}
func DrawBuffers(arg0 int32, arg1 *uint32) {
	gl.DrawBuffers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[DrawBuffers] code:" + strconv.Itoa(int(code)))
	}
}
func DrawElements(arg0 uint32, arg1 int32, arg2 uint32, arg3 unsafe.Pointer) {
	gl.DrawElements(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[DrawElements] code:" + strconv.Itoa(int(code)))
	}
}
func Enable(arg0 uint32) {
	gl.Enable(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Enable] code:" + strconv.Itoa(int(code)))
	}
}
func EnableVertexAttribArray(arg0 uint32) {
	gl.EnableVertexAttribArray(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[EnableVertexAttribArray] code:" + strconv.Itoa(int(code)))
	}
}
func Flush() {
	gl.Flush()

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Flush] code:" + strconv.Itoa(int(code)))
	}
}
func FramebufferRenderbuffer(arg0 uint32, arg1 uint32, arg2 uint32, arg3 uint32) {
	gl.FramebufferRenderbuffer(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[FramebufferRenderbuffer] code:" + strconv.Itoa(int(code)))
	}
}
func FramebufferTexture(arg0 uint32, arg1 uint32, arg2 uint32, arg3 int32) {
	gl.FramebufferTexture(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[FramebufferTexture] code:" + strconv.Itoa(int(code)))
	}
}
func FramebufferTexture2D(arg0 uint32, arg1 uint32, arg2 uint32, arg3 uint32, arg4 int32) {
	gl.FramebufferTexture2D(arg0, arg1, arg2, arg3, arg4)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[FramebufferTexture2D] code:" + strconv.Itoa(int(code)))
	}
}
func GenBuffers(arg0 int32, arg1 *uint32) {
	gl.GenBuffers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[GenBuffers] code:" + strconv.Itoa(int(code)))
	}
}
func GenerateMipmap(arg0 uint32) {
	gl.GenerateMipmap(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[GenerateMipmap] code:" + strconv.Itoa(int(code)))
	}
}
func GenFramebuffers(arg0 int32, arg1 *uint32) {
	gl.GenFramebuffers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[GenFramebuffers] code:" + strconv.Itoa(int(code)))
	}
}
func GenRenderbuffers(arg0 int32, arg1 *uint32) {
	gl.GenRenderbuffers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[GenRenderbuffers] code:" + strconv.Itoa(int(code)))
	}
}
func GenSamplers(arg0 int32, arg1 *uint32) {
	gl.GenSamplers(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[GenSamplers] code:" + strconv.Itoa(int(code)))
	}
}
func GenTextures(arg0 int32, arg1 *uint32) {
	gl.GenTextures(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[GenTextures] code:" + strconv.Itoa(int(code)))
	}
}
func GenVertexArrays(arg0 int32, arg1 *uint32) {
	gl.GenVertexArrays(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[GenVertexArrays] code:" + strconv.Itoa(int(code)))
	}
}
func GetIntegerv(arg0 uint32, arg1 *int32) {
	gl.GetIntegerv(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[GetIntegerv] code:" + strconv.Itoa(int(code)))
	}
}
func GetProgramInfoLog(arg0 uint32, arg1 int32, arg2 *int32, arg3 *uint8) {
	gl.GetProgramInfoLog(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[GetProgramInfoLog] code:" + strconv.Itoa(int(code)))
	}
}
func GetProgramiv(arg0 uint32, arg1 uint32, arg2 *int32) {
	gl.GetProgramiv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[GetProgramiv] code:" + strconv.Itoa(int(code)))
	}
}
func GetShaderInfoLog(arg0 uint32, arg1 int32, arg2 *int32, arg3 *uint8) {
	gl.GetShaderInfoLog(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[GetShaderInfoLog] code:" + strconv.Itoa(int(code)))
	}
}
func GetShaderiv(arg0 uint32, arg1 uint32, arg2 *int32) {
	gl.GetShaderiv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[GetShaderiv] code:" + strconv.Itoa(int(code)))
	}
}
func GetTexImage(arg0 uint32, arg1 int32, arg2 uint32, arg3 uint32, arg4 unsafe.Pointer) {
	gl.GetTexImage(arg0, arg1, arg2, arg3, arg4)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[GetTexImage] code:" + strconv.Itoa(int(code)))
	}
}
func GetUniformLocation(arg0 uint32, arg1 *uint8) {
	gl.GetUniformLocation(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[GetUniformLocation] code:" + strconv.Itoa(int(code)))
	}
}
func LinkProgram(arg0 uint32) {
	gl.LinkProgram(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[LinkProgram] code:" + strconv.Itoa(int(code)))
	}
}
func MapBuffer(arg0 uint32, arg1 uint32) unsafe.Pointer {
	ret := gl.MapBuffer(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[MapBuffer] code:" + strconv.Itoa(int(code)))
	}

	return ret
}
func ReadBuffer(arg0 uint32) {
	gl.ReadBuffer(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[ReadBuffer] code:" + strconv.Itoa(int(code)))
	}
}
func ReadPixels(arg0 int32, arg1 int32, arg2 int32, arg3 int32, arg4 uint32, arg5 uint32, arg6 unsafe.Pointer) {
	gl.ReadPixels(arg0, arg1, arg2, arg3, arg4, arg5, arg6)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[ReadPixels] code:" + strconv.Itoa(int(code)))
	}
}
func RenderbufferStorage(arg0 uint32, arg1 uint32, arg2 int32, arg3 int32) {
	gl.RenderbufferStorage(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[RenderbufferStorage] code:" + strconv.Itoa(int(code)))
	}
}
func SamplerParameteri(arg0 uint32, arg1 uint32, arg2 int32) {
	gl.SamplerParameteri(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[SamplerParameteri] code:" + strconv.Itoa(int(code)))
	}
}
func ShaderSource(arg0 uint32, arg1 int32, arg2 **uint8, arg3 *int32) {
	gl.ShaderSource(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[ShaderSource] code:" + strconv.Itoa(int(code)))
	}
}
func StencilFunc(arg0 uint32, arg1 int32, arg2 uint32) {
	gl.StencilFunc(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[StencilFunc] code:" + strconv.Itoa(int(code)))
	}
}
func StencilOp(arg0 uint32, arg1 uint32, arg2 uint32) {
	gl.StencilOp(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[StencilOp] code:" + strconv.Itoa(int(code)))
	}
}
func TexImage2D(arg0 uint32, arg1 int32, arg2 int32, arg3 int32, arg4 int32, arg5 int32, arg6 uint32, arg7 uint32, arg8 unsafe.Pointer) {
	gl.TexImage2D(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[TexImage2D] code:" + strconv.Itoa(int(code)))
	}
}
func TexParameterf(arg0 uint32, arg1 uint32, arg2 float32) {
	gl.TexParameterf(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[TexParameterf] code:" + strconv.Itoa(int(code)))
	}
}
func TexParameterfv(arg0 uint32, arg1 uint32, arg2 *float32) {
	gl.TexParameterfv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[TexParameterfv] code:" + strconv.Itoa(int(code)))
	}
}
func TexParameteri(arg0 uint32, arg1 uint32, arg2 int32) {
	gl.TexParameteri(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[TexParameteri] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform1f(arg0 int32, arg1 float32) {
	gl.Uniform1f(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform1f] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform1fv(arg0 int32, arg1 int32, arg2 *float32) {
	gl.Uniform1fv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform1fv] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform1i(arg0 int32, arg1 int32) {
	gl.Uniform1i(arg0, arg1)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform1i] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform1iv(arg0 int32, arg1 int32, arg2 *int32) {
	gl.Uniform1iv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform1iv] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform2f(arg0 int32, arg1 float32, arg2 float32) {
	gl.Uniform2f(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform2f] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform2fv(arg0 int32, arg1 int32, arg2 *float32) {
	gl.Uniform2fv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform2fv] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform2i(arg0 int32, arg1 int32, arg2 int32) {
	gl.Uniform2i(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform2i] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform2iv(arg0 int32, arg1 int32, arg2 *int32) {
	gl.Uniform2iv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform2iv] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform3f(arg0 int32, arg1 float32, arg2 float32, arg3 float32) {
	gl.Uniform3f(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform3f] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform3fv(arg0 int32, arg1 int32, arg2 *float32) {
	gl.Uniform3fv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform3fv] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform3i(arg0 int32, arg1 int32, arg2 int32, arg3 int32) {
	gl.Uniform3i(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform3i] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform3iv(arg0 int32, arg1 int32, arg2 *int32) {
	gl.Uniform3iv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform3iv] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform4f(arg0 int32, arg1 float32, arg2 float32, arg3 float32, arg4 float32) {
	gl.Uniform4f(arg0, arg1, arg2, arg3, arg4)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform4f] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform4fv(arg0 int32, arg1 int32, arg2 *float32) {
	gl.Uniform4fv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform4fv] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform4i(arg0 int32, arg1 int32, arg2 int32, arg3 int32, arg4 int32) {
	gl.Uniform4i(arg0, arg1, arg2, arg3, arg4)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform4i] code:" + strconv.Itoa(int(code)))
	}
}
func Uniform4iv(arg0 int32, arg1 int32, arg2 *int32) {
	gl.Uniform4iv(arg0, arg1, arg2)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Uniform4iv] code:" + strconv.Itoa(int(code)))
	}
}
func UniformMatrix2fv(arg0 int32, arg1 int32, arg2 bool, arg3 *float32) {
	gl.UniformMatrix2fv(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[UniformMatrix2fv] code:" + strconv.Itoa(int(code)))
	}
}
func UniformMatrix3fv(arg0 int32, arg1 int32, arg2 bool, arg3 *float32) {
	gl.UniformMatrix3fv(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[UniformMatrix3fv] code:" + strconv.Itoa(int(code)))
	}
}
func UniformMatrix4fv(arg0 int32, arg1 int32, arg2 bool, arg3 *float32) {
	gl.UniformMatrix4fv(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[UniformMatrix4fv] code:" + strconv.Itoa(int(code)))
	}
}
func UnmapBuffer(arg0 uint32) bool {
	ret := gl.UnmapBuffer(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[UnmapBuffer] code:" + strconv.Itoa(int(code)))
	}

	return ret
}
func UseProgram(arg0 uint32) {
	gl.UseProgram(arg0)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[UseProgram] code:" + strconv.Itoa(int(code)))
	}
}
func VertexAttribPointer(arg0 uint32, arg1 int32, arg2 uint32, arg3 bool, arg4 int32, arg5 unsafe.Pointer) {
	gl.VertexAttribPointer(arg0, arg1, arg2, arg3, arg4, arg5)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[VertexAttribPointer] code:" + strconv.Itoa(int(code)))
	}
}
func Viewport(arg0 int32, arg1 int32, arg2 int32, arg3 int32) {
	gl.Viewport(arg0, arg1, arg2, arg3)

	code := gl.GetError()
	if code != gl.NO_ERROR {
		log.Printf("trace:[Viewport] code:" + strconv.Itoa(int(code)))
	}
}
