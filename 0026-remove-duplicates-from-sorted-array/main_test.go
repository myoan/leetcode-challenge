package main

import "testing"

func Test_removeDuplicates(t *testing.T) {
	testcases := []struct {
		msg         string
		input       []int
		expect      int
		expectArray []int
	}{
		{
			msg:         "empty case",
			input:       []int{},
			expect:      0,
			expectArray: []int{},
		},
		{
			msg:         "simple",
			input:       []int{1},
			expect:      1,
			expectArray: []int{1},
		},
		{
			msg:         "example 1",
			input:       []int{1, 1, 2},
			expect:      2,
			expectArray: []int{1, 2, 2},
		},
		{
			msg:         "example 2",
			input:       []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expect:      5,
			expectArray: []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4},
		},
	}

	for _, tt := range testcases {
		actual := removeDuplicates(tt.input)
		if actual != tt.expect {
			t.Errorf("[%s(check result)] act: %d, exp: %d\n", tt.msg, actual, tt.expect)
		}
		if !matchIntArray(tt.input, tt.expectArray) {
			t.Errorf("[%s(check array)] act: %v, exp: %v\n", tt.msg, tt.input, tt.expectArray)
		}
	}
}
