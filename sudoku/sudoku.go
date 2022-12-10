package sudoku

import (
	"sort"
)

type row = []int
type coordsMap = map[int]row

// Grid sudoku 2d matrix
type Grid = []row

// defining squares
var sqr1 = []int{0, 2}
var sqr2 = []int{3, 5}
var sqr3 = []int{6, 8}
var blks = Grid{sqr1, sqr2, sqr3}

// square labels of a 9x9 sudoku grid
var sqrWins = Grid{
	{1, 4, 7},
	{2, 5, 8},
	{3, 6, 9},
	{10, 13, 16},
	{11, 14, 17},
	{12, 15, 18},
	{19, 22, 25},
	{20, 23, 26},
	{21, 24, 27},
}

// desired row in a 9x9 sudoku game
var valsRow = row{1, 2, 3, 4, 5, 6, 7, 8, 9}

func Run(c []row) bool {
	mtx := buildCoords(c)
	sqrChk := chkSquares(mtx)
	colChk := chkVert(c)
	rowChk := chkHoriz(c)

	return sqrChk && rowChk && colChk
}

func chkVert(gr Grid) bool {
	for i := 0; i < len(valsRow); i++ {
		col := row{}
		for _, v := range gr {
			col = append(col, v[i])
		}
		res := compareSlices(col, valsRow)
		if !res {
			return false
		}
	}
	return true
}

// check columns in grid
func chkHoriz(gr Grid) bool {
	for i := 0; i < len(gr); i++ {
		r := gr[i]
		sort.Ints(r)
		res := compareSlices(r, valsRow)
		if !res {
			return false
		}
	}
	return true
}

// check all squares in grid
func chkSquares(cm coordsMap) bool {
	for _, v := range sqrWins {
		resSlice := []int{}
		counter := 0
		for _, v1 := range v {
			counter++
			resSlice = append(resSlice, cm[v1]...)
			if counter == 3 {
				res := compareSlices(resSlice, valsRow)
				if !res {
					return false
				}
			}
		}
	}
	return true
}

// utility to compare slices
func compareSlices(res, exp row) bool {
	sort.Ints(res)
	for k, v := range res {
		if v != exp[k] {
			return false
		}
	}
	return true
}

// utility to label triplets' coords:
// 1  | 2  | 3
// 4  | 5  | 6
// 7  | 8  | 9

// 10 | 11 | 12
// 13 | 14 | 15
// 16 | 17 | 18

// 19 | 20 | 21
// 22 | 23 | 24
// 25 | 26 | 27
func buildCoords(m [][]int) coordsMap {
	cs := coordsMap{}
	counter := 0
	for _, v := range blks {
		for i := v[0]; i <= v[1]; i++ {
			for _, v1 := range blks {
				row := m[i]
				first := v1[0]
				last := v1[1]
				counter++
				record := []int{}
				for j := first; j <= last; j++ {
					record = append(record, row[j])
				}
				cs[counter] = record
			}
		}
	}
	return cs
}
