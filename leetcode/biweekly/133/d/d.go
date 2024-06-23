package main

import "slices"

// https://space.bilibili.com/206214
func numberOfPermutations(n int, requirements [][]int) int {
	const mod = 1_000_000_007
	req := make([]int, n)
	for i := 1; i < n; i++ {
		req[i] = -1
	}
	for _, p := range requirements {
		req[p[0]] = p[1]
	}
	if req[0] > 0 {
		return 0
	}

	m := slices.Max(req)
	f := make([]int, m+1)
	f[0] = 1
	for i := 1; i < n; i++ {
		mx := m
		if req[i] >= 0 {
			mx = req[i]
		}
		if r := req[i-1]; r >= 0 {
			clear(f[:r])
			for j := r + 1; j <= min(i+r, mx); j++ {
				f[j] = f[r]
			}
			clear(f[min(i+r, mx)+1:])
		} else {
			for j := 1; j <= mx; j++ {
				f[j] = (f[j] + f[j-1]) % mod
			}
			for j := mx; j > i; j-- {
				f[j] = (f[j] - f[j-i-1] + mod) % mod
			}
		}
	}
	return f[req[n-1]]
}

func numberOfPermutations2(n int, requirements [][]int) int {
	const mod = 1_000_000_007
	req := make([]int, n)
	for i := 1; i < n; i++ {
		req[i] = -1
	}
	for _, p := range requirements {
		req[p[0]] = p[1]
	}
	if req[0] > 0 {
		return 0
	}

	m := slices.Max(req)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	f[0][0] = 1
	for i := 1; i < n; i++ {
		mx := m
		if req[i] >= 0 {
			mx = req[i]
		}
		if r := req[i-1]; r >= 0 {
			for j := r; j <= min(i+r, m); j++ {
				f[i][j] = f[i-1][r]
			}
		} else {
			for j := 0; j <= mx; j++ {
				for k := 0; k <= min(i, j); k++ {
					f[i][j] = (f[i][j] + f[i-1][j-k]) % mod
				}
			}
		}
	}
	return f[n-1][req[n-1]]
}

func numberOfPermutations1(n int, requirements [][]int) int {
	const mod = 1_000_000_007
	req := make([]int, n)
	for i := 1; i < n; i++ {
		req[i] = -1
	}
	for _, p := range requirements {
		req[p[0]] = p[1]
	}
	if req[0] > 0 {
		return 0
	}

	m := slices.Max(req)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i == 0 {
			return 1
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if r := req[i-1]; r >= 0 {
			if j < r || j-i > r {
				return 0
			}
			return dfs(i-1, r)
		}
		for k := 0; k <= min(i, j); k++ {
			res += dfs(i-1, j-k)
		}
		return res % mod
	}
	return dfs(n-1, req[n-1])
}
