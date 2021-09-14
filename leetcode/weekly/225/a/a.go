package main

import "fmt"

// github.com/EndlessCheng/codeforces-go
func maximumTime(time string) string {
	for h := 23; ; h-- {
	o:
		for m := 59; m >= 0; m-- {
			ans := fmt.Sprintf("%02d:%02d", h, m)
			for i, b := range time {
				if b != '?' && time[i] != ans[i] {
					continue o
				}
			}
			return ans
		}
	}
}
