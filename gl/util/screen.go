package util

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"unsafe"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"

	"github.com/dabasan/goglf/gl/shader"
	"github.com/dabasan/goglf/gl/texture"
	"github.com/dabasan/goglf/gl/transferrer"
	"github.com/dabasan/goglf/gl/wrapper"
	"github.com/go-gl/gl/all-core/gl"
)

type Screen struct {
	fbo_id          uint32
	renderbuffer_id uint32
	texture_id      uint32

	screen_width  int32
	screen_height int32

	program     *shader.ShaderProgram
	transferrer *transferrer.FullscreenQuadTransferrerWithUV

	texture_handle int
}

func NewScreen(width int, height int) *Screen {
	s := new(Screen)

	s.screen_width = int32(width)
	s.screen_height = int32(height)

	s.setupRenderbuffer()
	s.setupTexture()
	s.setupFramebuffer()

	s.program, _ = shader.NewShaderProgram("texture_drawer")
	s.transferrer = transferrer.NewFullscreenQuadTransferrerWithUV(true)

	s.texture_handle = -1

	return s
}
func (s *Screen) setupRenderbuffer() {
	wrapper.GenRenderbuffers(1, &s.renderbuffer_id)
	wrapper.BindRenderbuffer(gl.RENDERBUFFER, s.renderbuffer_id)
	wrapper.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH_COMPONENT, s.screen_width, s.screen_height)
	wrapper.BindRenderbuffer(gl.RENDERBUFFER, 0)
}
func (s *Screen) setupTexture() {
	wrapper.GenTextures(1, &s.texture_id)
	wrapper.BindTexture(gl.TEXTURE_2D, s.texture_id)
	wrapper.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA,
		s.screen_width, s.screen_height, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)
	wrapper.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	wrapper.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	wrapper.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	wrapper.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	wrapper.BindTexture(gl.TEXTURE_2D, 0)
}
func (s *Screen) setupFramebuffer() {
	wrapper.GenFramebuffers(1, &s.fbo_id)

	wrapper.BindFramebuffer(gl.FRAMEBUFFER, s.fbo_id)

	wrapper.FramebufferTexture2D(gl.FRAMEBUFFER,
		gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, s.texture_id, 0)
	wrapper.FramebufferRenderbuffer(gl.FRAMEBUFFER,
		gl.DEPTH_ATTACHMENT, gl.RENDERBUFFER, s.renderbuffer_id)

	draw_buffers := [...]uint32{gl.COLOR_ATTACHMENT0}
	wrapper.DrawBuffers(1, &draw_buffers[0])

	status := wrapper.CheckFramebufferStatus(gl.FRAMEBUFFER)
	if status != gl.FRAMEBUFFER_COMPLETE {
		log.Printf("warn: Incomplete framebuffer. status=%v", status)
	}

	wrapper.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (s *Screen) GetScreenSize() (int, int) {
	return int(s.screen_width), int(s.screen_height)
}

func (s *Screen) Enable() {
	wrapper.BindFramebuffer(gl.FRAMEBUFFER, s.fbo_id)
	wrapper.Viewport(0, 0, s.screen_width, s.screen_height)
}
func (s *Screen) Disable() {
	wrapper.BindFramebuffer(gl.FRAMEBUFFER, 0)
}
func (s *Screen) Clear() {
	gl.Clear(gl.DEPTH_BUFFER_BIT | gl.COLOR_BUFFER_BIT)
}

func (s *Screen) Associate() int {
	texture_handle := texture.AssociateTexture(s.texture_id, s.screen_width, s.screen_height)
	return texture_handle
}

func (s *Screen) Draw() {
	s.program.Enable()

	wrapper.ActiveTexture(gl.TEXTURE0)
	wrapper.BindTexture(gl.TEXTURE_2D, s.texture_id)
	s.program.SetUniform1i("texture_sampler", 0)

	s.transferrer.Transfer()
}

func (s *Screen) DrawRect(x int, y int, width int, height int) {
	wrapper.Viewport(int32(x), int32(y), int32(width), int32(height))
	s.Draw()
}

func (s *Screen) BindScreenTexture() {
	wrapper.BindTexture(gl.TEXTURE_2D, s.texture_id)
}
func (s *Screen) UnbindScreenTexture() {
	wrapper.BindTexture(gl.TEXTURE_2D, 0)
}

//TakeScreenshot takes a screenshot of this screen.
//You have to specify the format such as "jpg" or "png" by the argument.
//Currently supported formats are jpg, png, bmp and tiff.
//Call example: TakeScreenshot("screenshot.png","png")
func (s *Screen) TakeScreenshot(filename string, format string) error {
	log.Printf("info: Taking a screenshot. filename=%v format=%v", filename, format)

	width := int(s.screen_width)
	height := int(s.screen_height)

	data := make([]byte, width*height*4)

	wrapper.BindTexture(gl.TEXTURE_2D, s.texture_id)
	wrapper.GetTexImage(gl.TEXTURE_2D, 0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&data[0]))
	wrapper.BindTexture(gl.TEXTURE_2D, 0)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	pos := 0
	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			r := data[pos]
			g := data[pos+1]
			b := data[pos+2]
			a := data[pos+3]

			img.Set(x, y, color.RGBA{R: r, G: g, B: b, A: a})

			pos += 4
		}
	}

	var err error
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	switch format {
	case "jpg", "jpeg":
		err = jpeg.Encode(file, img, nil)
	case "png":
		err = png.Encode(file, img)
	case "bmp":
		err = bmp.Encode(file, img)
	case "tiff":
		err = tiff.Encode(file, img, nil)
	default:
		return fmt.Errorf("Unsupported file format. format=%v", format)
	}

	if err != nil {
		return err
	}

	return nil
}
