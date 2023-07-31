package backtracking

import (
	"fmt"
	"testing"
)

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

func subsetsRecursive(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	first := nums[0]
	others := nums[1:]
	subsets := subsetsRecursive(others)

	var newSets [][]int
	for _, subset := range subsets {
		subset = append(subset, first)
		newSets = append(newSets, subset)
	}

	subsets = append(subsets, newSets...)
	return subsets
}

func subsetsRecursive2(nums []int) [][]int {
	var result [][]int
	var backtrack func(int, []int)
	backtrack = func(start int, track []int) {
		// make a copy of track
		trackCopy := make([]int, len(track))
		copy(trackCopy, track)
		result = append(result, trackCopy)

		for i := start; i < len(nums); i++ {
			// make a choice
			track = append(track, nums[i])
			// backtrack
			backtrack(i+1, track)
			// undo the choice
			track = track[:len(track)-1]
		}
	}

	backtrack(0, []int{})
	return result
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
		fmt.Println(actual)
		fmt.Println(test.output)
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

func TestSubsetsRecursive(t *testing.T) {
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
		actual := subsetsRecursive(test.input)
		fmt.Println(actual)
		fmt.Println(test.output)
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

func TestSubsetsRecursive2(t *testing.T) {
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
		actual := subsetsRecursive2(test.input)
		fmt.Println(actual)
		fmt.Println(test.output)
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
