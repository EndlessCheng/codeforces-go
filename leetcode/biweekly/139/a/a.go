package main

// https://space.bilibili.com/206214
func stableMountains(height []int, threshold int) (ans []int) {
	for i, h := range height[:len(height)-1] {
		if h > threshold {
			ans = append(ans, i+1)
		}
	}
	return
}
