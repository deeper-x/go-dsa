package bubblesort

import "testing"

var data = []int{6, 1, 3, 7, 2, 5, 4, 8, 9}
var expected = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestBubbleSort(t *testing.T) {
	for k := range data {
		if data[k] != expected[k] {
			t.Errorf("slices not equal")
			return
		}
	}
}
