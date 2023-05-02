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
			name:       "simple",
			candidates: []int{1},
			target:     1,
			expect:     [][]int{{1}},
		},
		{
			name:       "simple add",
			candidates: []int{1, 2},
			target:     3,
			expect:     [][]int{{1, 2}},
		},
		{
			name:       "first example",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expect:     [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			name:       "2nd example",
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expect:     [][]int{{1, 2, 2}, {5}},
		},
		{
			name: "time exceed",
			candidates: []int{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			target: 30,
			expect: [][]int{
				{
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				},
			},
		},
		{
			name:       "time exceed2",
			candidates: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			target:     30,
			expect: [][]int{
				{
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				},
				{
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1, 2,
				},
			},
		},
		{
			name:       "sample",
			candidates: []int{1, 2, 3, 4},
			target:     30,
			expect:     [][]int{{1, 1, 1}},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := combinationSum(tc.candidates, tc.target)
			if ok := assert.Equal(t, len(tc.expect), len(actual)); !ok {
				t.Errorf("actual: %v, expect: %v\n", actual, tc.expect)
			}
			for i, exp := range tc.expect {
				assert.Exactly(t, exp, actual[i])
			}
		})
	}
}
