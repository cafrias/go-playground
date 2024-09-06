package problems

import (
	"slices"
)

// MergeIntervals Solved LeetCode Problem 56
func MergeIntervals(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a []int, b []int) int {
		return a[0] - b[0]
	})

	res := make([][]int, 0)
	cStart := intervals[0][0]
	cEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		iStart := intervals[i][0]
		iEnd := intervals[i][1]

		if cEnd >= iStart {
			// THey overlap
			cEnd = max(iEnd, cEnd)
		} else {
			// They don't overlap, commit current to intervals
			res = append(res, []int{cStart, cEnd})
			cStart = iStart
			cEnd = iEnd
		}
	}

	res = append(res, []int{cStart, cEnd})

	return res
}
