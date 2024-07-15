package problems

// JumpGame solves the leet code problem 55
func JumpGame(nums []int) bool {
	r := nums[0]

	for _, v := range nums[:len(nums)-1] {
		if r < v {
			r = v
		}

		if r <= 0 {
			return false
		}

		r--
	}

	return true
}

func JumpGameSolution2(nums []int) bool {
	goal := len(nums) - 1

	for i := goal - 1; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}

	return goal == 0
}
