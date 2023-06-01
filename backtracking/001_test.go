package backtracking

import "testing"

// Given an integer array nums of unique elements, return all possible subsets (the power set).

// The solution set must not contain duplicate subsets. Return the solution in any order.

// Test case 1:
// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// Test case 2:
// Input: nums = [0]
// Output: [[],[0]]

func subsets(nums []int) [][]int {
	var allSubsets [][]int
	allSubsets = append(allSubsets, []int{}) // add the empty set
	for _, num := range nums {
		for _, subset := range allSubsets {
			allSubsets = append(allSubsets, append(subset, num))
		}
	}
	return allSubsets
}

func TestSubsets(t *testing.T) {
	tests := []struct {
		input  []int
		output [][]int
	}{
		{
			input:  []int{1, 2, 3},
			output: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			input:  []int{0},
			output: [][]int{{}, {0}},
		},
	}

	for _, test := range tests {
		actual := subsets(test.input)
		if len(test.output) != len(actual) {
			t.Errorf("expected len of %d, got len of %d", len(test.output), len(actual))
		}

		for j, subset := range test.output {
			if len(subset) != len(actual[j]) {
				t.Errorf("expected len of %d, got len of %d", len(subset), len(actual[j]))
			}
			for k, e := range subset {
				if e != actual[j][k] {
					t.Errorf("expected %d, got %d", e, actual[j][k])
				}
			}
		}
	}
}
