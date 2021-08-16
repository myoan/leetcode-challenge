package main

func longestValidParentheses(s string) int {
	dp := make([]int, len(s)+1)
	for i := 1; i < len(s); i++ {
		k := dp[i]
		if s[i-1] == '(' && s[i] == ')' {
			dp[i+1] = 2 + dp[i-1]
		} else if s[i-1] == ')' && s[i] == ')' {
			if i-k-1 >= 0 && s[i-k-1] == '(' {
				dp[i+1] = 2 + dp[i] + dp[i-k-1]
			} else {
				dp[i+1] = 0
			}
		}
	}

	max := 0
	for _, i := range dp {
		if max < i {
			max = i
		}
	}
	return max
}
