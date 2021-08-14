package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextPermutation(t *testing.T) {
	testcases := []struct {
		msg    string
		input  []int
		expect []int
	}{
		{
			msg:    "empty case",
			input:  []int{},
			expect: []int{},
		},
		{
			msg:    "single number",
			input:  []int{1},
			expect: []int{1},
		},
		{
			msg:    "double number",
			input:  []int{1, 2},
			expect: []int{2, 1},
		},
		{
			msg:    "triple number 1",
			input:  []int{1, 2, 3},
			expect: []int{1, 3, 2},
		},
		{
			msg:    "triple number 2",
			input:  []int{2, 3, 1},
			expect: []int{3, 1, 2},
		},
		{
			msg:    "impossible arrangement",
			input:  []int{3, 2, 1},
			expect: []int{1, 2, 3},
		},
		{
			msg:    "same value",
			input:  []int{1, 1, 1},
			expect: []int{1, 1, 1},
		},
		{
			msg:    "wrong anser",
			input:  []int{1, 3, 2},
			expect: []int{2, 1, 3},
		},
	}

	for _, tt := range testcases {
		nextPermutation(tt.input)
		assert.Equal(t, tt.expect, tt.input, tt.msg)
	}
}
