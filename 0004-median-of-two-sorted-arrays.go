package main

import "fmt"

func mergeArray(nums1 []int, nums2 []int) []int {
	len1 := len(nums1)
	len2 := len(nums2)
	length := len1 + len2
	j1 := 0
	j2 := 0
	ret := make([]int, length)

	for i := 0; i < length; i++ {
		if len1 > j1 && len2 > j2 {
			if nums1[j1] > nums2[j2] {
				ret[i] = nums2[j2]
				j2++
			} else {
				ret[i] = nums1[j1]
				j1++
			}
		} else if len1 > j1 {
			ret[i] = nums1[j1]
			j1++
		} else if len2 > j2 {
			ret[i] = nums2[j2]
			j2++
		} else {
			fmt.Println("else!!")
		}
	}
	return ret
}

func getCenter(nums1 []int, nums2 []int) []int {
	a := mergeArray(nums1, nums2)
	len1 := len(nums1)
	len2 := len(nums2)
	length := len1 + len2

	if length == 0 {
		return []int{}
	}

	center := length / 2
	if length%2 == 0 {
		return []int{a[center-1], a[center]}
	}
	return []int{a[center]}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	center := getCenter(nums1, nums2)
	if len(center) == 0 {
		return 0
	}
	sum := 0
	for _, e := range center {
		sum += e
	}

	return float64(sum) / float64(len(center))
}
