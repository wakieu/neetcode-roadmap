package binarysearch

// https://leetcode.com/problems/koko-eating-bananas/

import "math"

func MinEatingSpeed(piles []int, h int) int {
	l, r := 1, max(piles)
	var middle, hrs int

	for l < r {
		middle = (l + r) / 2
		hrs = 0

		for _, pile := range piles {
			hrs += int(math.Ceil(float64(pile) / float64(middle)))
		}

		if hrs <= h {
			r = middle
		} else {
			l = middle + 1
		}
	}
	return r
}

func max(arr []int) int {
	max := arr[0]
	for _, e := range arr {
		if max < e {
			max = e
		}
	}
	return max
}
