package main

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -1 * n
}

func divideTwoIntegers(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	ret := 0

	de := abs(dividend)
	ds := abs(divisor)

	tmp := de
	for {
		if tmp < ds {
			if (dividend >= 0 && divisor >= 0) || (dividend <= 0 && divisor <= 0) {
				if ret > 2147483647 {
					return 2147483647
				}
				return ret
			}
			if -1*ret < -2147483648 {
				return -2147483648
			}
			return -1 * ret
		}
		tmp -= ds
		ret++
	}
}
