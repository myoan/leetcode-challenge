package main

import "testing"

func Test_addTwoNumbers(t *testing.T) {
	testcases := []struct {
		msg    string
		input1 []int
		input2 []int
		expect []int
	}{
		{
			msg:    "sample 1",
			input1: []int{2, 4, 3},
			input2: []int{5, 6, 4},
			expect: []int{7, 0, 8},
		},
		{
			msg:    "sample 2",
			input1: []int{0},
			input2: []int{0},
			expect: []int{0},
		},
		{
			msg:    "sample 3",
			input1: []int{9, 9, 9, 9, 9, 9, 9},
			input2: []int{9, 9, 9, 9},
			expect: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			msg:    "simple case2",
			input1: []int{2, 4, 9},
			input2: []int{5, 6, 4, 9},
			expect: []int{7, 0, 4, 0, 1},
		},
		{
			msg:    "huge number",
			input1: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			input2: []int{5, 6, 4},
			expect: []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
	}

	for _, tt := range testcases {
		list := addTwoNumbers(arrayToList(tt.input1), arrayToList(tt.input2))
		actual := listToArray(list)
		if !matchSlice(actual, tt.expect) {
			t.Errorf("[%s] act: %d, exp: %d\n", tt.msg, actual, tt.expect)
		}
	}
}
