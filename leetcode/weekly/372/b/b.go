package main

// https://space.bilibili.com/206214
func minimumSteps(s string) (ans int64) {
	cnt1 := 0
	for _, c := range s {
		if c == '0' {
			ans += int64(cnt1)
		} else {
			cnt1++
		}
	}
	return
}
