package stackk

// https://leetcode.com/problems/valid-parentheses/

func IsValid(s string) bool {
	var st []rune

	for _, x := range s {

		if x == '(' {
			st = append(st, x)
			continue
		}

		if x == '[' {
			st = append(st, x)
			continue
		}

		if x == '{' {
			st = append(st, x)
			continue
		}

		l := len(st) - 1

		if l < 0 {
			return false
		}

		if x == ')' {
			if st[l] != '(' {
				return false
			} else {
				st = st[:l]
			}
			continue
		}

		if x == ']' {
			if st[l] != '[' {
				return false
			} else {
				st = st[:l]
			}
			continue
		}

		if x == '}' {
			if st[l] != '{' {
				return false
			} else {
				st = st[:l]
			}
			continue
		}
	}

	return len(st) == 0

}
