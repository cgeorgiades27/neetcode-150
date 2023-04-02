package nc001

import "testing"

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

const MAX_NUM = 1000

func containsDuplicate(nums []int) bool {
	m := make(map[int]int, MAX_NUM)
	for _, i := range nums {
		if m[i] > 0 {
			return true
		}
		m[i]++
	}
	return false
}

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		input []int
		want  bool
	}{
		{input: []int{1, 2, 3, 1}, want: true},
		{input: []int{1, 2, 3, 4}, want: false},
		{input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},
	}
	for i, test := range tests {
		if actual := containsDuplicate(test.input); actual != test.want {
			t.Fatalf("%d: wanted %t, got %t", i, actual, test.want)
		}
	}
}
