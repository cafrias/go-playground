package problems

// Hindex solves the Leet code problem 274. H-Index
func Hindex(citations []int) int {
	count := make([]int, len(citations)+1)

	maxIdx := len(count) - 1
	for _, v := range citations {
		cIdx := min(v, maxIdx)
		count[cIdx]++
	}

	r := 0
	for i := len(count) - 1; i >= 0; i-- {
		c := count[i]
		if r+c >= i {
			return i
		}
		r += c
	}

	return 0
}
