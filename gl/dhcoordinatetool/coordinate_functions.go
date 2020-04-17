package dhcoordinatetool

func NormalizeCoordinate_Int(value int, max int) float32 {
	return 2.0*float32(value)/float32(max) - 1.0
}
func NormalizeCoordinate_Float32(value float32, max float32) float32 {
	return 2.0*value/max - 1.0
}

func ConvertWindowCoordinateAndOpenGLCoordinate_Y(y int, height int) int {
	return height - y
}
