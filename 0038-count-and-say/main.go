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
	str := countAndSay(n-1)
	num := 0
	count := 0
	for i := 0; i < len(str); i++ {
		n, _ :=  strconv.Atoi(string(str[i]))
		if i == 0 {
			num = n
		}
		if num != n {
			// ret += fmt.Sprintf("%d%d", count, num)
			ret = append(ret, byte(count + 48))
			ret = append(ret, byte(num + 48))
			num = n
			count = 0
		} 
		count += 1
		fmt.Printf("%s[%d]: %d, '%s'\n", str, i, n, ret)
	}
	ret = append(ret, byte(count + 48))
	ret = append(ret, byte(num + 48))
	fmt.Printf("%s: %d, '%s'\n", str, n, ret)
	return string(ret)
}

func main() {
	fmt.Println(countAndSay(1))
}
