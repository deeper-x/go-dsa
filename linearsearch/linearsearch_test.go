package linearsearch

import (
	"testing"
)

var haystack = []int{1, 5, 4, 5, 6, 3, 6, 7, 3, 1}
var needleOK = 7
var needleKO = 12

func TestLinearsearchOK(t *testing.T) {
	ok := Run(haystack, needleOK)

	if !ok {
		t.Error("needle should have been found")
	}
}

func TestLinearSearchKO(t *testing.T) {
	ok := Run(haystack, needleKO)
	if ok {
		t.Error("needle shouldn't have been found")
	}
}
