package main

// https://space.bilibili.com/206214/dynamic
func distributeCookies(cookies []int, k int) int {
	m := 1 << len(cookies)
	sum := make([]int, m)
	for i, v := range cookies {
		for mask, bit := 0, 1<<i; mask < bit; mask++ {
			sum[bit|mask] = sum[mask] + v
		}
	}

	f := append([]int{}, sum...)
	for i := 1; i < k; i++ {
		for j := m - 1; j > 0; j-- {
			for s := j; s > 0; s = (s - 1) & j {
				f[j] = min(f[j], max(f[j^s], sum[s]))
			}
		}
	}
	return f[m-1]
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
