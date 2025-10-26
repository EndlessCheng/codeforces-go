package main

import "slices"

// https://space.bilibili.com/206214
func countCoprime1(mat [][]int) int {
	const mod = 1_000_000_007
	mx := 0
	for _, row := range mat {
		mx = max(mx, slices.Max(row))
	}

	m := len(mat)
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, mx+1)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, g int) (res int) {
		if i < 0 {
			if g == 1 {
				return 1
			}
			return
		}
		p := &memo[i][g]
		if *p != -1 { // 之前计算过
			return *p
		}
		for _, x := range mat[i] {
			res += dfs(i-1, gcd(g, x))
		}
		res %= mod
		*p = res // 记忆化
		return
	}
	return dfs(m-1, 0)
}

func countCoprime2(mat [][]int) int {
	const mod = 1_000_000_007
	mx := 0
	for _, row := range mat {
		mx = max(mx, slices.Max(row))
	}

	m := len(mat)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, mx+1)
	}
	f[0][1] = 1
	for i, row := range mat {
		for g := 0; g <= mx; g++ {
			res := 0
			for _, x := range row {
				res += f[i][gcd(g, x)]
			}
			f[i+1][g] = res % mod
		}
	}
	return f[m][0]
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

const maxVal = 151

var divisors [maxVal][]int

func init() {
	for i := 1; i < maxVal; i++ {
		for j := i; j < maxVal; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}
}

func countCoprime(mat [][]int) int {
	const mod = 1_000_000_007
	// 预处理每行的因子个数
	divisorCnt := make([][]int, len(mat))
	mx := 0
	for i, row := range mat {
		rowMax := slices.Max(row)
		mx = max(mx, rowMax)
		divisorCnt[i] = make([]int, rowMax+1)
		for _, x := range row {
			for _, d := range divisors[x] {
				divisorCnt[i][d]++
			}
		}
	}

	cntGcd := make([]int, mx+1)
	for i := mx; i > 0; i-- {
		// 每行选一个 i 的倍数的方案数
		res := 1
		for _, cnt := range divisorCnt {
			if i >= len(cnt) || cnt[i] == 0 {
				res = 0
				break
			}
			res = res * cnt[i] % mod // 乘法原理
		}

		for j := i; j <= mx; j += i {
			res -= cntGcd[j] // 注意这里有减法，可能导致 res 是负数
		}

		cntGcd[i] = res % mod
	}
	return (cntGcd[1] + mod) % mod // 保证结果非负
}
