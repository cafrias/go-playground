package problems

// InsertInterval solves LeetCode problem 57. Insert Interval
func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals))

	for i, curr := range intervals {
		if newInterval[1] < curr[0] {
			res = append(res, newInterval)
			return append(res, intervals[i:]...)
		}

		if newInterval[0] > curr[1] {
			res = append(res, curr)
		} else {
			newInterval[0] = min(newInterval[0], curr[0])
			newInterval[1] = max(newInterval[1], curr[1])
		}
	}

	res = append(res, newInterval)

	return res
}
