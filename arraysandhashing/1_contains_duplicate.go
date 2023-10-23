package arraysandhashing

// https://leetcode.com/problems/contains-duplicate/

func ContainsDuplicate(nums []int) bool {
    m := make(map[int]int)
    for _, e := range nums {
        _, ok := m[e]
        if ok {
            return true
        } else {
            m[e] = 1
        }
    }
    return false
}