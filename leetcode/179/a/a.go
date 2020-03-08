package main

import "strings"

func generateTheString(n int) string {
	if n&1 == 1 {
		return strings.Repeat("a", n)
	}
	return strings.Repeat("a", n-1) + "b"
}
