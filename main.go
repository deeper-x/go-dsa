package main

import (
	"log"

	"github.com/deeper-x/go-dsa/binarysearch"
	"github.com/deeper-x/go-dsa/linearsearch"
)

// DataSearch result object
type DataSearch struct {
	Name   string
	Result bool
}

var needle = 11
var haystack = []int{1, 5, 6, 7, 8, 9, 10, 11, 13, 15}

func main() {
	ls := NewData("Linear Search", linearsearch.Run(needle, haystack))
	bs := NewData("Binary Search", binarysearch.Run(needle, haystack))

	printSearch(ls)
	printSearch(bs)
}

// printSearch data printing utility
func printSearch(ds DataSearch) {
	log.Printf("%s: %t", ds.Name, ds.Result)
}

// NewData builder
func NewData(name string, result bool) DataSearch {
	return DataSearch{
		Name:   name,
		Result: result,
	}
}
