package main

import "testing"

func Test_intToRoman(t *testing.T) {
	testcases := []struct {
		msg    string
		input  int
		expect string
	}{
		{
			"case 1",
			1,
			"I",
		},
		{
			"case 2",
			2,
			"II",
		},
		{
			"case 3",
			3,
			"III",
		},
		{
			"case 4",
			4,
			"IV",
		},
		{
			"case 5",
			5,
			"V",
		},
		{
			"case 6",
			6,
			"VI",
		},
		{
			"case 7",
			7,
			"VII",
		},
		{
			"case 8",
			8,
			"VIII",
		},
		{
			"case 9",
			9,
			"IX",
		},
		{
			"case 10",
			10,
			"X",
		},
		{
			"case 27",
			27,
			"XXVII",
		},
		{
			"example 4",
			58,
			"LVIII",
		},
		{
			"example 5",
			1994,
			"MCMXCIV",
		},
	}

	for _, tt := range testcases {
		actual := intToRoman(tt.input)
		if actual != tt.expect {
			t.Errorf("[%s] act: '%s', expecte: '%s'\n", tt.msg, actual, tt.expect)
		}
	}
}
