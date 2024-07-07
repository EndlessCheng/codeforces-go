package main

// https://space.bilibili.com/206214
func numberOfAlternatingGroups(a []int) (ans int) {
	a = append(a, a...)
	c := 0
	for i, v := range a {
		if i > 0 && v == a[i-1] {
			c = 0
		}
		c++
		if i >= len(a)/2 && c >= 3 {
			ans++
		}
	}
	return
}
