package main

func romanToInt(s string) int {
	l := len(s)
	ret := 0
	for i, c := range s {
		switch c {
		case 'I':
			if i+1 < l && (s[i+1] == 'V' || s[i+1] == 'X') {
				ret -= 1
			} else {
				ret += 1
			}
		case 'X':
			if i+1 < l && (s[i+1] == 'L' || s[i+1] == 'C') {
				ret -= 10
			} else {
				ret += 10
			}
		case 'C':
			if i+1 < l && (s[i+1] == 'D' || s[i+1] == 'M') {
				ret -= 100
			} else {
				ret += 100
			}
		case 'M':
			ret += 1000
		case 'V':
			ret += 5
		case 'L':
			ret += 50
		case 'D':
			ret += 500
		}
	}
	return ret
}
