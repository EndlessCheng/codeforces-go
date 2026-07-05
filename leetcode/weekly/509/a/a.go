package main

// https://space.bilibili.com/206214
func maxDigitRange(nums []int) (ans int) {
	maxRange := 0
	for _, x := range nums {
		mn, mx := 9, 0
		for v := x; v > 0; v /= 10 {
			d := v % 10
			mn = min(mn, d)
			mx = max(mx, d)
		}

		r := mx - mn
		if r > maxRange {
			maxRange = r
			ans = x // 重新累加
		} else if r == maxRange {
			ans += x
		}
	}
	return
}
