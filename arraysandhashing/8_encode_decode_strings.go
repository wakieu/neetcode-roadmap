package arraysandhashing

// https://leetcode.com/problems/encode-and-decode-strings/
// FREE LINK: https://www.lintcode.com/problem/659/

import "strconv"

func Encode(strings []string) string {
	var s string
	for _, x := range strings {
		s += strconv.Itoa(len(x)) + "|" + x
	}
	return s
}

func Decode(s string) []string {
	var decodedString []string
	var size string

	for len(s) > 0 {
		if s[0] == '|' {
			length, err := strconv.Atoi(size)
			if err != nil {
				panic("int conversion error")
			}
			decodedString = append(decodedString, s[1:length+1])
			s = s[length+1:]
			size = ""
		} else {
			size += string(s[0])
			s = s[1:]
		}
	}

	return decodedString
}
