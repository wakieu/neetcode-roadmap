package arraysandhashing

// https://leetcode.com/problems/group-anagrams/

import "sort"

func GroupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, e := range strs {
		rs := []rune(e)
		sort.Slice(rs, func(i, j int) bool {
			return rs[i] < rs[j]
		})
		hash := string(rs)
		m[hash] = append(m[hash], e)
	}

	var r [][]string

	for _, e := range m {
		r = append(r, e)
	}

	return r
}
