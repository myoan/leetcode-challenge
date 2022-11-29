package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestValidParentheses(t *testing.T) {
	testcases := []struct {
		msg    string
		input  string
		expect int
	}{
		{
			msg:    "empty case",
			input:  "",
			expect: 0,
		},
		{
			msg:    "simple case",
			input:  "()",
			expect: 2,
		},
		{
			msg:    "simple case 2",
			input:  "()()",
			expect: 4,
		},
		{
			msg:    "simple nest",
			input:  "(())",
			expect: 4,
		},
		{
			msg:    "simple nest 2",
			input:  "(()())",
			expect: 6,
		},
		{
			msg:    "complext example",
			input:  "()(()()))",
			expect: 8,
		},
		{
			msg:    "example 1",
			input:  "(()",
			expect: 2,
		},
		{
			msg:    "example 2",
			input:  ")())())",
			expect: 2,
		},
		{
			msg:    "less close parens",
			input:  "()((()",
			expect: 2,
		},
	}
	for _, tt := range testcases {
		actual := longestValidParentheses(tt.input)
		assert.Equal(t, tt.expect, actual, tt.msg)
	}
}
