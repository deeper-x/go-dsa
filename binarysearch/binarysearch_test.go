package binarysearch

import "testing"

var haystack = []int{4, 6, 8, 9, 12, 56}
var OK = 8
var KO = 11

func TestBinarySearchOK(t *testing.T) {
	ok := Run(OK, haystack)

	if !ok {
		t.Error("needle should be found")
	}
}

func TestBinarySearchKO(t *testing.T) {
	ok := Run(KO, haystack)

	if ok {
		t.Error("needle shouldnt be found")
	}
}
