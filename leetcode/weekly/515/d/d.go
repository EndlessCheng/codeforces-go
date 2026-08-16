package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func elevatorRequests1(n int, start int, requests [][]int) int64 {
	m := len(requests)
	memo := make([][]int, 1<<m)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}

	// 返回处理完请求集合 mask，且电梯停在 requests[i][1]，所需的最短时间
	var dfs func(int, int) int
	dfs = func(mask, i int) int {
		mask ^= 1 << i
		req := requests[i]
		t, x := req[0], req[1]
		if mask == 0 {
			// i 是第一个被处理的请求
			return max(abs(x-start), t)
		}

		p := &memo[mask][i]
		if *p != -1 { // 之前计算过
			return *p
		}

		res := math.MaxInt
		for j, r := range requests {
			if mask>>j&1 > 0 {
				// 处理完请求 j 的时间 + 从 j 到 i 的时间
				res = min(res, dfs(mask, j)+abs(x-r[1]))
			}
		}
		// 处理完请求 i 的时间不能早于 t
		res = max(res, t)

		*p = res // 记忆化
		return res
	}

	ans := math.MaxInt
	for i := range m {
		ans = min(ans, dfs(1<<m-1, i))
	}
	return int64(ans)
}

func elevatorRequests(n int, start int, requests [][]int) int64 {
	m := len(requests)
	f := make([][]int, 1<<m)
	for i := range f {
		f[i] = make([]int, m)
	}

	for i, req := range requests {
		f[1<<i][i] = max(abs(req[1]-start), req[0])
	}

	for mask := 1; mask < 1<<m; mask++ {
		if mask&(mask-1) == 0 {
			continue
		}
		for i, req := range requests {
			if mask>>i&1 == 0 {
				continue
			}
			res := math.MaxInt
			msk := mask ^ 1<<i
			t, x := req[0], req[1]
			for j, r := range requests {
				if msk>>j&1 > 0 {
					res = min(res, f[msk][j]+abs(x-r[1]))
				}
			}
			f[mask][i] = max(res, t)
		}
	}

	return int64(slices.Min(f[1<<m-1]))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
