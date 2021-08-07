package main

import "testing"

func Test_removeNthFromEnd(t *testing.T) {
	testcases := []struct {
		msg    string
		input  []int
		target int
		expect []int
	}{
		{
			msg:    "sample 1",
			input:  []int{1, 2, 3, 4, 5},
			target: 2,
			expect: []int{1, 2, 3, 5},
		},
		{
			msg:    "sample 2",
			input:  []int{1, 2},
			target: 2,
			expect: []int{2},
		},
	}

	for _, tt := range testcases {
		list := removeNthFromEnd(arrayToList(tt.input), tt.target)
		actual := listToArray(list)
		if !matchSlice(actual, tt.expect) {
			t.Errorf("[%s] act: %v, exp: %v\n", tt.msg, actual, tt.expect)
		}
	}
}
