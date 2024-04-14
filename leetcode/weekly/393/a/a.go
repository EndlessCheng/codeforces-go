package main

import "fmt"

// https://space.bilibili.com/206214
func findLatestTime(s string) string {
	for h := 11; ; h-- {
		if s[0] != '?' && s[0]-'0' != byte(h/10) || s[1] != '?' && s[1]-'0' != byte(h%10) {
			continue
		}
		for m := 59; m >= 0; m-- {
			if s[3] != '?' && s[3]-'0' != byte(m/10) || s[4] != '?' && s[4]-'0' != byte(m%10) {
				continue
			}
			return fmt.Sprintf("%02d:%02d", h, m)
		}
	}
}
