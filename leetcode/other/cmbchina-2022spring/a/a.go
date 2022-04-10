package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func deleteText(s string, i int) (ans string) {
	if s[i] == ' ' {
		return s
	}
	sp := strings.Split(s, " ")
	cnt := strings.Count(s[:i], " ")
	sp = append(sp[:cnt], sp[cnt+1:]...)
	ans = strings.Join(sp, " ")
	return
}
