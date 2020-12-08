package main

import (
	"fmt"
	"time"
)

// github.com/EndlessCheng/codeforces-go
func reformatDate(date string) string {
	var y, d int
	var s string
	fmt.Sscanf(date, "%d%s%s%d", &d, &s, &s, &y)
	t, _ := time.Parse("Jan", s)
	return fmt.Sprintf("%d-%02d-%02d", y, t.Month(), d)
}
