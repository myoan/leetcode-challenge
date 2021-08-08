package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func arrayToList(a []int) *ListNode {
	var next *ListNode
	if len(a) == 0 {
		return nil
	}
	if len(a) == 1 {
		next = nil
	} else {
		next = arrayToList(a[1:])
	}
	return &ListNode{
		Val:  a[0],
		Next: next,
	}
}

func listToArray(list *ListNode) []int {
	l := list
	a := make([]int, 0)
	for {
		if l == nil {
			return a
		}
		a = append(a, l.Val)
		l = l.Next
	}
}

func matchSlice(a1 []int, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i, e := range a1 {
		if a2[i] != e {
			return false
		}
	}
	return true
}

func (l *ListNode) Show() {
	fmt.Printf("[ %d ", l.Val)
	tmp := l.Next
	for {
		if tmp == nil {
			fmt.Println("]")
			return
		}
		fmt.Printf("-> %d ", tmp.Val)
		tmp = tmp.Next
	}
}

func matchStringArray(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, e := range a {
		if e != b[i] {
			return false
		}
	}
	return true
}

func matchIntArray(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, e := range a {
		if e != b[i] {
			return false
		}
	}
	return true
}

func matchIntArrayOfArray(a1, a2 [][]int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for _, e1 := range a1 {
		flag := false
		for _, e2 := range a2 {
			if matchIntArray(e1, e2) {
				flag = true
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
