package twopointers

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

func TwoSum(numbers []int, target int) []int {
    r := len(numbers) - 1
	l := 0

	for {
		if r <= l {
			break
		}
		if numbers[l] + numbers[r] > target {
			r--
			continue
		}
		if numbers[l] + numbers[r] < target {
			l++
			continue
		}
		return []int{l+1, r+1}
	}

    return []int{0, 0}
}