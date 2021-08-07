package main

func min(strs []string) string {
	min := strs[0]
	for _, s := range strs {
		if len(min) > len(s) {
			min = s
		}
	}
	return min
}

func longestCommonPrefix(strs []string) string {
	minstr := min(strs)
	prefix := ""
	tmp := ""

	for i := 0; i < len(minstr); i++ {
		tmp = tmp + string(minstr[i])
		for _, s := range strs {
			for j := 0; j <= i; j++ {
				if s[j] != tmp[j] {
					return prefix
				}
			}
		}
		prefix = tmp
	}
	return prefix
}
