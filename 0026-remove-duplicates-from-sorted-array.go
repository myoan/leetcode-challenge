package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	uniq := 1
	i := 1
	for {
		if i >= len(nums) {
			break
		}
		if nums[uniq-1] != nums[i] {
			nums[uniq] = nums[i]
			uniq++
		}
		i++
	}
	return uniq
}
