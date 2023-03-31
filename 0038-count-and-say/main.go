package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	ret := []byte{}
	str := countAndSay(n - 1)
	num := 0
	count := 0
	for i := 0; i < len(str); i++ {
		n, _ := strconv.Atoi(string(str[i]))
		if i == 0 {
			num = n
		}
		if num != n {
			ret = append(ret, byte(count+48))
			ret = append(ret, byte(num+48))
			num = n
			count = 0
		}
		count += 1
	}
	ret = append(ret, byte(count+48))
	ret = append(ret, byte(num+48))
	return string(ret)
}

func main() {
	fmt.Println(countAndSay(1))
}
