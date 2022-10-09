package main

// https://space.bilibili.com/206214
func findArray(a []int) (ans []int) {
	ans = append(ans, a[0])
	for i := 1; i < len(a); i++ {
		ans = append(ans, a[i]^a[i-1])
	}
	return
}
