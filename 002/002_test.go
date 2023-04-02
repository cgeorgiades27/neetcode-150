package nc002

import (
	"sort"
	"testing"
)

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

func isAnagram(s string, t string) bool {
	rs := []byte(s)
	rt := []byte(t)

	sort.SliceStable(rs, func(i, j int) bool {
		return rs[i] > rs[j]
	})
	sort.SliceStable(rt, func(i, j int) bool {
		return rt[i] > rt[j]
	})
	for i, r := range rs {
		if rt[i] != r {
			return false
		}
	}
	return len(rs) == len(rt)
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		input []string
		want  bool
	}{
		{input: []string{"anagram", "nagaram"}, want: true},
		{input: []string{"rat", "car"}, want: false},
	}
	for i, test := range tests {
		if actual := isAnagram(test.input[0], test.input[1]); actual != test.want {
			t.Fatalf("%d: got %t, want %t", i, actual, test.want)
		}
	}
}
