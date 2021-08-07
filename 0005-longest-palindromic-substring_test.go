package main

import "testing"

func Test_insertDummy(t *testing.T) {
	testcases := []struct {
		msg    string
		input  string
		expect string
	}{
		{
			"simple",
			"a",
			"#a#",
		},
		{
			"even characters",
			"aa",
			"#a#a#",
		},
		{
			"sample",
			"hoge",
			"#h#o#g#e#",
		},
	}

	for _, tt := range testcases {
		actual := insertDummy(tt.input)
		if actual != tt.expect {
			t.Errorf("[%s] act: %s, exp: %s\n", tt.msg, actual, tt.expect)
		}
	}
}

func Test_deleteDummy(t *testing.T) {
	testcases := []struct {
		msg    string
		input  string
		expect string
	}{
		{
			"sample",
			"#h#o#g#e#",
			"hoge",
		},
	}

	for _, tt := range testcases {
		actual := deleteDummy(tt.input)
		if actual != tt.expect {
			t.Errorf("[%s] act: %s, exp: %s\n", tt.msg, actual, tt.expect)
		}
	}
}

func Test_manacher(t *testing.T) {
	testcases := []struct {
		msg    string
		input  string
		expect string
	}{
		{
			"example4 with dummy",
			"#a#c#",
			"#a#",
		},
	}

	for _, tt := range testcases {
		actual := manacher(tt.input)
		if actual != tt.expect {
			t.Errorf("[%s] act: %s, exp: %s\n", tt.msg, actual, tt.expect)
		}
	}
}

func Test_longestPalindrome(t *testing.T) {
	testcases := []struct {
		msg    string
		input  string
		expect string
	}{
		{
			"example1",
			"babad",
			"bab",
		},
		{
			"example2",
			"cbbd",
			"bb",
		},
		{
			"example3",
			"a",
			"a",
		},
		{
			"example4",
			"ac",
			"a",
		},
	}

	for _, tt := range testcases {
		actual := longestPalindrome(tt.input)
		if actual != tt.expect {
			t.Errorf("[%s] act: %s, exp: %s\n", tt.msg, actual, tt.expect)
		}
	}
}
