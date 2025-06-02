package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1670F(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, l, r, z int
	Fscan(in, &n, &l, &r, &z)
	cn := make([]int, n+1)
	inv := make([]int, n+1)
	cn[0] = 1
	cn[1] = n
	inv[1] = 1
	for i := 2; i <= n; i++ {
		inv[i] = (mod - mod/i) * inv[mod%i] % mod
		cn[i] = cn[i-1] * (n + 1 - i) % mod * inv[i] % mod
	}

	calc := func(r int) int {
		if r < z {
			return 0
		}
		m := bits.Len(uint(r))
		dp := make([][][2]int, m)
		for i := range dp {
			dp[i] = make([][2]int, n)
			for j := range dp[i] {
				dp[i][j] = [2]int{-1, -1}
			}
		}
		var f func(int, int, int) int
		f = func(i, s, lessEq int) (res int) {
			if s>>(m-i) > 0 {
				return 0
			}
			if i == m {
				return lessEq
			}
			p := &dp[i][s][lessEq]
			if *p >= 0 {
				return *p
			}
			le := [2]int{1, lessEq}
			if r>>i&1 == 0 {
				le = [2]int{lessEq, 0}
			}
			for j := z >> i & 1; j <= n; j += 2 {
				res = (res + f(i+1, (s+j)>>1, le[(s+j)&1])*cn[j]) % mod
			}
			*p = res
			return
		}
		return f(0, 0, 1)
	}
	Fprint(out, (calc(r)-calc(l-1)+mod)%mod)
}

//func main() { cf1670F(os.Stdin, os.Stdout) }
