package linkedlist

import "testing"

func TestListToSlice(t *testing.T) {
	head := &ListNode[int]{
		Val: 0,
		Next: &ListNode[int]{
			Val: 1,
			Next: &ListNode[int]{
				Val: 2,
				Next: &ListNode[int]{
					Val: 3,
					Next: &ListNode[int]{
						Val: 4,
						Next: &ListNode[int]{
							Val: 5,
						}}}},
		},
	}
	slc := []int{0, 1, 2, 3, 4, 5}
	actual := ListToSlice(head)
	if len(slc) != len(actual) {
		t.Errorf("expected len of %d, got len of %d", len(slc), len(actual))
	}

	for i, e := range slc {
		if e != actual[i] {
			t.Errorf("expected %d, got %d", e, actual[i])
		}
	}
}
