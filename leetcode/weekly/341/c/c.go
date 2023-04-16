package main

// https://space.bilibili.com/206214
func addMinimum(s string) int {
	t := 1
	for i := 1; i < len(s); i++ {
		if s[i-1] >= s[i] {
			t++
		}
	}
	return t*3 - len(s)
}
