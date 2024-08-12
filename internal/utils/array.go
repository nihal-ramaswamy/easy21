package utils

func Make3dArray(x, y, z int) [][][]float64 {
	arr := make([][][]float64, x)
	for i := range arr {
		arr[i] = make([][]float64, y)
		for j := range arr[i] {
			arr[i][j] = make([]float64, z)
		}
	}
	return arr
}
