package main

import (
	"fmt"
	"github.com/russmack/bloomfilter"
)

func main() {
	// Add some items to the filter.
	b := bloomfilter.NewBloomFilter(22)
	b.Add("cat")
	b.Add("dog")
	b.Add("Mars")
	b.Add("Venus")

	// Display the probability of false positives.
	fmt.Println("\nProbability of false positives:", b.GetFalsePositiveProbability(), "\n")

	// Check if some items exist in the filter.
	subjects := []string{"frog", "manitee", "dog", "Deimos", "Venus", "Titan"}
	for _, j := range subjects {
		if b.Exists(j) {
			fmt.Println(j, "might be in the filter.")
		} else {
			fmt.Println(j, "is certainly not in the filter.")
		}
	}
}
