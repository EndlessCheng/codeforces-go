package main

import "math"

// https://space.bilibili.com/206214
func maxRemovals(source, pattern string, targetIndices []int) int {
	m := len(pattern)
	f := make([]int, m+1)
	for i := 1; i <= m; i++ {
		f[i] = math.MinInt
	}
	k := 0
	for i := range source {
		if k < len(targetIndices) && targetIndices[k] < i {
			k++
		}
		isDel := 0
		if k < len(targetIndices) && targetIndices[k] == i {
			isDel = 1
		}
		for j := min(i, m-1); j >= 0; j-- {
			f[j+1] += isDel
			if source[i] == pattern[j] {
				f[j+1] = max(f[j+1], f[j])
			}
		}
		f[0] += isDel
	}
	return f[m]
}

func maxRemovals2(source, pattern string, targetIndices []int) int {
	targetSet := map[int]int{}
	for _, idx := range targetIndices {
		targetSet[idx] = 1
	}
	n, m := len(source), len(pattern)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < j {
			return math.MinInt
		}
		if i < 0 {
			return 0
		}
		p := &memo[i][j+1]
		if *p != -1 {
			return *p
		}
		res := dfs(i-1, j) + targetSet[i]
		if j >= 0 && source[i] == pattern[j] {
			res = max(res, dfs(i-1, j-1))
		}
		*p = res
		return res
	}
	return dfs(n-1, m-1)
}
