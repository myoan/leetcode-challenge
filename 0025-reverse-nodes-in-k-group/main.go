package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	tmp := head
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	ret := dummy
	group := make([]*ListNode, k)

	for {

		if tmp == nil {
			break
		}
		if tmp.Next == nil {
			dummy.Next = tmp
			break
		}

		for i := 0; i < k; i++ {
			if tmp == nil {
				dummy.Next = group[0]
				goto L
			}
			group[i] = tmp
			tmp = tmp.Next
		}

		for i := 0; i < k; i++ {
			group[i].Next = nil
		}

		for i := k - 1; i >= 0; i-- {
			dummy.Next = group[i]
			dummy = dummy.Next
		}
	}
L:

	return ret.Next
}
