package stackk

// https://leetcode.com/problems/evaluate-reverse-polish-notation/

import "strconv"

func EvalRPN(tokens []string) int {
    var i int = 0
    for len(tokens) > 1 {
        if tokens[i] == "+" {
			n1, err1 := strconv.Atoi(tokens[i-2])
			n2, err2 := strconv.Atoi(tokens[i-1])
			r := strconv.Itoa(n1 + n2)
			if err1 != nil || err2 != nil {
				panic("int conversion error")
			}
			tokens = append(append(tokens[:i-2], r), tokens[i+1:]...)
			i = i - 2
        } else if tokens[i] == "-" {
			n1, err1 := strconv.Atoi(tokens[i-2])
			n2, err2 := strconv.Atoi(tokens[i-1])
			r := strconv.Itoa(n1 - n2)
			if err1 != nil || err2 != nil {
				panic("int conversion error")
			}
			tokens = append(append(tokens[:i-2], r), tokens[i+1:]...)
			i = i - 2
        } else if tokens[i] == "*" {
			n1, err1 := strconv.Atoi(tokens[i-2])
			n2, err2 := strconv.Atoi(tokens[i-1])
			r := strconv.Itoa(n1 * n2)
			if err1 != nil || err2 != nil {
				panic("int conversion error")
			}
			tokens = append(append(tokens[:i-2], r), tokens[i+1:]...)
			i = i - 2
        } else if tokens[i] == "/" {
			n1, err1 := strconv.Atoi(tokens[i-2])
			n2, err2 := strconv.Atoi(tokens[i-1])
			r := strconv.Itoa(n1 / n2)
			if err1 != nil || err2 != nil {
				panic("int conversion error")
			}
			tokens = append(append(tokens[:i-2], r), tokens[i+1:]...)
			i = i - 2
        } else {
            i++
        }
    }

	x, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic("int conversion error")
	}

    return x
}
