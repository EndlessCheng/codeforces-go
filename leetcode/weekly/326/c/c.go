package main

// https://space.bilibili.com/206214
func minimumPartition(s string, k int) int {
	ans, x := 1, 0
	for _, c := range s {
		v := int(c - '0')
		if v > k {
			return -1
		}
		x = x*10 + v
		if x > k {
			ans++
			x = v
		}
	}
	return ans
}
