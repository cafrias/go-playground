package problems

// LongestConsecutive solves LeetCode problem "128. Longest Consecutive Sequence"
func LongestConsecutive(nums []int) int {
	results := make(map[int]uint32, len(nums))
	for _, v := range nums {
		results[v] = 0
	}

	for i := range results {
		var count uint32 = 1
		for {
			nIdx := i + int(count)
			nVal, nOk := results[nIdx]
			if !nOk {
				break
			}

			if nVal > 0 {
				count += nVal
				break
			}

			results[nIdx] = 1
			count += 1
		}
		results[i] = count
	}

	var largest uint32 = 0
	for _, v := range results {
		if v > largest {
			largest = v
		}
	}

	return int(largest)
}

func LongestConsecutive2(nums []int) int {
	results := make(map[int]bool, len(nums))
	for _, v := range nums {
		results[v] = true
	}

	longest := 0
	for i := range results {
		ok := results[i-1]
		if !ok {
			c := 1
			for {
				ok := results[i+c]
				if !ok {
					break
				}
				c += 1
			}

			if c > longest {
				longest = c
			}
		}
	}

	return longest
}
