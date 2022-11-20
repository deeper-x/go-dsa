package main

import (
	"github.com/deeper-x/go-dsa/binarysearch"
	"github.com/deeper-x/go-dsa/linearsearch"
)

var needle = 10
var haystack = []int{1, 5, 6, 7, 8, 9, 10, 11, 13, 15}

func main() {
	linearsearch.Run(needle, haystack)
	binarysearch.Run(needle, haystack)
}
