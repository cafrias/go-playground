package problems

import (
	"slices"
)

func RotateArray(nums []int, k int) {
	rKey := k % len(nums)

	slices.Reverse(nums)
	slices.Reverse(nums[:rKey])
	slices.Reverse(nums[rKey:])
}
