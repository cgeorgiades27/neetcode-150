package arrays

import (
	"sort"
	"testing"
)

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

func GroupAnagrams(strs []string) [][]string {
	var groups [][]string
	m := make(map[string][]string)
	for _, str := range strs {
		word := []byte(str)
		sort.SliceStable(word, func(i, j int) bool {
			return word[i] < word[j]
		})
		m[string(word)] = append(m[string(word)], str)
	}
	for _, v := range m {
		groups = append(groups, v)
	}
	return groups
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input  []string
		output [][]string
	}{
		{
			input:  []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			output: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			input:  []string{""},
			output: [][]string{{""}},
		},
		{
			input:  []string{"a"},
			output: [][]string{{"a"}},
		},
	}

	for testNum, test := range tests {
		sort.SliceStable(test.output, func(i, j int) bool {
			return len(test.output[i]) < len(test.output[j])
		})
		for _, out := range test.output {
			sort.SliceStable(out, func(i, j int) bool {
				return out[i] < out[j]
			})
		}

		actual := GroupAnagrams(test.input)
		sort.SliceStable(actual, func(i, j int) bool {
			return len(actual[i]) < len(actual[j])
		})

		for index, group := range actual {
			sort.SliceStable(group, func(i, j int) bool {
				return group[i] < group[j]
			})
			if len(group) != len(test.output[index]) {
				t.Fatalf("%d - wanted: %v, got: %v", testNum, test.output[index], group)
			}
			for gIndex, entry := range group {
				if test.output[index][gIndex] != entry {
					t.Fatalf("%d - wanted: %v, got: %v", testNum, test.output[index], group)
				}
			}
		}
	}
}
