package main

// https://space.bilibili.com/206214
func maxSubarrays(n int, conflictingPairs [][]int) int64 {
	groups := make([][2]int, n+1) // [][2]int 比 [][]int 快
	for i := range groups {
		groups[i] = [2]int{n + 1, n + 1}
	}
	for _, p := range conflictingPairs {
		a, b := p[0], p[1]
		if a > b {
			a, b = b, a
		}
		g := &groups[a]
		if b < g[0] {
			g[0], g[1] = b, g[0]
		} else if b < g[1] {
			g[1] = b
		}
	}

	var ans, maxExtra, extra int
	b0, b1 := n+1, n+1
	for i := n; i > 0; i-- {
		preB0 := b0
		for _, b := range groups[i] {
			if b < b0 {
				b0, b1 = b, b0
			} else if b < b1 {
				b1 = b
			}
		}

		ans += b0 - i
		if b0 != preB0 { // 重新统计连续相同 b0 的 extra
			extra = 0
		}
		extra += b1 - b0
		maxExtra = max(maxExtra, extra)
	}

	return int64(ans + maxExtra)
}

func maxSubarrays2(n int, conflictingPairs [][]int) int64 {
	groups := make([][2]int, n+1)
	for i := range groups {
		groups[i] = [2]int{n + 1, n + 1}
	}
	for _, p := range conflictingPairs {
		a, b := p[0], p[1]
		if a > b {
			a, b = b, a
		}
		g := &groups[a]
		if b < g[0] {
			g[0], g[1] = b, g[0]
		} else if b < g[1] {
			g[1] = b
		}
	}

	var ans, extra, maxExtra int
	b0, b1 := n+1, n+1
	for i := n; i > 0; i-- {
		preB0 := b0

		b, c := groups[i][0], groups[i][1]
		if b < b0 {
			b1 = min(b0, c)
			b0 = b
		} else if b < b1 {
			b1 = b
		} else if c < b1 {
			b1 = c
		}

		ans += b0 - i
		if b0 != preB0 {
			extra = 0
		}
		extra += b1 - b0
		maxExtra = max(maxExtra, extra)
	}

	return int64(ans + maxExtra)
}
