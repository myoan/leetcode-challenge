package main

import "testing"

func Test_swapPairs(t *testing.T) {
	testcases := []struct {
		msg    string
		input  []int
		expect []int
	}{
		{
			msg:    "example 1",
			input:  []int{1, 2, 3, 4},
			expect: []int{2, 1, 4, 3},
		},
		{
			msg:    "example 2",
			input:  []int{},
			expect: []int{},
		},
		{
			msg:    "example 3",
			input:  []int{1},
			expect: []int{1},
		},
	}

	for _, tt := range testcases {
		list := arrayToList(tt.input)
		actual := swapPairs(list)
		aa := listToArray(actual)
		if !matchIntArray(aa, tt.expect) {
			t.Errorf("[%s] act: %v, exp: %v\n", tt.msg, aa, tt.expect)
		}
	}
}
