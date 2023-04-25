package arrays

import (
	"cgeorgiades27/neetcode-150/internal/intheap"
	"container/heap"
	"sort"
	"testing"
)

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 100000)
	for _, num := range nums {
		m[num]++
	}

	h := &intheap.IntHeap{}
	for _, count := range m {
		if count > 0 {
			heap.Push(h, count)
		}
	}

	var resSlc []int
	for i := 0; i < k; i++ {
		val := heap.Pop(h)
		resSlc = append(resSlc, val.(int))
	}

	return resSlc
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
		sort.SliceStable(test.expexted, func(i, j int) bool {
			return test.expexted[i] < test.expexted[j]
		})
		actual := TopKFrequent(test.nums, test.target)
		sort.SliceStable(actual, func(i, j int) bool {
			return actual[i] < actual[j]
		})
		for index, elem := range actual {
			if elem != test.expexted[index] {
				t.Fatalf("%d - got: %d, wanted: %d", i, elem, test.expexted[index])
			}
		}
	}
}
