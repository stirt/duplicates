package main

import (
	"flag"
	"fmt"

	dups "github.com/stirt/duplicates/pkg/duplicates"
)

func main() {
	n := flag.Int("n", 1000000, "define number of items in createDuplicates list")
	maxDups := flag.Int("maxDups", 1, "define a maximum number of duplicates to be generated")
	flag.Parse()
	duplicates := dups.CreateDuplicates(*n, *maxDups)
	foundDups := dups.FindDuplicates(duplicates)
	fmt.Println(foundDups)
}
