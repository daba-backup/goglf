package model

import (
	"log"
	"strings"

	"github.com/dabasan/go-dh3dbasis/matrix"
	"github.com/dabasan/go-dh3dbasis/vector"

	"github.com/dabasan/go-dhtool/filename"

	"github.com/dabasan/goglf/gl/model/bd1"
	"github.com/dabasan/goglf/gl/model/obj"
	"github.com/dabasan/goglf/gl/shader"
	"github.com/dabasan/goglf/gl/shape"
)

var count int
var models_map map[int]*ModelMgr

var keep_order_if_possible bool

func init() {
	count = 0
	models_map = make(map[int]*ModelMgr)

	keep_order_if_possible = false
}

func SetKeepOrderIfPossible(a_keep_order_if_possible bool) {
	keep_order_if_possible = a_keep_order_if_possible
}

func LoadModel(model_filename string, option FlipVOption) int {
	log.Printf("info: Start loading a model. model_filename=%v", model_filename)

	extension := filename.GetFileExtension(model_filename)
	extension = strings.ToLower(extension)

	var model *ModelMgr
	var err error
	switch extension {
	case "bd1":
		if keep_order_if_possible == true {
			model, err = loadBD1_KeepOrder(model_filename, option)
		} else {
			model, err = loadBD1(model_filename, option)
		}
	case "obj":
		model, err = loadOBJ(model_filename, option)
	default:
		log.Printf("error: Unsupported model format. extension=%v", extension)
		return -1
	}

	if err != nil {
		log.Printf("error: Failed to load a model. model_filename=%v", model_filename)
		log.Printf("error: %v", err)
		return -1
	}

	model_handle := count
	models_map[model_handle] = model
	count++

	return model_handle
}
func loadBD1(model_filename string, option FlipVOption) (*ModelMgr, error) {
	buffered_vertices_list, err := bd1.LoadBD1(model_filename)
	if err != nil {
		return nil, err
	}

	model := NewModelMgr(buffered_vertices_list, option)
	return model, nil
}
func loadBD1_KeepOrder(model_filename string, option FlipVOption) (*ModelMgr, error) {
	buffered_vertices_list, err := bd1.LoadBD1_KeepOrder(model_filename)
	if err != nil {
		return nil, err
	}

	model := NewModelMgr(buffered_vertices_list, option)
	return model, nil
}
func loadOBJ(model_filename string, option FlipVOption) (*ModelMgr, error) {
	buffered_vertices_list, err := obj.LoadOBJ(model_filename)
	if err != nil {
		return nil, err
	}

	model := NewModelMgr(buffered_vertices_list, option)
	return model, nil
}

func CopyModel(model_handle int) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	copied_model := model.Copy()

	copied_model_handle := count
	count++

	models_map[copied_model_handle] = copied_model

	return copied_model_handle
}
func DeleteModel(model_handle int) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	model.DeleteBuffers()
	delete(models_map, model_handle)

	return 0
}
func DeleteAllModels() {
	for _, model := range models_map {
		model.DeleteBuffers()
	}

	models_map = make(map[int]*ModelMgr, 0)
}

func ModelExists(model_handle int) bool {
	_, ok := models_map[model_handle]
	return ok
}

func AddProgram(model_handle int, program *shader.ShaderProgram) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	model.AddProgram(program)

	return 0
}
func RemoveAllPrograms(model_handle int) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	model.RemoveAllPrograms()

	return 0
}
func SetDefaultProgram(model_handle int) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	program, _ := shader.NewShaderProgram("texture")
	model.RemoveAllPrograms()
	model.AddProgram(program)

	return 0
}

func DrawModelWithProgram(model_handle int,
	program *shader.ShaderProgram, sampler_name string, texture_unit int) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	model.DrawWithProgram(program, sampler_name, texture_unit)

	return 0
}
func DrawModel(model_handle int, sampler_name string, texture_unit int) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	model.Draw(sampler_name, texture_unit)

	return 0
}
func DrawModel_Simple(model_handle int) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	model.Draw_Simple()

	return 0
}
func TransferModel(model_handle int) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	model.Transfer()

	return 0
}
func DrawModelElements(model_handle int, sampler_name string, texture_unit int, bound int) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	model.DrawElements(sampler_name, texture_unit, bound)

	return 0
}
func DrawModelElements_Simple(model_handle int, bound int) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	model.DrawElements_Simple(bound)

	return 0
}

func GetModelElementNum(model_handle int) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	ret := model.GetElementNum()
	return ret
}

func SetModelMatrix(model_handle int, m matrix.Matrix) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	model.SetMatrix(m)

	return 0
}

func TranslateModel(model_handle int, translate vector.Vector) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	translate_mat := matrix.MGetTranslate(translate)
	model.SetMatrix(translate_mat)

	return 0
}
func RotateModel(model_handle int, rotate vector.Vector) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	rot_x := matrix.MGetRotX(rotate.X)
	rot_y := matrix.MGetRotY(rotate.Y)
	rot_z := matrix.MGetRotZ(rotate.Z)
	rot := matrix.MMult(rot_y, rot_x)
	rot = matrix.MMult(rot_z, rot)

	model.SetMatrix(rot)

	return 0
}
func RotateModelLocally(model_handle int, origin vector.Vector, rotate vector.Vector) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	to_world_origin_vec := vector.VSub(vector.VGet(0.0, 0.0, 0.0), origin)
	to_world_origin_mat := matrix.MGetTranslate(to_world_origin_vec)
	rot_x := matrix.MGetRotX(rotate.X)
	rot_y := matrix.MGetRotY(rotate.Y)
	rot_z := matrix.MGetRotZ(rotate.Z)
	to_local_origin_mat := matrix.MGetTranslate(origin)

	mat := matrix.MMult(rot_x, to_world_origin_mat)
	mat = matrix.MMult(rot_y, mat)
	mat = matrix.MMult(rot_z, mat)
	mat = matrix.MMult(to_local_origin_mat, mat)

	model.SetMatrix(mat)

	return 0
}
func RescaleModel(model_handle int, scale vector.Vector) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	scale_mat := matrix.MGetScale(scale)
	model.SetMatrix(scale_mat)

	return 0
}

func ChangeModelTexture(model_handle int, material_index int, new_texture_handle int) int {
	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return -1
	}

	model.ChangeTexture(material_index, new_texture_handle)

	return 0
}

func GetModelFaces(model_handle int) []*shape.Triangle {
	ret := make([]*shape.Triangle, 0)

	model, ok := models_map[model_handle]
	if !ok {
		log.Printf("trace: No such model. model_handle=%v", model_handle)
		return ret
	}

	ret = model.GetFaces()
	return ret
}
