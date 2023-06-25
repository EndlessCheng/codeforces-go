package main

// https://space.bilibili.com/206214
func minimizeConcatenatedLength(words []string) (ans int) {
	n := len(words)
	memo := make([][26][2]int, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1}
		}
	}
	var dfs func(int, byte, int) int
	dfs = func(i int, c byte, j int) (res int) {
		if i == n {
			return
		}
		dv := &memo[i][c][j]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()

		pre, cur := words[i-1], words[i]
		pl, pr := pre[0]-'a', pre[len(pre)-1]-'a'
		cl, cr := cur[0]-'a', cur[len(cur)-1]-'a'
		if j == 0 {
			res1 := dfs(i+1, c, 0)
			if cr == pl {
				res1--
			}
			res2 := dfs(i+1, pl, 1)
			if cl == c {
				res2--
			}
			return min(res1, res2) + len(cur)
		} else {
			res1 := dfs(i+1, pr, 0)
			if cr == c {
				res1--
			}
			res2 := dfs(i+1, c, 1)
			if cl == pr {
				res2--
			}
			return min(res1, res2) + len(cur)
		}
	}
	return dfs(1, words[0][0]-'a', 1) + len(words[0])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
