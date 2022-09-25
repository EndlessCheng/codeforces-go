package main

import "math"

// https://space.bilibili.com/206214
func unSuitability(operate []int) int {
	const inf = math.MaxInt32
	mx := 0
	for _, x := range operate {
		mx = max(mx, x)
	}
	mx *= 2
	pre := make([]int, mx+1)
	for i := range pre {
		pre[i] = inf
	}
	pre[0] = 0
	f := make([]int, mx+1)
	for _, x := range operate {
		for i := range f {
			f[i] = inf
		}
		for j, dis := range pre {
			if pre[j] == inf {
				continue
			}
			if j+x <= mx {
				f[j+x] = min(f[j+x], max(dis, j+x))
			}
			if j >= x {
				f[j-x] = min(f[j-x], dis)
			} else {
				f[0] = min(f[0], dis-j+x)
			}
		}
		pre, f = f, pre
	}
	ans := inf
	for _, x := range pre {
		ans = min(ans, x)
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
