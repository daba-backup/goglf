package bd1

func getFaceCorrespondingVertexIndices(face_index int) []int {
	var ret [4]int

	switch face_index {
	case 0:
		ret[0] = 0
		ret[1] = 1
		ret[2] = 2
		ret[3] = 3
	case 1:
		ret[0] = 5
		ret[1] = 4
		ret[2] = 7
		ret[3] = 6
	case 2:
		ret[0] = 1
		ret[1] = 0
		ret[2] = 4
		ret[3] = 5
	case 3:
		ret[0] = 2
		ret[1] = 1
		ret[2] = 5
		ret[3] = 6
	case 4:
		ret[0] = 3
		ret[1] = 2
		ret[2] = 6
		ret[3] = 7
	case 5:
		ret[0] = 0
		ret[1] = 3
		ret[2] = 7
		ret[3] = 4
	default:
		for i := 0; i < 4; i++ {
			ret[i] = 0
		}
	}

	return ret[:]
}
func getFaceCorrespondingUVIndices(face_index int) []int {
	var ret [4]int

	switch face_index {
	case 0:
		ret[0] = 3
		ret[1] = 2
		ret[2] = 1
		ret[3] = 0
	case 1:
		ret[0] = 7
		ret[1] = 6
		ret[2] = 5
		ret[3] = 4
	case 2:
		ret[0] = 9
		ret[1] = 8
		ret[2] = 11
		ret[3] = 10
	case 3:
		ret[0] = 13
		ret[1] = 12
		ret[2] = 15
		ret[3] = 14
	case 4:
		ret[0] = 17
		ret[1] = 16
		ret[2] = 19
		ret[3] = 18
	case 5:
		ret[0] = 21
		ret[1] = 20
		ret[2] = 23
		ret[3] = 22
	default:
		for i := 0; i < 4; i++ {
			ret[i] = 0
		}
	}

	return ret[:]
}
