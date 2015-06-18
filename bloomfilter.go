// Package bloomfilter provides a bloom filter.
// A Bloom filter provides a quick and memory-efficient, probabilistic check for
// the existence of a member in a set.
// The efficiency comes at the speed of certainty.
// False positives are possible, wrongly indicating that an element is in the set,
// ie that the element is possibly in the set.
// False negatives are not possible - if the filter indicates that an element is
// not in the set, then it is definitely not in the set.
package bloomfilter

import (
	"fmt"
	"hash/fnv"
	"math"
)

// Hash32Fn is a function type for 32 bit hashing functions.
type Hash32Fn func(string) uint32

// BloomFilter is the public struct.
type BloomFilter struct {
	// TODO: use bits, not bools.
	Filter       []bool
	Size         uint32
	HashFuncs    []Hash32Fn
	TotalFlipped uint32
}

// NewBloomFilter creates a new BloomFilter with the specified number of switches,
// and a list of the hash functions to use when adding elements to the set, and
// when checking for existence.
func NewBloomFilter(filterSize uint32) *BloomFilter {
	b := BloomFilter{}
	b.Filter = make([]bool, filterSize)
	b.Size = filterSize
	// TODO: inject preferred choice of hashers.
	b.HashFuncs = []Hash32Fn{hashFnv1, hashFnv1a}
	return &b
}

// hashFnv1 puts a string through the golang stdlib 32-bit FNV-1 hash.
// A string is reduced to an int.
func hashFnv1(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}

// hashFnv1a puts a string through the golang stdlib 32-bit FNV-1a hash.
// A string is reduced to an int.
func hashFnv1a(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// getIndex32 hashes a string, then reduces that hash to an index within the bounds of the filter.
func (b *BloomFilter) getIndex32(s string, hashFn Hash32Fn) uint32 {
	h := hashFn(s)
	return h % b.Size
}

// Add a string to the Bloom filter.
func (b *BloomFilter) Add(s string) {
	fmt.Println("\nAdding:", s)
	// Iterate over the list of hash functions, using each to reduce the string
	// to a single index in the filter, which is then flipped on.
	for i, j := range b.HashFuncs {
		idx := b.getIndex32(s, j)
		fmt.Println("Hash function #", i, "hashed to index:", idx)
		b.Filter[idx] = true
		fmt.Println("Updated filter to:", b.Filter)
		b.TotalFlipped++
	}
}

// Exists checks if the given string is in the Bloom Filter.
func (b *BloomFilter) Exists(s string) bool {
	// Put the candidate string through each of the hash functions, and for each
	// index returned get the value at that index in the bloom filter, and put
	// those values into the results slice.
	results := make([]bool, len(b.HashFuncs))
	for i, j := range b.HashFuncs {
		idx := b.getIndex32(s, j)
		val := b.Filter[idx]
		results[i] = val
	}
	allTrue := true
	// Iterate over the switch values retrieved from the bloom filter.
	for _, j := range results {
		// If the switch values retrieved are all true we'll end up returning true.
		allTrue = allTrue && j
	}
	return allTrue
}

// False positive probability:
// The number of true cells in the filter,
// divided by the length of the filter,
// to the power of the number of hash functions.
func (b *BloomFilter) GetFalsePositiveProbability() float64 {
	x := float64(b.TotalFlipped) / float64(b.Size)
	y := float64(len(b.HashFuncs))
	z := math.Pow(x, y) * 100
	p := math.Floor(z + .5)
	return p
}
