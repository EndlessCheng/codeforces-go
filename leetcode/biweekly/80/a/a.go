package main

import (
	"unicode"
)

// https://space.bilibili.com/206214/dynamic
func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 { return false }
	mask := 0
	for i, c := range password {
		if i > 0 && byte(c) == password[i-1] { return false }
		switch {
		case unicode.IsLower(c): mask |= 1
		case unicode.IsUpper(c): mask |= 2
		case unicode.IsDigit(c): mask |= 4
		default: mask |= 8
		}
	}
	return mask == 15
}
