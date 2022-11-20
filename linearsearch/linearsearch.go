package linearsearch

// Run linearsearch
func Run(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}
