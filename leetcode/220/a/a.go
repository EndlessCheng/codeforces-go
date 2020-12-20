package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func reformatNumber(s string) (ans string) {
	s = strings.NewReplacer(" ", "", "-", "").Replace(s)
	t := []byte{}
	for ; len(s) > 4; s = s[3:] {
		t = append(t, s[:3]+"-"...)
	}
	if len(s) < 4 {
		t = append(t, s...)
	} else {
		t = append(t, s[:2]+"-"+s[2:]...)
	}
	return string(t)
}
