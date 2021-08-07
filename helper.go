package main

func mergesort(a []int) []int {
	ln := len(a)
	if ln == 1 {
		return a
	}
	half := ln / 2
	sort1 := mergesort(a[0:half])
	sort2 := mergesort(a[half:])
	ret := make([]int, ln)
	i := 0
	j := 0
	k := 0
	for {
		if k >= ln {
			break
		}
		if half <= i {
			ret[k] = sort2[j]
			j++
		} else if (ln - half) <= j {
			ret[k] = sort1[i]
			i++
		} else {
			if sort1[i] > sort2[j] {
				ret[k] = sort2[j]
				j++
			} else {
				ret[k] = sort1[i]
				i++
			}
		}
		k++
	}
	return ret
}
