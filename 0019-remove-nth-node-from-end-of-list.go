package main

func RemoveChild(n *ListNode) {
	child := n.Next
	n.Next = child.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := head
	parent := head
	i := 0
	if head.Next == nil {
		return nil
	}
	for {
		if tmp.Next == nil {
			break
		}
		if i > n-1 {
			parent = parent.Next
		}
		tmp = tmp.Next
		i++
	}
	if i+1 == n {
		return head.Next
	}
	RemoveChild(parent)
	return head
}
