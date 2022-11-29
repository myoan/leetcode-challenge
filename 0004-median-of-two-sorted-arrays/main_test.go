package main

import "testing"

func Test_mergeArray(t *testing.T) {
	testcases := []struct {
		msg    string
		input1 []int
		input2 []int
		expect []int
	}{
		{
			"both empty",
			[]int{},
			[]int{},
			[]int{},
		},
		{
			"one is empty",
			[]int{},
			[]int{1},
			[]int{1},
		},
		{
			"example 1",
			[]int{1, 3},
			[]int{2},
			[]int{1, 2, 3},
		},
		{
			"example 2",
			[]int{1, 2, 3, 4, 5, 6},
			[]int{7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			"same number",
			[]int{1, 2, 3},
			[]int{2},
			[]int{1, 2, 2, 3},
		},
	}

	for _, tt := range testcases {
		actual := mergeArray(tt.input1, tt.input2)
		if !matchSlice(actual, tt.expect) {
			t.Errorf("[%s] act: %v, exp: %v\n", tt.msg, actual, tt.expect)
		}
	}
}

func Test_getCenter(t *testing.T) {
	testcases := []struct {
		msg    string
		input1 []int
		input2 []int
		expect []int
	}{
		{
			"both empty",
			[]int{},
			[]int{},
			[]int{},
		},
		{
			"one is empty",
			[]int{},
			[]int{1},
			[]int{1},
		},
		{
			"same length",
			[]int{2},
			[]int{1},
			[]int{1, 2},
		},
		{
			"example 1",
			[]int{1, 3},
			[]int{2},
			[]int{2},
		},
		{
			"example 2",
			[]int{1, 2, 3, 4, 5, 6},
			[]int{7},
			[]int{4},
		},
		{
			"same number",
			[]int{1, 2, 3},
			[]int{2},
			[]int{2, 2},
		},
	}

	for _, tt := range testcases {
		actual := getCenter(tt.input1, tt.input2)
		if !matchSlice(actual, tt.expect) {
			t.Errorf("[%s] act: %v, exp: %v\n", tt.msg, actual, tt.expect)
		}
	}
}

func Test_findMedianSortedArrays(t *testing.T) {
	testcases := []struct {
		msg    string
		input1 []int
		input2 []int
		expect float64
	}{
		{
			"empty",
			[]int{},
			[]int{},
			0,
		},
		{
			"example 1",
			[]int{1, 3},
			[]int{2},
			2,
		},
		{
			"example 2",
			[]int{1, 2},
			[]int{3, 4},
			2.50000,
		},
		{
			"example 3",
			[]int{0, 0},
			[]int{0, 0},
			0,
		},
		{
			"example 4",
			[]int{},
			[]int{1},
			1,
		},
		{
			"example 5",
			[]int{2},
			[]int{},
			2,
		},
	}

	for _, tt := range testcases {
		actual := findMedianSortedArrays(tt.input1, tt.input2)
		if actual != tt.expect {
			t.Errorf("[%s] act: %f, exp: %f\n", tt.msg, actual, tt.expect)
		}
	}
}
