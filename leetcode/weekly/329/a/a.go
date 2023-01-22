package main

// https://space.bilibili.com/206214
func alternateDigitSum(n int) (ans int) {
	sign := 1
	for ; n > 0; n /= 10 {
		ans += n % 10 * sign
		sign = -sign
	}
	return ans * -sign
}
