package bd1

import (
	dhbyte "github.com/dabasan/go-dhtool/byte"
	"github.com/dabasan/go-dhtool/file"
	"github.com/dabasan/go-dhtool/filename"
)

type bd1Reader struct {
	texture_filenames_map map[int]string
	blocks                []*bd1Block
}

func newbd1Reader() *bd1Reader {
	p := new(bd1Reader)
	p.texture_filenames_map = make(map[int]string)
	p.blocks = make([]*bd1Block, 0)

	return p
}

func (p *bd1Reader) read(bd1_filename string) error {
	bin, err := file.GetFileAllBin(bd1_filename)
	if err != nil {
		return err
	}

	pos := 0

	//Texture filenames
	for i := 0; i < 10; i++ {
		var texture_filenames_buffer [31]byte
		var texture_filenames_temp string
		var first_null_pos int

		for j := 0; j < 31; j++ {
			texture_filenames_buffer[j] = bin[pos]
			pos++
		}
		texture_filenames_temp = string(texture_filenames_buffer[:])

		for j := 0; j < 30; j++ {
			if texture_filenames_temp[j] == '\x00' {
				first_null_pos = j
				break
			}
		}

		texture_filenames_temp = texture_filenames_temp[0:first_null_pos]
		texture_filenames_temp = filename.ReplaceWindowsDelimiterWithLinuxDelimiter(texture_filenames_temp)

		p.texture_filenames_map[i] = texture_filenames_temp
	}

	//Number of blocks
	block_num, err := dhbyte.GetUint16ValueFromBin_LE(bin, pos)
	if err != nil {
		return err
	}
	pos += 2

	i_block_num := int(block_num)
	p.blocks = make([]*bd1Block, i_block_num)

	//Blocks
	for i := 0; i < i_block_num; i++ {
		var block bd1Block
		var coordinate_temp float32
		var err error

		//Vertex positions
		for j := 0; j < 8; j++ {
			coordinate_temp, err = dhbyte.GetFloat32ValueFromBin_LE(bin, pos)
			if err != nil {
				return err
			}
			block.Vertex_positions[j].X = coordinate_temp
			pos += 4
		}
		for j := 0; j < 8; j++ {
			coordinate_temp, err = dhbyte.GetFloat32ValueFromBin_LE(bin, pos)
			if err != nil {
				return err
			}
			block.Vertex_positions[j].Y = coordinate_temp
			pos += 4
		}
		for j := 0; j < 8; j++ {
			coordinate_temp, err = dhbyte.GetFloat32ValueFromBin_LE(bin, pos)
			if err != nil {
				return err
			}
			block.Vertex_positions[j].Z = coordinate_temp
			pos += 4
		}

		//UV coordinates
		for j := 0; j < 24; j++ {
			coordinate_temp, err = dhbyte.GetFloat32ValueFromBin_LE(bin, pos)
			if err != nil {
				return err
			}
			block.Us[j] = coordinate_temp
			pos += 4
		}
		for j := 0; j < 24; j++ {
			coordinate_temp, err = dhbyte.GetFloat32ValueFromBin_LE(bin, pos)
			if err != nil {
				return err
			}
			block.Vs[j] = coordinate_temp
			pos += 4
		}

		//Texture IDs
		for j := 0; j < 6; j++ {
			texture_id := bin[pos]
			block.Texture_ids[j] = int(texture_id)
			pos += 4
		}

		//Enabled flag
		enabled_flag := bin[pos]
		if enabled_flag != 0 {
			block.Enabled_flag = true
		} else {
			block.Enabled_flag = false
		}
		pos += 4

		p.blocks[i] = &block
	}

	return nil
}

func (p *bd1Reader) getTextureFilenames() map[int]string {
	return p.texture_filenames_map
}
func (p *bd1Reader) getBlocks() []*bd1Block {
	return p.blocks
}
