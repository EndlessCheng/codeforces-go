package main

// https://space.bilibili.com/206214
func maxValue(nums []int, k int) (ans int) {
	const mx = 1 << 7
	n := len(nums)
	suf := make([][mx]bool, n)
	f := make([][mx]bool, k+1)
	f[0][0] = true
	for i := n - 1; i >= k; i-- {
		v := nums[i]
		// 注意当 i 比较大的时候，循环次数应和 i 有关，因为更大的 j，对应的 f[j] 全为 false
		for j := min(k-1, n-1-i); j >= 0; j-- {
			for x, hasX := range f[j] {
				if hasX {
					f[j+1][x|v] = true
				}
			}
		}
		suf[i] = f[k]
	}

	pre := make([][mx]bool, k+1)
	pre[0][0] = true
	for i, v := range nums[:n-k] {
		for j := min(k-1, i); j >= 0; j-- {
			for x, hasX := range pre[j] {
				if hasX {
					pre[j+1][x|v] = true
				}
			}
		}
		if i < k-1 {
			continue
		}
		for x, hasX := range pre[k] {
			if hasX {
				for y, hasY := range suf[i+1] {
					if hasY {
						ans = max(ans, x^y)
					}
				}
			}
		}
		if ans == mx-1 {
			return
		}
	}
	return
}
