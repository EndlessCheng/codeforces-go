package main

import (
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func largestGoodInteger(num string) string {
	mx := byte(0)
	cnt := 1
	for i := 1; i < len(num); i++ {
		d := num[i]
		if d != num[i-1] {
			cnt = 1
			continue
		}
		cnt++
		if cnt == 3 && d > mx {
			mx = d
		}
	}
	if mx == 0 {
		return ""
	}
	return strings.Repeat(string(mx), 3)
}

func largestGoodInteger2(num string) string {
	mx := byte(0)
	for i := range len(num) - 2 {
		d := num[i]
		if d > mx && d == num[i+1] && d == num[i+2] {
			mx = d
		}
	}
	if mx == 0 {
		return ""
	}
	return strings.Repeat(string(mx), 3)
}

func largestGoodInteger3(num string) string {
	for d := '9'; d >= '0'; d-- {
		s := strings.Repeat(string(d), 3)
		if strings.Contains(num, s) {
			return s
		}
	}
	return ""
}
