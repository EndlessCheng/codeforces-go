package main

import "slices"

// https://space.bilibili.com/206214
func makeStringGood(s string) int {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}

	targets := map[int]struct{}{}
	targets[cnt[0]] = struct{}{}
	for i := 1; i < len(cnt); i++ {
		x, y := cnt[i-1], cnt[i]
		targets[y] = struct{}{}
		targets[x+y] = struct{}{}
		targets[(x+y)/2] = struct{}{}
		targets[(x+y+1)/2] = struct{}{}
	}

	ans := len(s) // target = 0 时的答案
	f := [27]int{}
	for target := range targets {
		f[25] = min(cnt[25], abs(cnt[25]-target))
		for i := 24; i >= 0; i-- {
			x := cnt[i]
			if x == 0 {
				f[i] = f[i+1]
				continue
			}
			// 单独操作 x（变成 target 或 0）
			f[i] = f[i+1] + min(x, abs(x-target))
			// x 变成 target 或 0，y 变成 target
			y := cnt[i+1]
			if 0 < y && y < target { // 只有当 y 需要变大时，才去执行第三种操作
				if x > target { // x 变成 target
					f[i] = min(f[i], f[i+2]+max(x-target, target-y))
				} else { // x 变成 0
					f[i] = min(f[i], f[i+2]+max(x, target-y))
				}
			}
		}
		ans = min(ans, f[0])
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }

func makeStringGoodAC(s string) int {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	m := slices.Max(cnt[:])

	ans := len(s) // target = 0 时的答案
	f := [27]int{}
	for target := 1; target <= m; target++ {
		f[25] = min(cnt[25], abs(cnt[25]-target))
		for i := 24; i >= 0; i-- {
			x := cnt[i]
			if x == 0 {
				f[i] = f[i+1]
				continue
			}
			// 单独操作 x（变成 target 或 0）
			f[i] = f[i+1] + min(x, abs(x-target))
			// x 变成 target 或 0，y 变成 target
			y := cnt[i+1]
			if 0 < y && y < target { // 只有当 y 需要变大时，才去执行第三种操作
				if x > target { // x 变成 target
					f[i] = min(f[i], f[i+2]+max(x-target, target-y))
				} else { // x 变成 0
					f[i] = min(f[i], f[i+2]+max(x, target-y))
				}
			}
		}
		ans = min(ans, f[0])
	}
	return ans
}

func makeStringGood2(s string) int {
	n := len(s)
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	mx := slices.Max(cnt[:])

	ans := n
	for target := 0; target <= mx; target++ {
		calcOp := func(v int) int { return min(v, abs(target-v)) }
		memo := [26]int{}
		for i := range memo {
			memo[i] = -1
		}
		var dfs func(int) int
		dfs = func(i int) (res int) {
			if i >= 26 {
				return
			}
			p := &memo[i]
			if *p != -1 {
				return *p
			}
			defer func() { *p = res }()
			x := cnt[i]
			if x == 0 {
				return dfs(i + 1)
			}
			if i == 25 {
				return dfs(i+1) + calcOp(x)
			}
			y := cnt[i+1]
			mn := calcOp(x) + calcOp(y)
			// 只有当 y 需要变大时，才去执行第三种操作
			if y < target {
				if x > target { // x 夹带着第三种操作，变成 target
					mn = min(mn, max(x-target, target-y))
				} else { // x 夹带着第三种操作，变成 0
					mn = min(mn, max(x, target-y))
				}
			}
			return min(dfs(i+1)+calcOp(x), dfs(i+2)+mn)
		}
		ans = min(ans, dfs(0))
	}
	return ans
}
