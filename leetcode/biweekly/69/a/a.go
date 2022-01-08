package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func capitalizeTitle(title string) string {
	a := strings.Split(title, " ")
	for i, s := range a {
		a[i] = strings.ToLower(s)
		if len(s) > 2 {
			a[i] = strings.Title(a[i])
		}
	}
	return strings.Join(a, " ")
}
