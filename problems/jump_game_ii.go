package problems

// JumpGameII solves the leet code problem 45
func JumpGameII(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	r := 1
	i := 0
	for {
		maxIdx := i + nums[i]
		if maxIdx >= len(nums)-1 {
			break
		}

		nextIdx := maxIdx
		nextJump := nums[maxIdx] + maxIdx
		for j := maxIdx - 1; j > i; j-- {
			candidate := nums[j] + j
			if candidate > nextJump {
				nextJump = candidate
				nextIdx = j
			}
		}

		i = nextIdx
		r++
	}

	return r
}
