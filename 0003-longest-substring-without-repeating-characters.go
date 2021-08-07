package main

var (
	actualStr = ""
)

func isInclude(used []byte, target byte, len int) bool {
	for i := 0; i < len; i++ {
		if used[i] == target {
			return true
		}
	}
	return false
}

func getWord(str string, start, end int) string {
	i := 0
	ret := make([]byte, end-start)
	for j := start; j < end; j++ {
		ret[i] = str[j]
		i++
	}
	return string(ret)
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	idx := 0
	used := make([]byte, 95)

	for {
		count := 0
		j := 0
		for i := idx; i < len(s); i++ {
			if isInclude(used, s[i], j) {
				if count > max {
					actualStr = getWord(s, idx, i)
					max = count
				}
				break
				// j = 0
				// count = 0
			}
			used[j] = s[i]
			j++
			count++
		}
		if count > max {
			actualStr = getWord(s, idx, len(s))
			max = count
		}
		if idx > len(s) {
			return max
		}
		if len(s)-idx < max {
			return max
		}
		idx++
	}
}
