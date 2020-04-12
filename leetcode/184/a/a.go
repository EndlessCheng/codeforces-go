package main

import "strings"

func stringMatching(a []string) (ans []string) {
	for i, s := range a {
		for j, s2 := range a {
			if j != i && strings.Contains(s2, s) {
				ans = append(ans, s)
				break
			}
		}
	}
	return
}
