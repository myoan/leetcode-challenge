package main

import "fmt"

type VirtualArray struct {
	pivot int
	raw   []int
}

func NewVArray(array []int) *VirtualArray {
	pivot := 0
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			pivot = i
			break
		}
	}

	return &VirtualArray{
		pivot: pivot,
		raw:   array,
	}
}

func (va *VirtualArray) Get(i int) int {
	return va.raw[va.ShiftIdx(i)]
}

func (va *VirtualArray) Show() {
	fmt.Printf("[ %d", va.raw[va.ShiftIdx(0)])
	for i := 1; i < len(va.raw); i++ {
		fmt.Printf(", %d", va.raw[va.ShiftIdx(i)])
	}
	fmt.Println(" ]")
}

func (va *VirtualArray) ShiftIdx(i int) int {
	return (i + va.pivot) % len(va.raw)
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	a := NewVArray(nums)
	i := 0
	j := len(nums) - 1
	k := (i + j) / 2
	// a.Show()

	if a.Get(i) == target {
		return a.ShiftIdx(i)
	} else if a.Get(j) == target {
		return a.ShiftIdx(j)
	}

	for {
		// fmt.Printf("- i: %d, k: %d, j: %d\n", i, k, j)
		if i == k || j == k {
			break
		}

		if a.Get(k) == target {
			return a.ShiftIdx(k)
		} else if a.Get(k) < target {
			i = k
		} else {
			j = k
		}
		k = (i + j) / 2
	}
	return -1
}
