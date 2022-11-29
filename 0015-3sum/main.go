package main

import (
	"fmt"
)

func threeSum_filter(a []int) []int {
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

func threeSum(a []int) [][]int {
	var ret [][]int
	if len(a) < 3 {
		return [][]int{}
	}
	fmt.Printf("array:  %v\n", a)
	tmp := mergesort(a)
	fmt.Printf("sorted: %v\n", tmp)
	sort := threeSum_filter(tmp)
	fmt.Printf("filter: %v\n", sort)

	for l := 0; l < len(sort)-2; l++ {
		if sort[l] > 0 {
			break
		}
		for m := l + 1; m < len(sort)-1; m++ {
			for n := m + 1; n < len(sort); n++ {
				if sort[l]+sort[m]+sort[n] == 0 {
					ret = append(ret, []int{sort[l], sort[m], sort[n]})
				}
			}
		}
	}

	result := make([][]int, 0)
	for _, r := range ret {
		uniq := true
		for _, rr := range result {
			if r[0] == r[0] && r[1] == rr[1] && r[2] == rr[2] {
				uniq = false
			}
		}
		if uniq {
			result = append(result, r)
		}
	}
	return result
}
