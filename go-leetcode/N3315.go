package main

func minBitwiseArrayV2(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = -1

		bits := toBinary(nums[i])

		for j := len(bits) - 1; j >= 0; j-- {
			bit := bits[j]
			if bit == 0 {
				continue
			}
			if j == len(bits)-1 && j-1 > 0 && bits[j-1] == 0 {
				continue
			}

			ret := nums[i] & ^(1 << j)
			if ret|(ret+1) == nums[i] {
				ans[i] = ret
				break
			}
		}
	}
	return ans
}

func toBinary(num int) []int {
	var ans []int
	for num > 0 {
		ans = append(ans, num&1)
		num >>= 1
	}
	return ans
}
