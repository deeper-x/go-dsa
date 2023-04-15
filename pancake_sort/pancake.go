package pancakesort

type data []int32

func (dl data) PancakeSort() {
	for uns := len(dl) - 1; uns > 0; uns-- {
		// find largest in unsorted range
		lx, lg := 0, dl[0]
		for i := 1; i <= uns; i++ {
			if dl[i] > lg {
				lx, lg = i, dl[i]
			}
		}
		// move to final position in two flips
		dl.flip(lx)
		dl.flip(uns)
	}
}

func (dl data) flip(r int) {
	for l := 0; l < r; l, r = l+1, r-1 {
		dl[l], dl[r] = dl[r], dl[l]
	}
}
