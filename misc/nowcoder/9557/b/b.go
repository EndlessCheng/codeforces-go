package main

import "strconv"

// github.com/EndlessCheng/codeforces-go
func solve(t string) int64 {
	s := []int{}
	for i := 0; i < len(t); i++ {
		if t[i] >= '0' && t[i] <= '9' {
			st := i
			for ; t[i] != '#'; i++ {
			}
			v, _ := strconv.Atoi(t[st:i])
			s = append(s, v)
		} else {
			l, r := s[len(s)-2], s[len(s)-1]
			s = s[:len(s)-2]
			if t[i] == '+' {
				s = append(s, l+r)
			} else if t[i] == '-' {
				s = append(s, l-r)
			} else {
				s = append(s, l*r)
			}
		}
	}
	return int64(s[0])
}
