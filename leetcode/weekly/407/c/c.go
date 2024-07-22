package main

// https://space.bilibili.com/206214
func maxOperations(s string) (ans int) {
	cnt1 := 0
	for i, c := range s {
		if c == '1' {
			cnt1++
		} else if i > 0 && s[i-1] == '1' {
			ans += cnt1
		}
	}
	return
}
