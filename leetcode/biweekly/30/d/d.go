package main

// github.com/EndlessCheng/codeforces-go
// https://oeis.org/A030193
func winnerSquareGame1(n int) bool {
	memo := make([]int8, n+1)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) bool
	dfs = func(i int) bool {
		if i == 0 {
			return false
		}
		p := &memo[i]
		if *p != -1 {
			return *p == 1
		}
		for x := 1; x*x <= i; x++ {
			if !dfs(i - x*x) {
				*p = 1
				return true
			}
		}
		*p = 0
		return false
	}

	return dfs(n)
}

func winnerSquareGame2(n int) bool {
	f := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		for x := 1; x*x <= i; x++ {
			if !f[i-x*x] {
				f[i] = true
				break
			}
		}
	}
	return f[n]
}

func winnerSquareGame(n int) bool {
	f := make([]bool, n+1)
	for i := range n {
		if f[i] {
			continue
		}
		for x := 1; x*x <= n-i; x++ {
			f[i+x*x] = true
		}
	}
	return f[n]
}
