package stackk

// https://leetcode.com/problems/daily-temperatures/

func DailyTemperatures(temperatures []int) []int {

	// start result array
	result := make([]int, len(temperatures))

	// go through it in reverse
	for i := len(temperatures) - 1; i >= 0; i-- {
		// j in one position ahead od i
		j := i + 1

		// if temp[j] is lower/equal than temp[i] -> find a higher temp ahead
		for j < len(temperatures) && temperatures[j] <= temperatures[i] {
			// if result[j] == 0 , this means there is no higher temp than temp[j] ahead break
			if result[j] <= 0 {
				break
			}
			// walk j by result[j]
			j += result[j]
		}

		// if temp[j] high than temp[i] -> result[i] is their index diference
		if j < len(temperatures) && temperatures[j] > temperatures[i] {
			result[i] = j - i
		} // else the result already has 0

	}
	return result
}
