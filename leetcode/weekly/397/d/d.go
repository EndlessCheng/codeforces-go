package main

import (
	"math"
)

// https://space.bilibili.com/206214
func findPermutation(a []int) []int {
	n := len(a)
	u := 1<<n - 1
	f := make([][]int, 1<<n)
	g := make([][]int, 1<<n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = math.MaxInt
		}
		g[i] = make([]int, n)
	}
	for j := range f[u] {
		f[u][j] = abs(j - a[0])
		g[u][j] = -1
	}
	for s := u - 2; s > 0; s -= 2 { // 注意偶数不含 0，是无效状态
		for j := 0; j < n; j++ {
			if s>>j&1 == 0 { // 无效状态，因为 j 一定在 s 中
				continue
			}
			for k := 1; k < n; k++ {
				if s>>k&1 > 0 { // k 之前填过
					continue
				}
				v := f[s|1<<k][k] + abs(j-a[k])
				if v < f[s][j] {
					f[s][j] = v
					g[s][j] = k // 记录该状态下填了哪个数
				}
			}
		}
	}

	ans := make([]int, 0, n)
	for s, j := 0, 0; j >= 0; {
		ans = append(ans, j)
		s |= 1 << j
		j = g[s][j]
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findPermutation2(a []int) []int {
	n := len(a)
	memo := make([][]int, 1<<n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(s, j int) int {
		if s == 1<<n-1 {
			return abs(j - a[0])
		}
		p := &memo[s][j]
		if *p != -1 {
			return *p
		}
		res := math.MaxInt
		for k := 1; k < n; k++ {
			if s>>k&1 == 0 { // k 之前没填过
				res = min(res, dfs(s|1<<k, k)+abs(j-a[k]))
			}
		}
		*p = res
		return res
	}

	ans := make([]int, 0, n)
	var makeAns func(int, int)
	makeAns = func(s, j int) {
		ans = append(ans, j)
		if s == 1<<n-1 {
			return
		}
		finalRes := dfs(s, j)
		for k := 1; k < n; k++ {
			if s>>k&1 == 0 && dfs(s|1<<k, k)+abs(j-a[k]) == finalRes {
				makeAns(s|1<<k, k)
				break
			}
		}
	}
	makeAns(1, 0)
	return ans
}
