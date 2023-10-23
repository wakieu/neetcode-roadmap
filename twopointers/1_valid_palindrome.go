package twopointers

// https://leetcode.com/problems/valid-palindrome/

import (
    "regexp"
	"strings"
)

func IsPalindrome(s string) bool {
    var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

    s = nonAlphanumericRegex.ReplaceAllString(s, "")
	s = strings.ToLower(s)

    n := len(s)

    for i := 0; i < n/2; i++ {
        if s[i] != s[n-i-1] {
            return false
        }
    }

    return true
}