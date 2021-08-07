package main

import (
	"testing"
)

func matchArray(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, e := range a {
		if e != b[i] {
			return false
		}
	}
	return true
}

func Test_letterCombinations(t *testing.T) {
	testcases := []struct {
		msg    string
		input  string
		expect []string
	}{
		{
			"example 1",
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			"example 2",
			"223",
			[]string{
				"aad", "aae", "aaf", "abd", "abe", "abf", "acd", "ace", "acf",
				"bad", "bae", "baf", "bbd", "bbe", "bbf", "bcd", "bce", "bcf",
				"cad", "cae", "caf", "cbd", "cbe", "cbf", "ccd", "cce", "ccf",
			},
		},
	}

	for _, tt := range testcases {
		actual := letterCombinations(tt.input)
		if !matchArray(actual, tt.expect) {
			t.Errorf("[%s] act: %v, exp: %v\n", tt.msg, actual, tt.expect)
		}
	}
}
