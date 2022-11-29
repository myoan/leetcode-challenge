package main

func reverseNum(x int) int {
	ret := 0
	tmp := x
	for {
		if tmp == 0 {
			return ret
		}
		mod := tmp % 10
		ret = ret*10 + mod
		tmp = tmp / 10
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	r := reverseNum(x)
	return x-r == 0
}
