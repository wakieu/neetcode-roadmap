package arraysandhashing

// https://leetcode.com/problems/two-sum/

func TwoSum(nums []int, target int) []int {
    m := make(map[int][]int)
    for i, e := range nums {
        m[e] = append(m[e], i)
    }

    for num, indexArray := range m {
        complementarIndexArray, ok := m[target - num]

        if ok {
            if indexArray[0] != complementarIndexArray[len(complementarIndexArray)-1] {
                return []int{indexArray[0], complementarIndexArray[len(complementarIndexArray)-1]}
            }
        }
    }

    return []int{0, 0}
}