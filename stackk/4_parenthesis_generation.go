package stackk

// https://leetcode.com/problems/generate-parentheses/

import "strings"

func GenerateParenthesis(n int) []string {
    var stack []string
    var res []string

    var backtrack func(open, closed int)
    backtrack = func (open, closed int) {
        if open == closed && closed == n {
            res = append(res, strings.Join(stack, ""))
            return
        }

        if open < n {
            stack = append(stack, "(")
            backtrack(open + 1, closed)
            stack = stack[:len(stack)-1]
        }
        if closed < open {
            stack = append(stack, ")")
            backtrack(open, closed + 1)
            stack = stack[:len(stack)-1]
        }
    }

    backtrack(0, 0)

    return res
}
