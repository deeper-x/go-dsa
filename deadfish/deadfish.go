package deadfish

var tot int
var sum int
var res []int

// Run deadfish program
func Run(inStr string) []int {
	tot = 0
	sum = 0
	res = []int{}

	for _, v := range inStr {
		ok, cur := calc(v)

		if ok {
			res = append(res, cur)
			continue
		}

		sum += cur
	}

	return res
}

// calc given an input ascii value, returning its corresponding command with the "print" flag
func calc(inChar rune) (bool, int) {
	// d 100
	// i 105
	// o 111
	// s 115

	switch inChar {
	case 100:
		tot--
		return false, tot

	case 105:
		tot++
		return false, tot

	case 111:
		return true, tot

	case 115:
		tot = tot * tot
		return false, tot

	default:
		return false, 0
	}
}
