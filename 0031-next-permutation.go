package main

func swap(nums []int, r, l int) {
	tmp := nums[r]
	nums[r] = nums[l]
	nums[l] = tmp
}

func sort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				swap(nums, i, j)
			}
		}
	}
}

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			idx := i + 1
			for j := i + 1; j < len(nums); j++ {
				if nums[j] < nums[idx] && nums[j] > nums[i] {
					idx = j
				}
			}
			swap(nums, i, idx)
			sort(nums[i+1:])
			return
		}
	}
	sort(nums)
}
