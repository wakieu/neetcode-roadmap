package binarysearch

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

func FindMin(nums []int) int {
    l, r := 0, len(nums) - 1

    for l <= r {
        if nums[l] < nums[r] {
            return nums[l]
        }
        middle := (l+r)/2
        if nums[middle] > nums[l] {
            l = middle + 1
        } else if nums[middle] < nums[l] {
            r = middle
        } else {
            return nums[r]
        }
    }

    return -1
}