package linearsearch

// Run linearsearch
func Run(key int, datalist []int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}
