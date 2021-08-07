package main

import (
	"testing"
)

func Test_fourSum(t *testing.T) {
	testcases := []struct {
		msg    string
		input  []int
		target int
		expect [][]int
	}{
		{
			"example 1",
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			"example 2",
			[]int{2, 2, 2, 2, 2},
			8,
			[][]int{
				{2, 2, 2, 2},
			},
		},
	}

	for _, tt := range testcases {
		actual := fourSum(tt.input, tt.target)
		if !matchIntArrayOfArray(actual, tt.expect) {
			t.Errorf("[%s] act: '%v', exp: '%v'\n", tt.msg, actual, tt.expect)
		}
	}
}
