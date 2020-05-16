package obj

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/dabasan/go-dh3dbasis/vector"

	"github.com/dabasan/go-dhtool/filename"

	"github.com/dabasan/go-dhtool/file"

	"github.com/dabasan/go-dh3dbasis/coloru8"
	"github.com/dabasan/goglf/gl/model/buffer"
	gdfobj "github.com/mokiat/go-data-front/decoder/obj"
)

type Material struct {
	Name string

	Ambient_color       coloru8.ColorU8
	Diffuse_color       coloru8.ColorU8
	Specular_color      coloru8.ColorU8
	Specular_exponent   float32
	Dissolve            float32
	Diffuse_texture_map string
}

func LoadOBJ(obj_filename string) ([]*buffer.BufferedVertices, error) {
	log.Printf("info: Start loading an OBJ file. obj_filename=%v", obj_filename)

	ret := make([]*buffer.BufferedVertices, 0)

	decoder := gdfobj.NewDecoder(gdfobj.DefaultLimits())
	file, err := os.Open(obj_filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	model, err := decoder.Decode(file)
	if err != nil {
		return nil, err
	}

	obj_directory := filename.GetFileDirectory(obj_filename)
	mtl_filename := obj_directory + "/" + model.MaterialLibraries[0]

	materials, err := LoadMtl(mtl_filename)
	if err != nil {
		return nil, err
	}

	for _, o := range model.Objects {
		for _, mesh := range o.Meshes {
			material := findMaterialByName(mesh.MaterialName, materials)
			if material == nil {
				log.Printf("warn: No such material. material_name=%v", mesh.MaterialName)
				continue
			}

			buffered_vertices := buffer.NewBufferedVertices()
			buffered_vertices.SetAmbientColor(material.Ambient_color)
			buffered_vertices.SetDiffuseColor(material.Diffuse_color)
			buffered_vertices.SetSpecularColor(material.Specular_color)
			buffered_vertices.SetSpecularExponent(material.Specular_exponent)
			buffered_vertices.SetDissolve(material.Dissolve)
			buffered_vertices.SetDiffuseTextureMap(material.Diffuse_texture_map)

			indices_buffer := make([]uint32, 0)
			pos_buffer := make([]float32, 0)
			norm_buffer := make([]float32, 0)
			uv_buffer := make([]float32, 0)

			var vertex_count uint32 = 0

			for _, face := range mesh.Faces {
				vertex_num := len(face.References)
				if vertex_num != 3 && vertex_num != 4 {
					continue
				}

				if vertex_num == 3 {
					indices_buffer = append(indices_buffer, vertex_count)
					indices_buffer = append(indices_buffer, vertex_count+1)
					indices_buffer = append(indices_buffer, vertex_count+2)

					vertex_count += 3
				} else if vertex_num == 4 {
					indices_buffer = append(indices_buffer, vertex_count)
					indices_buffer = append(indices_buffer, vertex_count+1)
					indices_buffer = append(indices_buffer, vertex_count+2)

					indices_buffer = append(indices_buffer, vertex_count+2)
					indices_buffer = append(indices_buffer, vertex_count+3)
					indices_buffer = append(indices_buffer, vertex_count)

					vertex_count += 4
				}

				vertices := make([]vector.Vector, vertex_num)
				normals := make([]vector.Vector, vertex_num)
				us := make([]float32, vertex_num)
				vs := make([]float32, vertex_num)

				for i, ref := range face.References {
					vertices[i].X = float32(model.Vertices[ref.VertexIndex].X)
					vertices[i].Y = float32(model.Vertices[ref.VertexIndex].Y)
					vertices[i].Z = float32(model.Vertices[ref.VertexIndex].Z)

					normals[i].X = float32(model.Normals[ref.NormalIndex].X)
					normals[i].Y = float32(model.Normals[ref.NormalIndex].Y)
					normals[i].Z = float32(model.Normals[ref.NormalIndex].Z)

					us[i] = float32(model.TexCoords[ref.TexCoordIndex].U)
					vs[i] = float32(model.TexCoords[ref.TexCoordIndex].V)
				}

				for i := 0; i < vertex_num; i++ {
					pos_buffer = append(pos_buffer, vertices[i].X)
					pos_buffer = append(pos_buffer, vertices[i].Y)
					pos_buffer = append(pos_buffer, vertices[i].Z)

					norm_buffer = append(norm_buffer, normals[i].X)
					norm_buffer = append(norm_buffer, normals[i].Y)
					norm_buffer = append(norm_buffer, normals[i].Z)

					uv_buffer = append(uv_buffer, us[i])
					uv_buffer = append(uv_buffer, vs[i])
				}
			}

			buffered_vertices.SetIndicesBuffer(indices_buffer)
			buffered_vertices.SetPosBuffer(pos_buffer)
			buffered_vertices.SetNormBuffer(norm_buffer)
			buffered_vertices.SetUVBuffer(uv_buffer)

			ret = append(ret, buffered_vertices)
		}
	}

	return ret, nil
}
func LoadMtl(mtl_filename string) ([]*Material, error) {
	log.Printf("info: Start loading a MTL file. mtl_filename=%v", mtl_filename)

	ret := make([]*Material, 0)

	lines, err := file.GetFileAllLines(mtl_filename)
	if err != nil {
		return nil, err
	}

	newmtl_line_indexes := make([]int, 0)
	r := regexp.MustCompile("^newmtl")
	for i, line := range lines {
		if r.MatchString(line) {
			newmtl_line_indexes = append(newmtl_line_indexes, i)
		}
	}

	mtl_num := len(newmtl_line_indexes)
	for i := 0; i < mtl_num; i++ {
		var begin_line_index int
		var end_line_index int
		if i == mtl_num-1 {
			begin_line_index = newmtl_line_indexes[i]
			end_line_index = len(lines)
		} else {
			begin_line_index = newmtl_line_indexes[i]
			end_line_index = newmtl_line_indexes[i+1]
		}

		material := new(Material)
		material.Name = strings.Fields(lines[begin_line_index])[1]

		var err error
		for j := begin_line_index + 1; j < end_line_index; j++ {
			r_Ns := regexp.MustCompile("^Ns")
			r_Ka := regexp.MustCompile("^Ka")
			r_Kd := regexp.MustCompile("^Kd")
			r_Ks := regexp.MustCompile("^Ks")
			r_d := regexp.MustCompile("^d")
			r_MapKd := regexp.MustCompile("^Map_Kd")

			if r_Ns.MatchString(lines[j]) {
				Ns := strings.Fields(lines[j])[1]

				material.Specular_exponent, err = parseFloat32(Ns)
				if err != nil {
					return nil, err
				}
			} else if r_Ka.MatchString(lines[j]) {
				Ka := strings.Fields(lines[j])[1:]

				material.Ambient_color.R, err = parseFloat32(Ka[0])
				if err != nil {
					return nil, err
				}
				material.Ambient_color.G, err = parseFloat32(Ka[1])
				if err != nil {
					return nil, err
				}
				material.Ambient_color.B, err = parseFloat32(Ka[2])
				if err != nil {
					return nil, err
				}
				material.Ambient_color.A = 1.0
			} else if r_Kd.MatchString(lines[j]) {
				Kd := strings.Fields(lines[j])[1:]

				material.Diffuse_color.R, err = parseFloat32(Kd[0])
				if err != nil {
					return nil, err
				}
				material.Diffuse_color.G, err = parseFloat32(Kd[1])
				if err != nil {
					return nil, err
				}
				material.Diffuse_color.B, err = parseFloat32(Kd[2])
				if err != nil {
					return nil, err
				}
				material.Diffuse_color.A = 1.0
			} else if r_Ks.MatchString(lines[j]) {
				Ks := strings.Fields(lines[j])[1:]

				material.Specular_color.R, err = parseFloat32(Ks[0])
				if err != nil {
					return nil, err
				}
				material.Specular_color.G, err = parseFloat32(Ks[1])
				if err != nil {
					return nil, err
				}
				material.Specular_color.B, err = parseFloat32(Ks[2])
				if err != nil {
					return nil, err
				}
				material.Specular_color.A = 1.0
			} else if r_d.MatchString(lines[j]) {
				d := strings.Fields(lines[j])[1]

				material.Dissolve, err = parseFloat32(d)
				if err != nil {
					return nil, err
				}
			} else if r_MapKd.MatchString(lines[j]) {
				Map_Kd := strings.Fields(lines[j])[1]
				material.Diffuse_texture_map = Map_Kd
			}
		}

		ret = append(ret, material)
	}

	return ret, nil
}
func parseFloat32(str string) (float32, error) {
	f, err := strconv.ParseFloat(str, 32)
	return float32(f), err
}
func findMaterialByName(material_name string, materials []*Material) *Material {
	var ret *Material = nil
	for _, material := range materials {
		if material.Name == material_name {
			ret = material
			break
		}
	}

	return ret
}
