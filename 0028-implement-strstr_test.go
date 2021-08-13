package main

import "testing"

func Test_strStr(t *testing.T) {
	testcases := []struct {
		msg      string
		haystack string
		needle   string
		expect   int
	}{
		{
			msg:      "simple",
			haystack: "abcde",
			needle:   "e",
			expect:   4,
		},
		{
			msg:      "return first match",
			haystack: "helllllo",
			needle:   "lll",
			expect:   2,
		},
		{
			msg:      "example 1",
			haystack: "hello",
			needle:   "ll",
			expect:   2,
		},
		{
			msg:      "example 2",
			haystack: "aaaaa",
			needle:   "bba",
			expect:   -1,
		},
		{
			msg:      "example 3",
			haystack: "",
			needle:   "",
			expect:   0,
		},
	}

	for _, tt := range testcases {
		actual := strStr(tt.haystack, tt.needle)
		if actual != tt.expect {
			t.Errorf("[%s] act: %d, exp: %d\n", tt.msg, actual, tt.expect)
		}
	}
}
