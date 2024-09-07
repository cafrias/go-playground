package problems

import (
	"slices"
)

// MinArrowsBalloons solves LeetCode Problem 452. Minimum Number of Arrows to Burst Balloons
func MinArrowsBalloons(points [][]int) int {
	slices.SortFunc(points, func(a []int, b []int) int {
		return a[0] - b[0]
	})

	arrows := 1
	arrowInterval := points[0]

	for i := 1; i < len(points); i++ {
		if points[i][0] <= arrowInterval[1] && points[i][1] >= arrowInterval[0] {
			// They intersect, we can use the same arrow, but the interval for the arrow
			// gets reduced
			arrowInterval[0] = max(points[i][0], arrowInterval[0])
			arrowInterval[1] = min(points[i][1], arrowInterval[1])
		} else {
			arrowInterval = points[i]
			arrows++
		}
	}

	return arrows
}
