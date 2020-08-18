package main

import (
	"fmt"
	"strings"
	"time"
)

// github.com/EndlessCheng/codeforces-go
func reformatDate(s string) (ans string) {
	sp := strings.Split(s, " ")
	d := sp[0]
	if len(d) < 4 {
		d = "0" + d
	}
	t, _ := time.Parse("Jan", sp[1])
	return fmt.Sprintf("%s-%02d-%s", sp[2], t.Month(), d[:2])
}
