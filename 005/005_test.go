package ns005

import (
	"sort"
	"testing"
)

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 100000)
	for _, num := range nums {
		m[num]++
	}
	max, smax := 0, 0
	var res []int
	for index, count := range m {
		if count > smax {
			if count > max {
				temp := max
				max = count
				smax = temp
			} else {
				smax = count
			}
			res = append(res, index)
		}
	}
	return res[len(res)-k:]
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expexted []int
	}{
		{
			nums:     []int{1, 1, 1, 2, 2, 3},
			target:   2,
			expexted: []int{1, 2},
		},
		{
			nums:     []int{1},
			target:   1,
			expexted: []int{1},
		},
	}

	for i, test := range tests {
		sort.SliceStable(test.nums, func(i, j int) bool {
			return test.nums[i] < test.nums[j]
		})
		actual := TopKFrequent(test.nums, test.target)
		sort.SliceStable(actual, func(i, j int) bool {
			return actual[i] < actual[j]
		})
		for index, elem := range actual {
			if elem != test.nums[index] {
				t.Fatalf("%d - got: %d, wanted: %d", i, elem, test.nums[index])
			}
		}
	}
}
