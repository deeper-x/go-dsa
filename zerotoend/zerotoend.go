package zerotoend

// Write an algorithm that takes an array and moves all of the zeros to the end, preserving the order of the other elements.
// https://www.codewars.com/kata/52597aa56021e91c93000cb0/train/go
// MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1})
// returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }

// Run build an zero-filled list, then populate it w/ !=0 values, as soon they're found
func Run(inL []int) []int {
	l := len(inL)
	res := make([]int, l)

	c := 0
	for i := 0; i < l; i++ {
		if inL[i] != 0 {
			res[c] = inL[i]
			c++
		}
	}

	return res
}
