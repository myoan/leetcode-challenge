package main

import "testing"

func Test_mergeKLists(t *testing.T) {
	testcases := []struct {
		msg    string
		input  [][]int
		expect []int
	}{
		{
			msg: "single array",
			input: [][]int{
				{1, 2, 3},
			},
			expect: []int{1, 2, 3},
		},
		{
			msg: "double array",
			input: [][]int{
				{2, 4, 6},
				{1, 3, 5},
			},
			expect: []int{1, 2, 3, 4, 5, 6},
		},
		{
			msg: "example 1",
			input: [][]int{
				{1, 4, 5},
				{1, 3, 4},
				{2, 6},
			},
			expect: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			msg:    "sample 2",
			input:  [][]int{},
			expect: []int{},
		},
		{
			msg:    "sample 3",
			input:  [][]int{{}},
			expect: []int{},
		},
	}

	for _, tt := range testcases {
		list := make([]*ListNode, len(tt.input))
		for i, a := range tt.input {
			list[i] = arrayToList(a)
		}
		actual := mergeKLists(list)
		aa := listToArray(actual)
		if !matchIntArray(aa, tt.expect) {
			t.Errorf("[%s] act: %v, exp: %v\n", tt.msg, aa, tt.expect)
		}
	}
}
