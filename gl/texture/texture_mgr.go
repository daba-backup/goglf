package texture

import (
	"log"
	"unsafe"

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

func init() {
	textures_map = make(map[int]*Texture)
	count = 0

	generate_mipmap_flag = true
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

	var texture_object uint32
	texture_width := int32(rgba.Rect.Size().X)
	texture_height := int32(rgba.Rect.Size().Y)

	wrapper.GenTextures(1, &texture_object)
	wrapper.BindTexture(gl.TEXTURE_2D, texture_object)
	wrapper.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	wrapper.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	wrapper.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA,
		texture_width, texture_height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
	if generate_mipmap_flag == true {
		wrapper.GenerateMipmap(gl.TEXTURE_2D)
		wrapper.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST_MIPMAP_NEAREST)
	} else {
		wrapper.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	}
	wrapper.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	wrapper.BindTexture(gl.TEXTURE_2D, 0)

	texture := new(Texture)
	texture.object = texture_object
	texture.width = texture_width
	texture.height = texture_height

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

func GetTextureSize(texture_handle int) (int32, int32) {
	texture, ok := textures_map[texture_handle]
	if !ok {
		log.Printf("warn: No such texture. texture_handle=%v", texture_handle)
		return -1, -1
	}

	return texture.width, texture.height
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
