package problems

import (
	"slices"
)

// GroupAnagrams solves the leetcode problem 49. Group Anagrams
func GroupAnagrams(strs []string) [][]string {
	buckets := make(map[string][]string)

	for _, word := range strs {
		sortedWord := []rune(word)
		slices.Sort(sortedWord)

		sStr := string(sortedWord)
		if _, ok := buckets[sStr]; !ok {
			buckets[sStr] = []string{word}
		} else {
			buckets[sStr] = append(buckets[sStr], word)
		}
	}

	results := make([][]string, 0)
	for _, bucket := range buckets {
		results = append(results, bucket)
	}

	return results
}
