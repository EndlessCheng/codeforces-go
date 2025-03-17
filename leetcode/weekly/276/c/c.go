package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	f := make([]int, n+1)
	for i, q := range slices.Backward(questions) {
		j := min(i+q[1]+1, n)
		f[i] = max(f[i+1], f[j]+q[0])
	}
	return int64(f[0])
}

func mostPoints1(questions [][]int) int64 {
	n := len(questions)
	memo := make([]int, n)
	var dfs func(int) int
	dfs = func(i int) int {
		if i >= n {
			return 0
		}
		p := &memo[i]
		if *p == 0 {
			*p = max(dfs(i+1), dfs(i+questions[i][1]+1)+questions[i][0])
		}
		return *p
	}
	return int64(dfs(0))
}
