package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	f := make([]int64, n+1)
	for i, q := range questions {
		f[i+1] = max(f[i+1], f[i])
		j := min(i+q[1]+1, n)
		f[j] = max(f[j], f[i]+int64(q[0]))
	}
	return f[n]
}

func mostPoints2(questions [][]int) int64 {
	n := len(questions)
	f := make([]int64, n+1)
	for i, q := range slices.Backward(questions) {
		j := min(i+q[1]+1, n)
		f[i] = max(f[i+1], f[j]+int64(q[0]))
	}
	return f[0]
}

func mostPoints1(questions [][]int) int64 {
	n := len(questions)
	memo := make([]int64, n)
	var dfs func(int) int64
	dfs = func(i int) int64 {
		if i >= n {
			return 0
		}
		p := &memo[i]
		if *p == 0 {
			q := questions[i]
			*p = max(dfs(i+1), dfs(i+q[1]+1)+int64(q[0]))
		}
		return *p
	}
	return dfs(0)
}
