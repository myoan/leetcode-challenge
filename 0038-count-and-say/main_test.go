package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countAndSay(t *testing.T) {
	testcases := []struct {
		msg string
		input int
		expect string
	} {
		{
			msg: "1",
			input: 1,
			expect: "1",
		},
		{
			msg: "2",
			input: 2,
			expect: "11",
		},
		{
			msg: "3",
			input: 3,
			expect: "21",
		},
		{
			msg: "4",
			input: 4,
			expect: "1211",
		},
		{
			msg: "8",
			input: 8,
			expect: "1113213211",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.msg, func(t *testing.T) {
			actual := countAndSay(tc.input)
			assert.Equal(t, tc.expect, actual)
		})
	}
}
