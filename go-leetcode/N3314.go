package main

func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = -1
		for j := 1; j < nums[i]; j++ {
			if j|(j+1) == nums[i] {
				ans[i] = j
				break
			}
		}
	}
	return ans
}
