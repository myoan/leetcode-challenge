package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	beforeRet := dummy
	for {
		allnil := true
		for _, a := range lists {
			if a != nil {
				allnil = false
			}
		}
		if allnil {
			break
		}

		minidx := -1
		min := 100000
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil && min > lists[i].Val {
				min = lists[i].Val
				minidx = i
			}
		}
		dummy.Next = &ListNode{
			Val:  min,
			Next: nil,
		}
		dummy = dummy.Next
		lists[minidx] = lists[minidx].Next
	}
	return beforeRet.Next
}
