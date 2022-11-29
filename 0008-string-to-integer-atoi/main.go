package main

import "strconv"

func convert(s string) int {
	ret, _ := strconv.Atoi(s)
	if ret < -2147483648 {
		return -2147483648
	} else if ret > 2147483647 {
		return 2147483647
	}
	return ret
}

func myAtoi(s string) int {
	data := ""
	for i, c := range s {
		switch c {
		case ' ':
			if len(data) != 0 {
				goto L
			}
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			data += string(c)
		case '-':
			if len(data) != 0 {
				goto L
			}
			if i+1 < len(s) && 48 <= int(s[i+1]) && int(s[i+1]) < 58 {
				data += string(c)
			} else {
				goto L
			}
		case '+':
			if len(data) != 0 {
				goto L
			}
			if i+1 < len(s) && 48 <= int(s[i+1]) && int(s[i+1]) < 58 {
				// NOP
			} else {
				goto L
			}
		default:
			if len(data) == 0 {
				return 0
			}
			return convert(data)
		}
	}
L:
	return convert(data)
}
