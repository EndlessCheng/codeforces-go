package main

// https://space.bilibili.com/206214
func getSmallestString(s string) string {
	t := []byte(s)
	for i := 1; i < len(t); i++ {
		x, y := t[i-1], t[i]
		if x > y && x%2 == y%2 {
			t[i-1], t[i] = y, x
			break
		}
	}
	return string(t)
}
