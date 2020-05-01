package texture

import (
	"log"
	"unsafe"

	"github.com/dabasan/goglf/gl/shader"

	"github.com/dabasan/goglf/gl/common"

	"github.com/dabasan/goglf/gl/wrapper"
	"github.com/go-gl/gl/all-core/gl"
)

type Texture struct {
	object uint32
	width  int32
	height int32
}

var textures_map map[int]*Texture
var count int

var default_texture_handle int

var generate_mipmap_flag bool

var window_width int
var window_height int

func init() {
	textures_map = make(map[int]*Texture)
	count = 0

	generate_mipmap_flag = true

	window_width = common.WINDOW_DEFAULT_WIDTH
	window_height = common.WINDOW_DEFAULT_HEIGHT
}

func NewTexture() *Texture {
	t := new(Texture)

	var texture_object uint32
	wrapper.GenTextures(1, &texture_object)
	wrapper.BindTexture(gl.TEXTURE_2D, texture_object)
	wrapper.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	wrapper.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	if generate_mipmap_flag == true {
		wrapper.GenerateMipmap(gl.TEXTURE_2D)
		wrapper.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST_MIPMAP_NEAREST)
	} else {
		wrapper.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	}
	wrapper.BindTexture(gl.TEXTURE_2D, 0)

	t.object = texture_object

	return t
}

func Initialize() {
	default_texture_handle = LoadTexture("./Data/Texture/white.bmp")

	log.Printf("info: TextureMgr initialized.")
}

func LoadTexture(filename string) int {
	rgba, err := loadRGBAFromImage(filename)
	if err != nil {
		log.Printf("error: Failed to load a texture. filename=%v", filename)
		return -1
	}

	texture := NewTexture()
	texture.width = int32(rgba.Rect.Size().X)
	texture.height = int32(rgba.Rect.Size().Y)

	wrapper.BindTexture(gl.TEXTURE_2D, texture.object)
	wrapper.TexImage2D(gl.TEXTURE_2D, 0, gl.SRGB_ALPHA, texture.width, texture.height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
	wrapper.BindTexture(gl.TEXTURE_2D, 0)

	texture_handle := count
	textures_map[texture_handle] = texture
	count++

	return texture_handle
}
func DeleteTexture(texture_handle int) int {
	texture, ok := textures_map[texture_handle]
	if !ok {
		log.Printf("warn: No such texture. texture_handle=%v", texture_handle)
		return -1
	}

	texture_object := texture.object
	wrapper.DeleteTextures(1, &texture_object)

	delete(textures_map, texture_handle)

	return 0
}
func DeleteAllTextures() {
	for _, texture := range textures_map {
		texture_object := texture.object
		wrapper.DeleteTextures(1, &texture_object)
	}

	textures_map = make(map[int]*Texture)
}

func FlipTexture(texture_handle int, flip_vertically bool, flip_horizontally bool) int {
	texture, ok := textures_map[texture_handle]
	if !ok {
		log.Printf("warn: No such texture. texture_handle=%v", texture_handle)
		return -1
	}

	data := make([]uint8, texture.width*texture.height*4)

	wrapper.BindTexture(gl.TEXTURE_2D, texture.object)
	wrapper.GetTexImage(gl.TEXTURE_2D, 0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&data[0]))
	wrapper.BindTexture(gl.TEXTURE_2D, 0)

	data_r := make([]uint8, texture.width*texture.height)
	data_g := make([]uint8, texture.width*texture.height)
	data_b := make([]uint8, texture.width*texture.height)
	data_a := make([]uint8, texture.width*texture.height)

	pos := 0
	bound := int(texture.width * texture.height * 4)
	for i := 0; i < bound; i += 4 {
		data_r[pos] = data[i]
		data_g[pos] = data[i+1]
		data_b[pos] = data[i+2]
		data_a[pos] = data[i+3]
		pos++
	}

	flipped_data := make([]uint8, texture.width*texture.height*4)

	pos = 0
	if flip_vertically == true && flip_horizontally == true {
		for y := texture.height - 1; y >= 0; y-- {
			for x := texture.width - 1; x >= 0; x-- {
				flipped_data[pos] = data_r[y*texture.width+x]
				flipped_data[pos+1] = data_g[y*texture.width+x]
				flipped_data[pos+2] = data_b[y*texture.width+x]
				flipped_data[pos+3] = data_a[y*texture.width+x]
				pos += 4
			}
		}
	} else if flip_vertically == true {
		for y := texture.height - 1; y >= 0; y-- {
			for x := int32(0); x < texture.width; x++ {
				flipped_data[pos] = data_r[y*texture.width+x]
				flipped_data[pos+1] = data_g[y*texture.width+x]
				flipped_data[pos+2] = data_b[y*texture.width+x]
				flipped_data[pos+3] = data_a[y*texture.width+x]
				pos += 4
			}
		}
	} else if flip_horizontally == true {
		for y := int32(0); y < texture.height; y++ {
			for x := texture.width - 1; x >= 0; x-- {
				flipped_data[pos] = data_r[y*texture.width+x]
				flipped_data[pos+1] = data_g[y*texture.width+x]
				flipped_data[pos+2] = data_b[y*texture.width+x]
				flipped_data[pos+3] = data_a[y*texture.width+x]
				pos += 4
			}
		}
	} else {
		return 0
	}

	wrapper.BindTexture(gl.TEXTURE_2D, texture.object)
	wrapper.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA,
		texture.width, texture.height, 0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&flipped_data[0]))
	wrapper.BindTexture(gl.TEXTURE_2D, 0)

	return 0
}

