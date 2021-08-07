package main

import (
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	testcases := []struct {
		msg    string
		input  []int
		target int
		expect int
	}{
		{
			"example 1",
			[]int{-1, 2, 1, -4},
			1,
			2,
		},
		{
			"example 2",
			[]int{1, 1, 1, 1},
			0,
			3,
		},
		{
			"example 2",
			[]int{-100, -98, -2, -1},
			-101,
			-101,
		},
	}

	for _, tt := range testcases {
		actual := threeSumClosest(tt.input, tt.target)
		if actual != tt.expect {
			t.Errorf("[%s] act: %d, exp: %d\n", tt.msg, actual, tt.expect)
		}
	}
}
