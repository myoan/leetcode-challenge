package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteDuplicates(t *testing.T) {
	testcases := []struct {
		name   string
		nodes  *ListNode
		expect *ListNode
	}{
		{
			name: "simple",
			nodes: &ListNode{
				Val: 1,
			},
			expect: &ListNode{
				Val: 1,
			},
		},
		{
			name: "simple 2",
			nodes: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
				},
			},
			expect: &ListNode{
				Val: 1,
			},
		},
		{
			name: "2nd example",
			nodes: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
			},
			expect: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := deleteDuplicates(tc.nodes)
			assert.Equal(t, tc.expect, actual)
		})
	}
}
