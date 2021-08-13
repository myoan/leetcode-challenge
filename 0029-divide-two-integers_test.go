package main

import "testing"

func Test_dividendTwoInteger(t *testing.T) {
	testcases := []struct {
		msg      string
		dividend int
		divisor  int
		expect   int
	}{
		{
			msg:      "when dividend is zero",
			dividend: 0,
			divisor:  1,
			expect:   0,
		},
		{
			msg:      "simple case",
			dividend: 4,
			divisor:  2,
			expect:   2,
		},
		{
			msg:      "when dividend is negative",
			dividend: -4,
			divisor:  2,
			expect:   -2,
		},
		{
			msg:      "when divsor is negative",
			dividend: 4,
			divisor:  -2,
			expect:   -2,
		},
		{
			msg:      "when dividend and divsor are negative",
			dividend: -4,
			divisor:  -2,
			expect:   2,
		},
		{
			msg:      "example 1",
			dividend: 10,
			divisor:  3,
			expect:   3,
		},
		{
			msg:      "example 2",
			dividend: 7,
			divisor:  -3,
			expect:   -2,
		},
		{
			msg:      "example 3",
			dividend: 0,
			divisor:  1,
			expect:   0,
		},
		{
			msg:      "example 4",
			dividend: 1,
			divisor:  1,
			expect:   1,
		},
		{
			msg:      "error",
			dividend: -2147483648,
			divisor:  -1,
			expect:   2147483647,
		},
	}

	for _, tt := range testcases {
		actual := divideTwoIntegers(tt.dividend, tt.divisor)
		if actual != tt.expect {
			t.Errorf("[%s(check result)] act: %d, exp: %d\n", tt.msg, actual, tt.expect)
		}
	}
}
