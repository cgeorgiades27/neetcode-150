package arrays

import "testing"

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

func TwoSumBrute(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func TwoSumOptimal(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if ind, ok := m[num]; ok {
			return []int{ind, i}
		}
		m[target-num] = i
	}
	return []int{}
}

func TestTwoSumBrute(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		want   []int
	}{
		{input: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{input: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{input: []int{3, 3}, target: 6, want: []int{0, 1}},
	}
	for i, test := range tests {
		indices := TwoSumBrute(test.input, test.target)
		if indices[0] != test.want[0] && indices[1] != test.want[1] {
			t.Fatalf("%d: got %v, want :%v", i, indices, test.want)
		}
	}
}
func TestTwoSumOptimal(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		want   []int
	}{
		{input: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{input: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{input: []int{3, 3}, target: 6, want: []int{0, 1}},
	}
	for i, test := range tests {
		indices := TwoSumOptimal(test.input, test.target)
		if indices[0] != test.want[0] && indices[1] != test.want[1] {
			t.Fatalf("%d: got %v, want :%v", i, indices, test.want)
		}
	}
}

func BenchmarkTwoSumBrute(b *testing.B) {
	input := []int{2, 7, 11, 15}
	for i := 0; i < b.N; i++ {
		TwoSumBrute(input, 9)
	}
}
func BenchmarkTwoSumOptimal(b *testing.B) {
	input := []int{2, 7, 11, 15}
	for i := 0; i < b.N; i++ {
		TwoSumOptimal(input, 9)
	}
}
