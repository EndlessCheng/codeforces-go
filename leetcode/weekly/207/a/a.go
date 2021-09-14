package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func reorderSpaces(s string) (ans string) {
	c := strings.Count(s, " ")
	ws := strings.Fields(s)
	if len(ws) == 1 {
		return ws[0] + strings.Repeat(" ", c)
	}
	return strings.Join(ws, strings.Repeat(" ", c/(len(ws)-1))) + strings.Repeat(" ", c%(len(ws)-1))
}
