package arraysandhashing

// https://leetcode.com/problems/top-k-frequent-elements/

import "sort"

func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, e := range nums {
		m[e]++
	}

	keys := make([]int, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return keys[:k]
}
