package main

// https://space.bilibili.com/206214
func minAdjacentSwaps(nums []int, a, b int) (ans int) {
	const mod = 1_000_000_007
	cnt1, cnt2 := 0, 0
	for _, x := range nums {
		if x < a { // x 视作 0
			ans += cnt1 + cnt2
		} else if x <= b { // x 视作 1
			ans += cnt2
			cnt1++
		} else { // x 视作 2
			cnt2++
		}
	}
	return ans % mod
}
