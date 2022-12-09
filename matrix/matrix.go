package matrix

// Run given size returns a bidimensional matrix
func Run(size int) [][]int {
	mtx := [][]int{}
	for i := 1; i <= size; i++ {
		row := []int{}

		for y := 1; y <= size; y++ {
			res := i * y
			row = append(row, res)
		}
		mtx = append(mtx, row)
	}

	return mtx
}
