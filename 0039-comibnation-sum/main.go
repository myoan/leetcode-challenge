package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) (ret [][]int) {
	if len(candidates) == 0 {
		return ret
	}

	var backtrack func(cur []int, sum, idx int)
	backtrack = func(cur []int, sum, idx int) {
		if sum == target {
			a := make([]int, len(cur))
			copy(a, cur)
			ret = append(ret, a)
		} else if sum > target {
			return
		}

		for i := idx; i < len(candidates); i++ {
			cur = append(cur, candidates[i])
			backtrack(cur, sum+candidates[i], i)
			cur = cur[:len(cur)-1]
		}
	}

	backtrack(make([]int, 0), 0, 0)
	return
}

func main() {
	fmt.Println("Hello, world")
}
