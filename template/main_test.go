package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Hoge(t *testing.T) {
	testcases := []struct {
		name   string
		input  int
		expect int
	}{
		{
			name:   "hoge",
			input:  1,
			expect: 1,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var actual int
			// YOUR CODE HERE
			assert.Equal(t, tc.expect, actual)
		})
	}
}
