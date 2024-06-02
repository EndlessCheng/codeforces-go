package main

// https://space.bilibili.com/206214
func minimumChairs(s string) (ans int) {
	cnt := 0
	for _, b := range s {
		if b == 'E' {
			cnt++
			ans = max(ans, cnt)
		} else {
			cnt--
		}
	}
	return
}
