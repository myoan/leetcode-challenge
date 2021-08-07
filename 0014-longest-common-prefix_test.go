package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	testcases := []struct {
		msg    string
		inputs []string
		expect string
	}{
		{
			"example 1",
			[]string{"flower", "flow", "flight"},
			"fl",
		},
		{
			"example 2",
			[]string{"dog", "racecar", "car"},
			"",
		},
		{
			"example 3",
			[]string{"cir", "car"},
			"c",
		},
	}

	for _, tt := range testcases {
		actual := longestCommonPrefix(tt.inputs)
		if actual != tt.expect {
			t.Errorf("[%s] act: '%s', exp: '%s'\n", tt.msg, actual, tt.expect)
		}
	}
}
