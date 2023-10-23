package twopointers

// https://leetcode.com/problems/3sum/

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int{}

	for i, e := range nums {
		if i > 0 && e == nums[i-1] {
			continue
		}
		r := len(nums) - 1
		l := i + 1
		for l < r {
			if nums[l] + nums[r] + e > 0  {
				r--
				continue
			}
			if nums[l] + nums[r] + e < 0 {
				l++
				continue
			}
			res = append(res, []int{e, nums[l], nums[r]})
			for l < r && nums[l] == nums[l+1] {
				l++
			}
			for l < r && nums[l] == nums[r-1] {
				r--
			}
			r--
			l++
		}
	}
	return res
}
