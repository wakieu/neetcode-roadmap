package stackk

// https://leetcode.com/problems/largest-rectangle-in-histogram/

func LargestRectangleArea(heights []int) int {
	largest := 0
	last := 0
	for i, e := range heights {
		if e == 0 || e == last {
			continue
		}
		r := 0
		for j := i + 1; j < len(heights); j++ {
			if heights[j] >= e {
				r += e
			} else {
				break
			}
		}
		l := 0
		for j := i - 1; j >= 0; j-- {
			if heights[j] >= e {
				l += e
			} else {
				break
			}
		}

		if l+e+r > largest {
			largest = l + e + r
		}

		last = e
	}

	return largest
}
