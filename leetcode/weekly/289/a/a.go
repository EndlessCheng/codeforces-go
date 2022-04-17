package main

import (
	"strconv"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func digitSum(s string, k int) string {
	for len(s) > k {
		t := &strings.Builder{}
		for {
			sum := 0
			for i := 0; i < len(s) && i < k; i++ {
				sum += int(s[i] & 15)
			}
			t.WriteString(strconv.Itoa(sum))
			if len(s) <= k {
				break
			}
			s = s[k:]
		}
		s = t.String()
	}
	return s
}
