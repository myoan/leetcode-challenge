package main

import "testing"

func Test_removeElement(t *testing.T) {
	testcases := []struct {
		msg         string
		input       []int
		val         int
		expect      int
		expectArray []int
	}{
		{
			msg:         "empty case",
			input:       []int{1, 2, 3},
			val:         3,
			expect:      2,
			expectArray: []int{1, 2, 99},
		},
		{
			msg:         "empty case",
			input:       []int{3, 2, 2, 3},
			val:         3,
			expect:      2,
			expectArray: []int{2, 2, 99, 99},
		},
		{
			msg:         "example 1",
			input:       []int{3, 2, 2, 3},
			val:         3,
			expect:      2,
			expectArray: []int{2, 2, 99, 99},
		},
		{
			msg:         "example 2",
			input:       []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:         2,
			expect:      5,
			expectArray: []int{0, 1, 3, 0, 4, 99, 99, 99},
		},
	}

	for _, tt := range testcases {
		actual := removeElement(tt.input, tt.val)
		if actual != tt.expect {
			t.Errorf("[%s(check result)] act: %d, exp: %d\n", tt.msg, actual, tt.expect)
		}
		if !matchIntArray(tt.input, tt.expectArray) {
			t.Errorf("[%s(check array)] act: %v, exp: %v\n", tt.msg, tt.input, tt.expectArray)
		}
	}
}
