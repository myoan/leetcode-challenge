package main

import "testing"

func Test_reverseKGroup(t *testing.T) {
	testcases := []struct {
		msg    string
		input  []int
		k      int
		expect []int
	}{
		{
			msg:    "example 1",
			input:  []int{1, 2, 3, 4, 5},
			k:      2,
			expect: []int{2, 1, 4, 3, 5},
		},
		{
			msg:    "example 2",
			input:  []int{1, 2, 3, 4, 5},
			k:      3,
			expect: []int{3, 2, 1, 4, 5},
		},
		{
			msg:    "example 3",
			input:  []int{1, 2, 3, 4, 5},
			k:      1,
			expect: []int{1, 2, 3, 4, 5},
		},
		{
			msg:    "example 4",
			input:  []int{1},
			k:      1,
			expect: []int{1},
		},
	}

	for _, tt := range testcases {
		list := arrayToList(tt.input)
		actual := reverseKGroup(list, tt.k)
		aa := listToArray(actual)
		if !matchIntArray(aa, tt.expect) {
			t.Errorf("[%s] act: %v, exp: %v\n", tt.msg, aa, tt.expect)
		}
	}
}
