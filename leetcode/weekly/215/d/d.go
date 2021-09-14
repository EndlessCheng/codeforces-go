package main

// github.com/EndlessCheng/codeforces-go
const maxMask = 243 // 3**5
var cache [maxMask][maxMask]int

func init() {
	for i := 0; i < maxMask; i++ {
		for j := 0; j < maxMask; j++ {
			cache[i][j] = -1
		}
	}
}

// 教训：一开始没有写成数组的形式，导致下面 calc 函数老长一段逻辑 debug 半天
var self = [3]int{0, 120, 40}
var nb = [3]int{0, -30, 20}

// 技巧：能 cache 尽量 cache
func calc(cur, next int) (s int) {
	cc := &cache[cur][next]
	if *cc != -1 {
		return *cc
	}
	defer func() { *cc = s }()
	prev := false
	for cur > 0 {
		if tp := cur % 3; tp > 0 {
			s += self[tp]
			if prev {
				s += nb[tp]
			}
			if cur/3%3 > 0 {
				s += nb[tp]
			}
			if next%3 > 0 {
				s += nb[tp] + nb[next%3]
			}
			prev = true
		} else {
			prev = false
		}
		cur /= 3
		next /= 3
	}
	return
}

func getMaxGridHappiness(n, m, c1, c2 int) (ans int) {
	// 由于记录的是所有样例的运行时间，ifswap 会快不少
	if m > n {
		m, n = n, m
	}
	mx := 1
	for i := 0; i < m; i++ {
		mx *= 3
	}
	dp := make([][][][]int, n-1)
	for i := range dp {
		dp[i] = make([][][]int, mx)
		for j := range dp[i] {
			dp[i][j] = make([][]int, c1+1)
			for k := range dp[i][j] {
				dp[i][j][k] = make([]int, c2+1)
				for l := range dp[i][j][k] {
					dp[i][j][k][l] = -1
				}
			}
		}
	}

	var f func(p, cur, left1, left2 int) int
	f = func(p, cur, left1, left2 int) (res int) {
		if p >= 0 {
			if p == n-1 {
				return calc(cur, 0)
			}
			dv := &dp[p][cur][left1][left2]
			if *dv != -1 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		for next := 0; next < mx; next++ {
			c := [3]int{}
			for j := next; j > 0; j /= 3 {
				c[j%3]++
			}
			if c[1] <= left1 && c[2] <= left2 {
				if r := f(p+1, next, left1-c[1], left2-c[2]) + calc(cur, next); r > res {
					res = r
				}
			}
		}
		return
	}
	return f(-1, 0, c1, c2)
}
