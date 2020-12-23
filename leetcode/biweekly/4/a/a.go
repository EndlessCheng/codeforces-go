package main

import "time"

// github.com/EndlessCheng/codeforces-go
func numberOfDays(y, m int) int {
	return time.Date(y, time.Month(m+1), 0, 0, 0, 0, 0, time.UTC).Day()
}
