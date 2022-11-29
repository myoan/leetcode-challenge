package main

func strStr(haystack string, needle string) int {
	needlelen := len(needle)
	if needle == "" {
		return 0
	}
	for i := 0; i <= len(haystack)-needlelen; i++ {
		if haystack[i:i+needlelen] == needle {
			return i
		}
	}
	return -1
}
