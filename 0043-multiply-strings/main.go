package main

import "fmt"

func toInt(str string) int {
	ret := 0
	for i := 0; i < len(str); i++ {
		e := str[i]
		n := int(byte(e)) - 48
		dig := len(str) - i
		if dig > 0 {
			ret *= 10
		}
		ret += n
		// fmt.Printf("ret %d, i: %d, e: %s, n: %d\n", ret, i, string(e), n)
	}
	return ret
}

func toStr(n int) string {
	var tmp []byte
	i := 0
	for n > 0 {
		e := n % 10
		tmp = append(tmp, byte(e+48))
		n /= 10
		i++
	}

	if len(tmp) == 0 {
		return "0"
	}

	// reverse
	ret := make([]byte, len(tmp))
	for i := len(tmp) - 1; i >= 0; i-- {
		j := len(tmp) - i - 1
		ret[j] = tmp[i]
	}
	return string(ret)
}

func multiply(num1 string, num2 string) string {
	return toStr(toInt(num1) * toInt(num2))
}

func main() {
	fmt.Printf("hoge")
}
