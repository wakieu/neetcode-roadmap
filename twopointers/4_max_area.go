package twopointers

// https://leetcode.com/problems/container-with-most-water/

func MaxArea(height []int) int {
	l := 0
	r := len(height) - 1

	max := 0

	for l < r {
		h := 0
		if height[l] < height[r] {
			h = height[l]
		} else {
			h = height[r]
		}
		width := r - l
		area := h * width

		if area > max {
			max = area
		}

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return max
}
