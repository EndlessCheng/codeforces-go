package main

import "slices"

// https://space.bilibili.com/206214/dynamic
func distributeCookies(cookies []int, k int) int {
	u := 1 << len(cookies)
	sum := make([]int, u)
	for i, v := range cookies {
		highBit := 1 << i
		for j := range highBit {
			sum[highBit|j] = sum[j] + v
		}
	}

	f := slices.Clone(sum)
	for range k - 1 {
		for j := u - 1; j > 0; j-- {
			for s := j; s > 0; s = (s - 1) & j {
				f[j] = min(f[j], max(f[j^s], sum[s]))
			}
		}
	}
	return f[u-1]
}
