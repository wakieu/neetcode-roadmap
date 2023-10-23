package twopointers

// https://leetcode.com/problems/trapping-rain-water/

func Trap(height []int) int {
    l := 0
    r := len(height) - 1
    maxL := 0
    maxR := 0
    water := 0
    for l < r {
        if height[l] < height[r] {
            if height[l] >= maxL {
                maxL = height[l]
            } else {
                water += maxL - height[l]
            }
            l++
        } else {
            if height[r] >= maxR {
                maxR = height[r]
            } else {
                water += maxR - height[r]
            }
            r--
        }
    }

    return water
}