func GetTextureWidth(texture_handle int) int32 {
	texture, ok := textures_map[texture_handle]
	if !ok {
		log.Printf("warn: No such texture. texture_handle=%v", texture_handle)
		return -1
	}

	return texture.width
}
func GetTextureHeight(texture_handle int) int32 {
	texture, ok := textures_map[texture_handle]
	if !ok {
		log.Printf("warn: No such texture. texture_handle=%v", texture_handle)
		return -1
	}

	return texture.height
}

func AssociateTexture(texture_object uint32, texture_width int32, texture_height int32) int {
	var texture Texture
	texture.object = texture_object
	texture.width = texture_width
	texture.height = texture_height

	texture_handle := count
	textures_map[texture_handle] = &texture
	count++

	return texture_handle
}

func GetTextureImage(texture_handle int) ([]uint8, int) {
	texture, ok := textures_map[texture_handle]
	if !ok {
		log.Printf("warn: No such texture. texture_handle=%v", texture_handle)
		return nil, -1
	}

	data := make([]uint8, texture.width*texture.height*4)

	wrapper.BindTexture(gl.TEXTURE_2D, texture.object)
	wrapper.GetTexImage(gl.TEXTURE_2D, 0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&data[0]))
	wrapper.BindTexture(gl.TEXTURE_2D, 0)

	return data, len(data)
}

func TextureExists(texture_handle int) bool {
	_, ok := textures_map[texture_handle]
	return ok
}

func SetGenerateMipmapFlag(flag bool) {
	generate_mipmap_flag = flag
}

func SetWindowSize(width int, height int) {
	window_width = width
	window_height = height
}

func BindTexture(texture_handle int) int {
	texture, ok := textures_map[texture_handle]
	if !ok {
		texture, _ = textures_map[default_texture_handle]
	}

	wrapper.BindTexture(gl.TEXTURE_2D, texture.object)

	return 0
}
func UnbindTexture() {
	wrapper.BindTexture(gl.TEXTURE_2D, 0)
}

