package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/deeper-x/go-dsa/deadfish"
)

// DataSearch result object
type DataSearch struct {
	Name   string
	Result bool
}

var needle = 11
var haystack = []int{1, 5, 6, 7, 8, 9, 10, 11, 13, 15}
var randSlice = genRandomSlice(10)

func main() {
	// ls := NewData("Linear Search", linearsearch.Run(needle, haystack))
	// bs := NewData("Binary Search", binarysearch.Run(needle, haystack))

	// printSearch(ls)
	// printSearch(bs)

	// bubblesort.Run(randSlice)
	// log.Println("Sorted slice:", randSlice)

	// btree := binarytree.New()
	// btree.Add(10).Add(20).Add(30).Add(100).Add(15).Add(25)
	// binarytree.Dump(os.Stdout, btree.Root, 0, 'M')

	log.Println(deadfish.Run("iiisdoso"))
	log.Println(deadfish.Run("ooo"))
	log.Println(deadfish.Run("ioioio"))
	log.Println(deadfish.Run("codewars"))
	log.Println(deadfish.Run("isoisoiso"))
	log.Println(deadfish.Run("codewars"))
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

func genRandomSlice(size int) []int {
	res := make([]int, size)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		res[i] = rand.Intn(999)
	}

	return res
}
