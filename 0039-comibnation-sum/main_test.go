package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_compinationSum(t *testing.T) {
	testcases := []struct {
		name       string
		candidates []int
		target     int
		expect     [][]int
	}{
		{
			name:       "first example",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expect:     [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "2nd example",
			candidates: []int{2, 3, 5},
			target:     8,
			expect:     [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "3rd example",
			candidates: []int{2},
			target:     1,
			expect:     [][]int{},
		},
		{
			name:       "failed 1",
			candidates: []int{8, 7, 4, 3},
			target:     11,
			expect:     [][]int{{8, 3}, {7, 4}, {4, 4, 3}},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := combinationSum(tc.candidates, tc.target)
			assert.Equal(t, len(tc.expect), len(actual))
			for i, exp := range tc.expect {
				assert.Exactly(t, exp, actual[i])
			}
		})
	}
}
