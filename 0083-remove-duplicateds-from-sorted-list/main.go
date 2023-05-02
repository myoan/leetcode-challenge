package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var recur func(*ListNode)
	recur = func(head *ListNode) {
		if head == nil || head.Next == nil {
			return
		}
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		}
		recur(head.Next)
	}
	recur(head)
	return head
}
