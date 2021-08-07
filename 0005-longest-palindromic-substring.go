package main

var (
	dummy = "#"
)

func insertDummy(s string) string {
	ret := dummy
	for _, c := range s {
		ret += string(c)
		ret += dummy
	}
	return ret
}

func deleteDummy(s string) string {
	ret := ""
	for _, c := range s {
		if string(c) == dummy {
			continue
		}
		ret += string(c)
	}
	return ret
}

func getPalindrome(s string, l, r int) string {
	ret := ""
	if l == r {
		ret = string(s[l])
	} else {
		ret = s[l : r+1]
	}
	// fmt.Printf("paradrome: '%s', s: '%s', l: %d, r: %d\n", ret, s, l, r)
	return ret
}

func manacher(s string) string {
	center := 0
	edge := 0
	result := ""
	for {
		if center+edge >= len(s) {
			break
		}
		for {
			left := center - edge
			right := center + edge
			// fmt.Printf("center: %d, edge: %d, left: %d, right: %d\n", center, edge, left, right)

			edge++
			if left < 0 || right >= len(s) {
				break
			}
			if s[left] != s[right] {
				break
			}

			p := getPalindrome(s, left, right)
			if len(result) < len(p) {
				result = p
			}
		}
		center++
		edge = 0
	}
	return result
}

func longestPalindrome(s string) string {
	target := insertDummy(s)
	result := manacher(target)
	// fmt.Printf("s: %s, inserted: %s, palindrome: %s\n", s, target, result)
	return deleteDummy(result)
}
