package arraysandhashing

// https://leetcode.com/problems/product-of-array-except-self/

func ProductExceptSelf(nums []int) []int {
	product := 1
	isZero := make(map[int]bool)
	for i, e := range nums {
		if e != 0 {
			product *= e
		} else {
			isZero[i] = true
		}
	}

	if len(isZero) > 1 {
		return make([]int, len(nums))
	}

	var productArray []int

	for i, e := range nums {
		if isZero[i] {
			productArray = append(productArray, product)
		} else if len(isZero) > 0 {
			productArray = append(productArray, 0)
		} else {
			productArray = append(productArray, product/e)
		}
	}
	return productArray
}
