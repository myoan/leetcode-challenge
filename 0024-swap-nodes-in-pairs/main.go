package main

func swapPairs(head *ListNode) *ListNode {
	tmp := head
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	ret := dummy

	for {
		if tmp == nil {
			break
		}
		if tmp.Next == nil {
			dummy.Next = tmp
			break
		}

		a := tmp
		tmp = tmp.Next
		b := tmp
		tmp = tmp.Next

		a.Next = nil
		b.Next = nil

		dummy.Next = b
		dummy = dummy.Next
		dummy.Next = a
		dummy = dummy.Next
	}

	return ret.Next
}
