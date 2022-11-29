package main

func removeElement(nums []int, val int) int {
	i := 0
	rest := 0
	len := len(nums)
	for {
		if i >= len {
			break
		}
		if nums[i] != val {
			nums[rest] = nums[i]
			rest++
		}
		i++
	}
	for i := rest; i < len; i++ {
		nums[i] = 99
	}
	return rest
}
