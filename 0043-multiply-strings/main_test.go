package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toStr(t *testing.T) {
	testcases := []struct {
		n1     int
		expect string
	}{
		{n1: 1, expect: "1"},
		{n1: 43, expect: "43"},
	}
	for _, tc := range testcases {
		t.Run(tc.expect, func(t *testing.T) {
			actual := toStr(tc.n1)
			assert.Equal(t, tc.expect, actual)
		})
	}
}

func Test_toInt(t *testing.T) {
	testcases := []struct {
		n1     string
		expect int
	}{
		{n1: "1", expect: 1},
		{n1: "4", expect: 4},
		{n1: "43", expect: 43},
		{n1: "043", expect: 43},
		{n1: "123", expect: 123},
		{n1: "456", expect: 456},
	}
	for _, tc := range testcases {
		t.Run(tc.n1, func(t *testing.T) {
			actual := toInt(tc.n1)
			assert.Equal(t, tc.expect, actual)
		})
	}
}

func Test_multiply(t *testing.T) {
	testcases := []struct {
		name   string
		n1     string
		n2     string
		expect string
	}{
		{name: "simple", n1: "1", n2: "2", expect: "2"},
		{name: "simple2", n1: "123", n2: "456", expect: "56088"},
		{name: "over", n1: "498828660196", n2: "840477629533", expect: "419254329864656431168468"},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := multiply(tc.n1, tc.n2)
			assert.Equal(t, tc.expect, actual)
		})
	}
}
