package main

// github.com/EndlessCheng/codeforces-go
func minimumWhiteTiles(floor string, numCarpets, carpetLen int) int {
	m := len(floor)
	if numCarpets*carpetLen >= m {
		return 0
	}

	f := make([]int, m)
	f[0] = int(floor[0] - '0')
	for j := 1; j < m; j++ {
		f[j] = f[j-1] + int(floor[j]-'0')
	}
	for i := 1; i <= numCarpets; i++ {
		nf := make([]int, m)
		for j := carpetLen * i; j < m; j++ {
			nf[j] = min(nf[j-1]+int(floor[j]-'0'), f[j-carpetLen])
		}
		f = nf
	}
	return f[m-1]
}

func minimumWhiteTiles2(floor string, numCarpets, carpetLen int) int {
	m := len(floor)
	f := make([][]int, numCarpets+1)
	for i := range f {
		f[i] = make([]int, m)
	}
	// 单独计算 i=0 的情况，本质是 floor 的前缀和
	f[0][0] = int(floor[0] - '0')
	for j := 1; j < m; j++ {
		f[0][j] = f[0][j-1] + int(floor[j]-'0')
	}
	for i := 1; i <= numCarpets; i++ {
		for j := carpetLen * i; j < m; j++ {
			f[i][j] = min(f[i][j-1]+int(floor[j]-'0'), f[i-1][j-carpetLen])
		}
	}
	return f[numCarpets][m-1]
}

func minimumWhiteTiles1(floor string, numCarpets, carpetLen int) int {
	m := len(floor)
	memo := make([][]int, numCarpets+1)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if j < carpetLen*i {
			return
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		res = dfs(i, j-1) + int(floor[j]-'0')
		if i > 0 {
			res = min(res, dfs(i-1, j-carpetLen))
		}
		return
	}
	return dfs(numCarpets, m-1)
}
