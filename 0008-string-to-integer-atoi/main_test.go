package main

import "testing"

func Test_myAtoi(t *testing.T) {
	testcases := []struct {
		msg    string
		input  string
		expect int
	}{
		{
			"example1",
			"42",
			42,
		},
		{
			"example2",
			"   -42",
			-42,
		},
		{
			"hyphen example 1",
			"   4-2",
			4,
		},
		{
			"hyphen example 2",
			"-+42",
			0,
		},
		{
			"hyphen example 3",
			"+-12",
			0,
		},
		{
			"hyphen example 4",
			"42-",
			42,
		},
		{
			"hyphen example 5",
			"+-42",
			0,
		},
		{
			"example 3",
			"4193 with words",
			4193,
		},
		{
			"example 4",
			"words and 987",
			0,
		},
		{
			"example 5",
			"-91283472332",
			-2147483648,
		},
		{
			"bigger",
			"9999999999",
			2147483647,
		},
		{
			"pi",
			"3.14159",
			3,
		},
		{
			"space",
			"   +0 123",
			0,
		},
	}

	for _, tt := range testcases {
		actual := myAtoi(tt.input)
		if actual != tt.expect {
			t.Errorf("[%s] act: %d, exp: %d\n", tt.msg, actual, tt.expect)
		}
	}
}
