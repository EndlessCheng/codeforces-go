package main

// https://space.bilibili.com/206214
func beautifulBouquet(flowers []int, cnt int) (ans int) {
	c := map[int]int{}
	left := 0
	for right, x := range flowers {
		c[x]++
		for c[x] > cnt {
			c[flowers[left]]--
			left++
		}
		ans += right - left + 1
	}
	return ans % (1e9 + 7)
}
