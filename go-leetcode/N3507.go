package main

func minimumPairRemoval(nums []int) int {
	var ops int
	for !checkSequence(nums) {
		var ret int
		var retIdx int
		for i := 0; i < len(nums)-1; i++ {
			sum := nums[i] + nums[i+1]
			if i == 0 || sum < ret {
				ret = sum
				retIdx = i
			}
		}
		numsNew := append(nums[:retIdx], ret)
		if retIdx+2 < len(nums) {
			numsNew = append(numsNew, nums[retIdx+2:]...)
		}
		nums = numsNew
		ops++
	}
	return ops
}

func checkSequence(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
