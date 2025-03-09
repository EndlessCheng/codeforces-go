package main

import "slices"

// https://space.bilibili.com/206214
func maxSubarrays(n int, conflictingPairs [][]int) int64 {
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
	for a := n; a > 0; a-- {
		preB0 := b0
		for _, b := range groups[a] {
			if b < b0 {
				b0, b1 = b, b0
			} else if b < b1 {
				b1 = b
			}
		}
		ans += b0 - a
		if b0 != preB0 {
			extra = 0
		}
		extra += b1 - b0
		maxExtra = max(maxExtra, extra)
	}

	return int64(ans + maxExtra)
}

func maxSubarrays1(n int, conflictingPairs [][]int) int64 {
	groups := make([][]int, n+1)
	for _, p := range conflictingPairs {
		a, b := p[0], p[1]
		if a > b {
			a, b = b, a
		}
		groups[a] = append(groups[a], b)
	}

	ans := 0
	extra := make([]int, n+2)
	b := []int{n + 1, n + 1} // 维护最小 b 和次小 b
	for a := n; a > 0; a-- {
		listB := groups[a]
		if listB != nil {
			slices.Sort(listB)
			if len(listB) > 2 {
				listB = listB[:2]
			}
			b = append(b, listB...)
			slices.Sort(b)
			b = b[:2]
		}
		ans += b[0] - a
		extra[b[0]] += b[1] - b[0]
	}

	return int64(ans + slices.Max(extra))
}
