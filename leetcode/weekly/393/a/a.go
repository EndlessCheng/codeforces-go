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

func findLatestTime2(s string) string {
	t := []byte(s)
	if t[0] == '?' {
		if t[1] == '?' || t[1] <= '1' {
			t[0] = '1'
		} else {
			t[0] = '0'
		}
	}
	if t[1] == '?' {
		if t[0] == '1' {
			t[1] = '1'
		} else {
			t[1] = '9'
		}
	}
	if t[3] == '?' {
		t[3] = '5'
	}
	if t[4] == '?' {
		t[4] = '9'
	}
	return string(t)
}
