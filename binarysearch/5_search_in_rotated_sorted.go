package binarysearch

// https://leetcode.com/problems/search-in-rotated-sorted-array/

func SearchInRotatedSorted(nums []int, target int) int {

	// find array min element in log(n) time, split the array in that spot and binary search the correct slice

	minIndex := 0

	l, r := 0, len(nums)-1
	n := r

	for l <= r {
		if nums[l] < nums[r] {
			minIndex = l
			break
		}
		middle := (l + r) / 2
		if nums[middle] > nums[l] {
			l = middle + 1
		} else if nums[middle] < nums[l] {
			r = middle
		} else {
			minIndex = r
			break
		}
	}

	if target <= nums[n] {
		aux := binarySearch(nums[minIndex:], target)
		if aux == -1 {
			return -1
		} else {
			return len(nums[:minIndex]) + aux
		}
	} else {
		return binarySearch(nums[:minIndex], target)
	}
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	var middle int
	for l <= r {
		middle = (l + r) / 2
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
