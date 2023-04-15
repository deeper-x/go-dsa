package pancakesort

import (
	"testing"
)

func TestPancakeSort(t *testing.T) {
	expected := dataList{-26, -23, 11, 28, 44, 59, 158, 193, 503, 997}
	inList := dataList{28, 11, 59, -26, 503, 158, 997, 193, -23, 44}

	inList.PancakeSort()
	for k, v := range inList {
		if expected[k] != v {
			t.Error("pancakes are not sorted!")
		}
	}
}