func DrawTexture(
	texture_handle int, x int, y int, width int, height int,
	bottom_left_u float32, bottom_left_v float32,
	bottom_right_u float32, bottom_right_v float32,
	top_right_u float32, top_right_v float32,
	top_left_u float32, top_left_v float32) int {
	texture, ok := textures_map[texture_handle]
	if !ok {
		texture, _ = textures_map[default_texture_handle]
	}

	indices := make([]int32, 6)
	pos_buffer := make([]float32, 8)
	uv_buffer := make([]float32, 8)

	indices[0] = 0
	indices[1] = 1
	indices[2] = 2
	indices[3] = 2
	indices[4] = 3
	indices[5] = 0

	normalized_x := 2.0*float32(x)/float32(window_width) - 1.0
	normalized_y := 2.0*float32(y)/float32(window_height) - 1.0
	normalized_width := float32(width) / float32(window_width) * 2.0
	normalized_height := float32(height) / float32(window_height) * 2.0

	//Bottom left
	pos_buffer[0] = normalized_x
	pos_buffer[1] = normalized_y
	uv_buffer[0] = bottom_left_u
	uv_buffer[1] = bottom_left_v
	//Bottom right
	pos_buffer[2] = normalized_x + normalized_width
	pos_buffer[3] = normalized_y
	uv_buffer[2] = bottom_right_u
	uv_buffer[3] = bottom_right_v
	//Top right
	pos_buffer[4] = normalized_x + normalized_width
	pos_buffer[5] = normalized_y + normalized_height
	uv_buffer[4] = top_right_u
	uv_buffer[5] = top_right_v
	//Top left
	pos_buffer[6] = normalized_x
	pos_buffer[7] = normalized_y + normalized_height
	uv_buffer[6] = top_left_u
	uv_buffer[7] = top_left_v

	shader.UseProgram("texture_drawer")

	var indices_vbo uint32
	var pos_vbo uint32
	var uv_vbo uint32
	var vao uint32

	wrapper.GenBuffers(1, &indices_vbo)
	wrapper.GenBuffers(1, &pos_vbo)
	wrapper.GenBuffers(1, &uv_vbo)
	wrapper.GenVertexArrays(1, &vao)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, pos_vbo)
	wrapper.BufferData(gl.ARRAY_BUFFER,
		wrapper.SIZEOF_FLOAT*len(pos_buffer), unsafe.Pointer(&pos_buffer[0]), gl.STATIC_DRAW)
	wrapper.BindBuffer(gl.ARRAY_BUFFER, uv_vbo)
	wrapper.BufferData(gl.ARRAY_BUFFER,
		wrapper.SIZEOF_FLOAT*len(uv_buffer), unsafe.Pointer(&uv_buffer[0]), gl.STATIC_DRAW)

	wrapper.BindVertexArray(vao)

	wrapper.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indices_vbo)
	wrapper.BufferData(gl.ELEMENT_ARRAY_BUFFER,
		wrapper.SIZEOF_INT*len(indices), unsafe.Pointer(&indices[0]), gl.STATIC_DRAW)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, pos_vbo)
	wrapper.EnableVertexAttribArray(0)
	wrapper.VertexAttribPointer(0, 2, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*2, nil)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, uv_vbo)
	wrapper.EnableVertexAttribArray(1)
	wrapper.VertexAttribPointer(1, 2, gl.FLOAT, false, wrapper.SIZEOF_FLOAT*2, nil)

	wrapper.BindBuffer(gl.ARRAY_BUFFER, 0)

	wrapper.BindTexture(gl.TEXTURE_2D, texture.object)
	wrapper.Enable(gl.BLEND)
	wrapper.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
	wrapper.Disable(gl.BLEND)
	wrapper.BindTexture(gl.TEXTURE_2D, 0)

	wrapper.BindVertexArray(0)

	wrapper.DeleteBuffers(1, &indices_vbo)
	wrapper.DeleteBuffers(1, &pos_vbo)
	wrapper.DeleteBuffers(1, &uv_vbo)
	wrapper.DeleteVertexArrays(1, &vao)

	return 0
}
func DrawTexture_Simple(texture_handle int, x int, y int, width int, height int) int {
	ret := DrawTexture(texture_handle, x, y, width, height, 0.0, 0.0, 1.0, 0.0, 1.0, 1.0, 0.0, 1.0)
	return ret
}
