package binarysearch

// https://leetcode.com/problems/search-a-2d-matrix/

func SearchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1
	var middle int
	targetRow := -1
	for l <= r {
		middle = (l + r) / 2
		rowSize := len(matrix[middle]) - 1
		if target < matrix[middle][0] {
			r = middle - 1
		} else if target > matrix[middle][rowSize] {
			l = middle + 1
		} else {
			targetRow = middle
			break
		}
	}

	if targetRow == -1 {
		return false
	}

	l, r = 0, len(matrix[targetRow])-1
	for l <= r {
		middle = (l + r) / 2
		if target < matrix[targetRow][middle] {
			r = middle - 1
		} else if target > matrix[targetRow][middle] {
			l = middle + 1
		} else {
			return true
		}
	}

	return false
}
