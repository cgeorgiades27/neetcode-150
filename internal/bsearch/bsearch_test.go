package bsearch

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		target   int
		arr      []int
		expected int
	}{
		{
			target:   1,
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 0,
		},
		{
			target:   11,
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: -1,
		},
	}

	for _, test := range tests {
		actual := BinarySearch(test.arr, test.target)
		if actual != test.expected {
			t.Fail()
		}
	}
}
