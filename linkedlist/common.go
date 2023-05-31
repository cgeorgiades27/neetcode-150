package linkedlist

// Definition for singly-linked list.
type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

func ListToSlice[T any](head *ListNode[T]) []T {
	var slc []T
	for head != nil {
		slc = append(slc, head.Val)
		head = head.Next
	}
	return slc
}

func SliceToList[T any](slc []T) *ListNode[T] {
	head := &ListNode[T]{}
	curr := head
	for i := 0; i < len(slc); i++ {
		curr.Val = slc[i]
		if i == len(slc)-1 {
			break
		}
		curr.Next = &ListNode[T]{}
		curr = curr.Next
	}
	return head
}
