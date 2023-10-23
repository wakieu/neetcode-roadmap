package binarysearch

// https://leetcode.com/problems/binary-search/

func BinarySearch(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    var middle int
    for l <= r {
        middle = (l+r)/2
        if target < nums[middle] {
            r = middle - 1
        } else if target > nums[middle] {
            l = middle + 1
        } else {
            return middle
        }
    }
    return -1
}
