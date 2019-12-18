package main

import "strings"

func gcdOfStrings(s1 string, s2 string) string {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	len1, len2 := len(s1), len(s2)
	g := gcd(len1, len2)
	for i := g; i > 0; i-- {
		if g%i == 0 {
			if s := s1[:i]; strings.Count(s1, s) == len1/i && strings.Count(s2, s) == len2/i {
				return s
			}
		}
	}
	return ""
}
