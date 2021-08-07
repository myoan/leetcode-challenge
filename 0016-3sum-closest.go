package main

import (
	"fmt"
	"math"
)

func threeSumClosest_filter(a []int) []int {
	if len(a) < 3 {
		return a
	}
	var ret []int
	i := 0
	for {
		if i >= len(a) {
			break
		}
		if i+2 < len(a) && a[i] == a[i+1] && a[i] == a[i+2] {
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

func threeSumClosest(nums []int, target int) int {
	fmt.Printf("array:  %v\n", nums)
	tmp := mergesort(nums)
	fmt.Printf("sorted: %v\n", tmp)
	sort := threeSumClosest_filter(tmp)
	fmt.Printf("filter: %v\n", sort)

	closest := sort[0] + sort[1] + sort[2]

	for l := 0; l < len(sort)-2; l++ {
		if target > 0 && sort[l] > target {
			break
		}
		for m := l + 1; m < len(sort)-1; m++ {
			for n := m + 1; n < len(sort); n++ {
				cand := sort[l] + sort[m] + sort[n]
				fmt.Printf("cond: %d(diff: %f), closest: %d(diff: %f)\n", cand, math.Abs(float64(target-cand)), closest, math.Abs(float64(target-closest)))
				if math.Abs(float64(target-cand)) < math.Abs(float64(target-closest)) {
					closest = cand
				}
			}
		}
	}

	return closest
}
