package arraysandhashing

// https://leetcode.com/problems/longest-consecutive-sequence/

import "sort"

func LongestConsecutive(nums []int) int {

	sort.Ints(nums)

	var prev int
	var seq int
	var longest int

	for i, x := range nums {
		if i == 0 {
			prev = x
			seq++
			continue
		}

		if x-prev == 1 {
			seq++
		} else if x == prev {
			continue
		} else {
			if seq > longest {
				longest = seq
			}
			seq = 1
		}
		prev = x
	}

	if seq > longest {
		longest = seq
	}

	return longest
}
