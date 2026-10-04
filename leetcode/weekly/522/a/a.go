package main

// https://space.bilibili.com/206214
func minRotations(s string) (ans int) {
	pre := '0'
	for _, ch := range s {
		d := abs(int(ch) - int(pre))
		ans += min(d, 10-d)
		pre = ch
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
