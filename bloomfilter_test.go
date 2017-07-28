// Package bloomfilter implements a Bloom filter.
package bloomfilter

import (
	"testing"
)

// Test_Exists tests the Exists method of BloomFilter.
func Test_Exists(t *testing.T) {
	// Strings to add to the filter as sample existing data.
	loadStrings := []string{
		"cat", "dog", "mate", "frog", "moose",
		"el capitan", "spruce goose"}

	// Test struct to hold each test case.
	type TestCase struct {
		Input    string
		Expected bool
	}

	// Build a list of test cases.
	testcases := []TestCase{
		{"klingon", false},
		{"frog", true},
		{"donkey", true},
		{"tame", true},
		{"spruce goose", true},
		{"light speed", false},
	}

	// Create a new BloomFilter and add the sample data.
	b := NewBloomFilter(15)
	for _, j := range loadStrings {
		b.Add(j)
	}

	// Test each case.
	for _, j := range testcases {
		actual := b.Exists(j.Input)
		if actual != j.Expected {
			t.Error("\nExpected:", j.Expected, "\nGot:", actual, "\nFor:", j.Input)
		}
	}
}

func Benchmark_Add(b *testing.B) {
	b.ReportAllocs()

	// Strings to add to the filter as sample existing data.
	loadStrings := []string{
		"cat", "dog", "mate", "frog", "moose",
		"el capitan", "spruce goose"}

	f := NewBloomFilter(15)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, j := range loadStrings {
			f.Add(j)
		}
	}
}

func Benchmark_Exists(b *testing.B) {
	b.ReportAllocs()

	checkStrings := []string{
		"klingon", "frog", "donkey",
		"tame", "spruce goose", "light speed",
	}

	f := NewBloomFilter(15)

	// Strings to add to the filter as sample existing data.
	loadStrings := []string{
		"cat", "dog", "mate", "frog", "moose",
		"el capitan", "spruce goose"}

	for _, j := range loadStrings {
		f.Add(j)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, j := range checkStrings {
			_ = f.Exists(j)
		}
	}

}
