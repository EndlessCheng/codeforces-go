package main

// https://space.bilibili.com/206214
func separateDigits(nums []int) (ans []int) {
	for _, x := range nums {
		i0 := len(ans)
		for ; x > 0; x /= 10 {
			ans = append(ans, x%10)
		}
		b := ans[i0:]
		for i, n := 0, len(b); i < n/2; i++ {
			b[i], b[n-1-i] = b[n-1-i], b[i]
		}
	}
	return
}
