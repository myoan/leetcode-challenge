package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	testcases := []struct {
		msg    string
		input  int
		expect bool
	}{
		{
			"example1",
			121,
			true,
		},
		{
			"example2",
			-121,
			false,
		},
		{
			"example3",
			10,
			false,
		},
		{
			"example4",
			-101,
			false,
		},
		{
			"even palindrome",
			1001,
			true,
		},
		{
			"odd palindrome",
			101,
			true,
		},
	}

	for _, tt := range testcases {
		actual := isPalindrome(tt.input)
		if actual != tt.expect {
			t.Errorf("[%s]\n", tt.msg)
		}
	}
}
