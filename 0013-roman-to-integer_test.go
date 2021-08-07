package main

import "testing"

func Test_romanToInt(t *testing.T) {
	testcases := []struct {
		msg    string
		input  string
		expect int
	}{
		{
			"case 1",
			"I",
			1,
		},
		{
			"case 2",
			"II",
			2,
		},
		{
			"case 4",
			"IV",
			4,
		},
		{
			"case 5",
			"V",
			5,
		},
		{
			"case 8",
			"VIII",
			8,
		},
		{
			"case 9",
			"IX",
			9,
		},
		{
			"case 10",
			"X",
			10,
		},
		{
			"case 48",
			"XLVIII",
			48,
		},
		{
			"case 58",
			"LVIII",
			58,
		},
		{
			"case 1994",
			"MCMXCIV",
			1994,
		},
	}

	for _, tt := range testcases {
		actual := romanToInt(tt.input)
		if actual != tt.expect {
			t.Errorf("[%s] act: '%d', expecte: '%d'\n", tt.msg, actual, tt.expect)
		}
	}
}
