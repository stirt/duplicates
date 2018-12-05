package duplicates

import (
	"math/rand"
	"time"
)

// CreateDuplicates creates a slice of [0,n+1) numbers
// using param "n" as n.
// maxDups allows for a pseudo-randomly generated number of dups to
// be added, from [0, maxDups).
func CreateDuplicates(n int, maxDups int) []int {
	rand.Seed(time.Now().UnixNano())
	var dupsAdded int
	noOfDups := rand.Intn(maxDups)
	if noOfDups == 0 {
		noOfDups = 1
	}
	dups := rand.Perm(n - noOfDups + 1)
	for dupsAdded < noOfDups {
		insertAt := rand.Intn(len(dups) - 1)
		randIndex := rand.Intn(len(dups) - 1)
		dups = append(dups[:insertAt], append([]int{dups[randIndex]}, dups[insertAt:]...)...)
		dupsAdded++
	}
	return dups
}

// FindDuplicates finds duplicate occurances of ints in a given intSlice and.
// returns those duplicates in a new intSlice. We use maps here as it offers
// constant time getting and setting, along with the guarantee
// that duplicates cannot be stored in them(as keys must be unique).
func FindDuplicates(is []int) []int {
	found := make(map[int]int)
	seen := make(map[int]bool)
	for _, d := range is {
		if !seen[d] {
			seen[d] = true
			continue
		}
		// We could potentially store an empty struct{} here instead
		// as it takes up no space in memory.
		found[d]++
	}
	var dups []int
	for key := range found {
		dups = append(dups, key)
	}
	return dups
}
