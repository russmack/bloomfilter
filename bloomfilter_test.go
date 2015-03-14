package bloomfilter

import (
	"testing"
)

func Test_Exists(t *testing.T) {
	loadStrings := []string{
		"cat", "dog", "mate", "frog", "moose",
		"el capitan", "spruce goose"}
	type Test struct {
		Input    string
		Expected bool
	}
	tests := []Test{
		{"klingon", false},
		{"frog", true},
		{"donkey", true},
		{"tame", true},
		{"spruce goose", true},
		{"light speed", false},
	}
	b := NewBloomFilter(15)
	for _, j := range loadStrings {
		b.Add(j)
	}

	for _, j := range tests {
		actual := b.Exists(j.Input)
		if actual != j.Expected {
			t.Error("\nExpected:", j.Expected, "\nGot:", actual, "\nFor:", j.Input)
		}
	}
}
