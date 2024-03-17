package main

// https://space.bilibili.com/206214
func sumOfEncryptedInt(nums []int) (ans int) {
	for _, x := range nums {
		mx, base := 0, 0
		for ; x > 0; x /= 10 {
			mx = max(mx, x%10)
			base = base*10 + 1
		}
		ans += mx * base
	}
	return
}
