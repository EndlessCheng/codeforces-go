package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func p7961(in io.Reader, out io.Writer) {
	const mod = 998244353
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	const mx = 31
	var fac, invF [mx]int
	fac[0] = 1
	for i := 1; i < mx; i++ {
		fac[i] = fac[i-1] * i % mod
	}
	invF[mx-1] = pow(fac[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}

	var n, m, k, v int
	Fscan(in, &m, &n, &k)
	n++
	powV := make([][]int, n)
	for i := range powV {
		Fscan(in, &v)
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
		if i == n {
			if leftM == 0 && bits.OnesCount(uint(x)) <= leftK {
				return 1
			}
			return
		}
		p := &memo[i][leftM][x][leftK]
		if *p != -1 {
			return *p
		}
		for j := 0; j <= leftM; j++ {
			bit := (x + j) & 1
			if bit <= leftK {
				r := dfs(i+1, leftM-j, (x+j)>>1, leftK-bit)
				res = (res + r*powV[i][j]%mod*invF[j]) % mod
			}
		}
		*p = res
		return
	}
	Fprint(out, dfs(0, m, 0, k)*fac[m]%mod)
}

//func main() { p7961(bufio.NewReader(os.Stdin), os.Stdout) }
