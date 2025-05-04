package main

import (
	"math/bits"
)

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const mx = 31

var fac [mx]int  // fac[i] = i!
var invF [mx]int // invF[i] = i!^-1

func init() {
	fac[0] = 1
	for i := 1; i < mx; i++ {
		fac[i] = fac[i-1] * i % mod
	}
	invF[mx-1] = pow(fac[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func magicalSum(m, k int, nums []int) int {
	n := len(nums)
	powV := make([][]int, n)
	for i, v := range nums {
		powV[i] = make([]int, m+1)
		powV[i][0] = 1
		for j := 1; j <= m; j++ {
			powV[i][j] = powV[i][j-1] * v % mod
		}
	}

	memo := make([][][][]int, n)
	for i := range memo {
		memo[i] = make([][][]int, m+1)
		for j := range memo[i] {
			memo[i][j] = make([][]int, m/2+1)
			for p := range memo[i][j] {
				memo[i][j][p] = make([]int, k+1)
				for q := range memo[i][j][p] {
					memo[i][j][p][q] = -1
				}
			}
		}
	}
	var dfs func(int, int, int, int) int
	dfs = func(i, leftM, x, leftK int) (res int) {
		c1 := bits.OnesCount(uint(x))
		if c1+leftM < leftK { // 可行性剪枝
			return
		}
		if i == n {
			if leftM == 0 && c1 == leftK {
				return 1
			}
			return
		}
		p := &memo[i][leftM][x][leftK]
		if *p != -1 {
			return *p
		}
		for j := range leftM + 1 { // 枚举 I 中有 j 个下标 i
			// 这 j 个下标 i 对 S 的贡献是 j * pow(2, i)
			// 由于 x = S >> i，转化成对 x 的贡献是 j
			bit := (x + j) & 1 // 取最低位，提前从 leftK 中减去，其余进位到 x 中
			if bit <= leftK {
				r := dfs(i+1, leftM-j, (x+j)>>1, leftK-bit)
				res = (res + r*powV[i][j]%mod*invF[j]) % mod
			}
		}
		*p = res
		return
	}
	return dfs(0, m, 0, k) * fac[m] % mod
}

func magicalSum2(m, k int, nums []int) int {
	n := len(nums)
	powV := make([][]int, n)
	for i, v := range nums {
		powV[i] = make([]int, m+1)
		powV[i][0] = 1
		for j := 1; j <= m; j++ {
			powV[i][j] = powV[i][j-1] * v % mod
		}
	}

	f := make([][][][]int, n+1)
	for i := range f {
		f[i] = make([][][]int, m+1)
		for j := range f[i] {
			f[i][j] = make([][]int, m/2+1)
			for x := range f[i][j] {
				f[i][j][x] = make([]int, k+1)
			}
		}
	}
	for x := range m/2 + 1 {
		c1 := bits.OnesCount(uint(x))
		if c1 <= k {
			f[n][0][x][c1] = 1
		}
	}

	for i := n - 1; i >= 0; i-- {
		for leftM := range m + 1 {
			for x := range m/2 + 1 {
				for leftK := range k + 1 {
					res := 0
					for j := range min(leftM, m-x) + 1 {
						bit := (x + j) & 1
						if bit <= leftK {
							r := f[i+1][leftM-j][(x+j)>>1][leftK-bit]
							res = (res + r*powV[i][j]%mod*invF[j]) % mod
						}
					}
					f[i][leftM][x][leftK] = res
				}
			}
		}
	}
	return f[0][m][0][k] * fac[m] % mod
}
