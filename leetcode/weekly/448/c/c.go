package main

import "math"

// https://space.bilibili.com/206214
func minTravelTime(l, n, k int, position, time []int) int {
	s := make([]int, n)
	for i, t := range time[:n-1] {
		s[i+1] = s[i] + t
	}

	f := make([][][]int, k+1)
	for i := range f {
		f[i] = make([][]int, n)
		for j := range f[i] {
			f[i][j] = make([]int, n)
		}
	}
	for leftK := 1; leftK <= k; leftK++ {
		for pre := range n {
			f[leftK][n-1][pre] = math.MaxInt / 2
		}
	}

	for leftK := range f {
		for i := n - 2; i >= 0; i-- {
			for pre := range i + 1 {
				t := s[i+1] - s[pre]
				res := math.MaxInt
				for nxt := i + 1; nxt < min(n, i+2+leftK); nxt++ {
					res = min(res, f[leftK-(nxt-i-1)][nxt][i+1]+(position[nxt]-position[i])*t)
				}
				f[leftK][i][pre] = res
			}
		}
	}
	return f[k][0][0]
}

func minTravelTime1(l, n, k int, position, time []int) int {
	s := make([]int, n)
	for i, t := range time[:n-1] {
		s[i+1] = s[i] + t
	}

	memo := make([][][]int, k+1)
	for i := range memo {
		memo[i] = make([][]int, n-1)
		for j := range memo[i] {
			memo[i][j] = make([]int, n-1)
		}
	}
	var dfs func(int, int, int) int
	dfs = func(leftK, i, pre int) int {
		if i == n-1 {
			if leftK > 0 {
				return math.MaxInt / 2
			}
			return 0
		}
		p := &memo[leftK][i][pre]
		if *p > 0 {
			return *p
		}
		res := math.MaxInt
		t := s[i+1] - s[pre]
		for nxt := i + 1; nxt < min(n, i+2+leftK); nxt++ {
			res = min(res, dfs(leftK-(nxt-i-1), nxt, i+1)+(position[nxt]-position[i])*t)
		}
		*p = res
		return res
	}
	return dfs(k, 0, 0)
}
