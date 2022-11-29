package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	testcases := []struct {
		msg    string
		input  []int
		target int
		expect int
	}{
		{
			msg:    "empty case",
			input:  []int{},
			target: 0,
			expect: -1,
		},
		{
			msg:    "simple case",
			input:  []int{0, 1, 2, 3, 4, 5},
			target: 0,
			expect: 0,
		},
		{
			msg:    "simple case not found target",
			input:  []int{1, 2, 3, 4, 6},
			target: 0,
			expect: -1,
		},
		{
			msg:    "simple case not found target",
			input:  []int{1, 2, 3, 4, 5},
			target: 0,
			expect: -1,
		},
		{
			msg:    "no rotation and target is bigger than k",
			input:  []int{-3, -2, -1, 0, 1, 2, 3, 4, 5},
			target: 4,
			expect: 7,
		},
		{
			msg:    "no rotation and target is smaller than k",
			input:  []int{-3, -2, -1, 0, 1, 2, 3, 4, 5},
			target: -2,
			expect: 1,
		},
		{
			msg:    "no rotation and target is head of nums",
			input:  []int{-3, -2, -1, 0, 1, 2, 3, 4, 5},
			target: -3,
			expect: 0,
		},
		{
			msg:    "no rotation and target is tail of nums",
			input:  []int{-3, -2, -1, 0, 1, 2, 3, 4, 5},
			target: 5,
			expect: 8,
		},
		{
			msg:    "no rotation and target not found",
			input:  []int{-4, -2, -1, 0, 1, 2, 3, 4, 5},
			target: -3,
			expect: -1,
		},
		{
			msg:    "rotation",
			input:  []int{100, -100, 0, 1, 2, 3, 4, 5},
			target: 0,
			expect: 2,
		},
		{
			msg:    "example 1",
			input:  []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			expect: 4,
		},
	}
	for _, tt := range testcases {
		actual := search(tt.input, tt.target)
		assert.Equal(t, tt.expect, actual, tt.msg)
	}
}
