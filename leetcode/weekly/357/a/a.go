package main

// https://space.bilibili.com/206214
func finalString(s string) (ans string) {
	a := []rune{}
	for _, v := range s {
		if v == 'i' {
			for i, n := 0, len(a); i < n/2; i++ {
				a[i], a[n-1-i] = a[n-1-i], a[i]
			}
		} else {
			a = append(a, v)
		}
	}
	ans = string(a)
	return
}
