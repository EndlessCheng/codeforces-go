package main

// https://space.bilibili.com/206214
func countWinningSequences(s string) int {
	const mod = 1_000_000_007
	mp := [...]int{'F': 0, 'W': 1, 'E': 2}
	n := len(s)
	f := make([][][3]int, n+1)
	for i := range f {
		f[i] = make([][3]int, n*2+1)
	}
	for j := n + 1; j <= n*2; j++ {
		f[0][j] = [3]int{1, 1, 1}
	}
	pow2 := 1
	for i, c := range s {
		pow2 = pow2 * 2 % mod
		for j := -i; j < n-i; j++ {
			for pre := 0; pre < 3; pre++ {
				if j > i+1 {
					f[i+1][j+n][pre] = pow2
					continue
				}
				res := 0
				for cur := 0; cur < 3; cur++ {
					if i == n-1 || cur != pre {
						score := (cur - mp[c] + 3) % 3
						if score == 2 {
							score = -1
						}
						res += f[i][j+score+n][cur]
					}
				}
				f[i+1][j+n][pre] = res % mod
			}
		}
	}
	return f[n][n][0]
}

func countWinningSequences2(s string) int {
	const mod = 1_000_000_007
	mp := [...]int{'F': 0, 'W': 1, 'E': 2}
	n := len(s)
	f := [2][][3]int{}
	for i := range f {
		f[i] = make([][3]int, n*2+1)
	}
	for j := n + 1; j <= n*2; j++ {
		f[0][j] = [3]int{1, 1, 1}
	}
	pow2 := 1
	for i, c := range s {
		for j := -i; j < min(i+2, n-i); j++ {
			for pre := 0; pre < 3; pre++ {
				res := 0
				for cur := 0; cur < 3; cur++ {
					if i == n-1 || cur != pre {
						score := (cur - mp[c] + 3) % 3
						if score == 2 {
							score = -1
						}
						res += f[i%2][j+score+n][cur]
					}
				}
				f[(i+1)%2][j+n][pre] = res % mod
			}
		}
		pow2 = pow2 * 2 % mod
		for j := i + 2; j < n-i; j++ {
			f[(i+1)%2][j+n] = [3]int{pow2, pow2, pow2}
		}
	}
	return f[n%2][n][0]
}

func countWinningSequencesDFS(s string) int {
	const mod = 1_000_000_007
	n := len(s)
	pow2 := make([]int, (n+1)/2)
	pow2[0] = 1
	for i := 1; i < len(pow2); i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}

	mp := [...]int{'F': 0, 'W': 1, 'E': 2}
	memo := make([][][3]int, n)
	for i := range memo {
		memo[i] = make([][3]int, n*2+1)
		for j := range memo[i] {
			memo[i][j] = [3]int{-1, -1, -1}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, ban int) (res int) {
		if -j > i { // 必败
			return
		}
		if j > i+1 { // 必胜
			return pow2[i+1]
		}
		p := &memo[i][j+n][ban]
		if *p != -1 {
			return *p
		}
		for k := 0; k < 3; k++ { // 枚举当前召唤的生物
			// 判断 i==n-1 避免讨论 k == ban 的情况
			if i == n-1 || k != ban { // 不能连续两个回合召唤相同的生物
				score := (k - mp[s[i]] + 3) % 3
				if score == 2 {
					score = -1
				}
				res += dfs(i-1, j+score, k)
			}
		}
		res %= mod
		*p = res
		return
	}
	return dfs(n-1, 0, 0) // ban=0,1,2 都可以
}
