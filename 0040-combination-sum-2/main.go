package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) (ret [][]int) {
	sort.Ints(candidates)
	var backtrack func([]int, int, int)
	backtrack = func(cur []int, idx, sum int) {
		fmt.Printf("cur: %v\n", cur)
		if sum == target {
			rec := make([]int, len(cur))
			copy(rec, cur)
			ret = append(ret, rec)
			return
		}
		if sum > target {
			return
		}
		for i := idx; i < len(candidates); i++ {
			cur = append(cur, candidates[i])
			backtrack(cur, i+1, sum+candidates[i])
			cur = cur[:len(cur)-1]

			for i < len(candidates)-1 && candidates[i] == candidates[i+1] {
				i++
			}
		}
	}
	backtrack(make([]int, 0), 0, 0)
	return ret
}
