package main

import "strings"

func findOcurrences(text string, first string, second string) (ans []string) {
	ss := strings.Split(text, " ")
	for i := 0; i <= len(ss)-3; i++ {
		if ss[i] == first && ss[i+1] == second {
			ans = append(ans, ss[i+2])
		}
	}
	return ans
}
