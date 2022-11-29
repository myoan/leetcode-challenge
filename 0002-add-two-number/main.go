package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return _addTwoNumber(l1, l2, 0)
}

func callNextNode(l1 *ListNode, l2 *ListNode, val int) *ListNode {
	if val >= 10 {
		return &ListNode{
			Val:  val - 10,
			Next: _addTwoNumber(l1, l2, 1),
		}
	}
	return &ListNode{
		Val:  val,
		Next: _addTwoNumber(l1, l2, 0),
	}
}

func _addTwoNumber(l1 *ListNode, l2 *ListNode, inc int) *ListNode {

	if l1 == nil && l2 == nil {
		if inc != 0 {
			return &ListNode{
				Val:  1,
				Next: nil,
			}
		}
		return nil
	} else if l1 == nil {
		return callNextNode(nil, l2.Next, l2.Val+inc)
	} else if l2 == nil {
		return callNextNode(l1.Next, nil, l1.Val+inc)
	}

	return callNextNode(l1.Next, l2.Next, l1.Val+l2.Val+inc)
}
