package arraysandhashing

// https://leetcode.com/problems/valid-anagram/

func IsAnagram(s string, t string) bool {
	m := make(map[rune]int)
	l := 0
	for _, e := range s {
		m[e]++
		l++
	}

	for _, e := range t {
		m[e]--
		if m[e] < 0 {
			return false
		}
		l--
	}

	return l == 0
}