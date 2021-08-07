package main

import (
	"fmt"
)

func fourSum_filter(a []int) []int {
	if len(a) < 4 {
		return a
	}
	var ret []int
	i := 0
	for {
		if i >= len(a) {
			break
		}
		if i+3 < len(a) && a[i] == a[i+1] && a[i] == a[i+2] && a[i] == a[i+3] {
			ret = append(ret, a[i])
			ret = append(ret, a[i])
			ret = append(ret, a[i])
			j := i + 1
			for {
				if j >= len(a) || a[i] != a[j] {
					break
				}
				j++
			}
			i = j - 1
		}
		ret = append(ret, a[i])
		i++
	}
	return ret
}

func fourSum(nums []int, target int) [][]int {
	var ret [][]int
	if len(nums) < 4 {
		return [][]int{}
	}
	fmt.Printf("array:  %v\n", nums)
	tmp := mergesort(nums)
	fmt.Printf("sorted: %v\n", tmp)
	sort := fourSum_filter(tmp)
	fmt.Printf("filter: %v\n", sort)

	for l := 0; l < len(sort)-3; l++ {
		if target > 0 && sort[l] > target {
			break
		}
		for m := l + 1; m < len(sort)-2; m++ {
			i := m + 1
			j := len(sort) - 1
			for {
				if j <= i {
					break
				}
				// fmt.Printf("try: %d + %d + %d + %d == %d?\n", sort[l], sort[m], sort[i], sort[j], target)
				sum := sort[l] + sort[m] + sort[i] + sort[j]
				if sum == target {
					ret = append(ret, []int{sort[l], sort[m], sort[i], sort[j]})
				}
				if sum > target {
					j--
				} else {
					i++
				}
			}
		}
	}

	result := make([][]int, 0)
	for _, r := range ret {
		uniq := true
		for _, rr := range result {
			if r[0] == r[0] && r[1] == rr[1] && r[2] == rr[2] && r[3] == rr[3] {
				uniq = false
			}
		}
		if uniq {
			result = append(result, r)
		}
	}
	return result
}
