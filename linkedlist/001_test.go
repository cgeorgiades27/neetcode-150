package linkedlist

import (
	"fmt"
	"testing"
)

// Given the head of a singly linked list, reverse the list, and return the reversed list.

func ReverseList(head *ListNode[int]) *ListNode[int] {
	var slc []*ListNode[int]
	for head != nil {
		slc = append([]*ListNode[int]{head}, slc...)
		head = head.Next
	}

	newHead := &ListNode[int]{}
	if len(slc) > 0 {
		newHead = slc[0]
	}
	for i, node := range slc {
		if i+1 < len(slc)-1 {
			node.Next = slc[i+1]
		}
		fmt.Println(node)
	}
	return newHead
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			output: []int{5, 4, 3, 2, 1},
		},
		{
			input:  []int{1, 2},
			output: []int{2, 1},
		},
	}

	for i, test := range tests {
		testSlc := SliceToList(test.input)
		actual := ReverseList(testSlc)
		actualSlc := ListToSlice(actual)
		for j, e := range test.output {
			if actualSlc[j] != e {
				t.Errorf("%d - expected %d, got %d", i, e, actualSlc[j])
			}
		}
	}
}
