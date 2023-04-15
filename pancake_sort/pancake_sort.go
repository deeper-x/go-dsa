package pancakesort

type dataList []int64

func (dl dataList) PancakeSort() {
	// 1. start from the second to last element
	for nsrtd := len(dl) - 1; nsrtd > 0; nsrtd-- {
		// 2. pick the first and last in list
		lx, lg := 0, dl[0]
		// 3. iterate all elements in this sublist
		for i := 1; i <= nsrtd; i++ {
			// 4. if current > last overwrite first and last
			if dl[i] > lg {
				lx, lg = i, dl[i]
			}
		}
		// 5. move to final position in two flips
		dl.flip(lx)
		dl.flip(nsrtd)
	}
}

func (dl dataList) flip(r int) {
	for l := 0; l < r; l, r = l+1, r-1 {
		dl[l], dl[r] = dl[r], dl[l]
	}
}
