package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
const mod = 998244353

func run(in io.Reader, out io.Writer) {
	const mx = 500
	F := [mx]int{1}
	for i := 1; i < mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF := [mx]int{mx - 1: pow(F[mx-1], mod-2)}
	for i := mx - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	C := func(n, k int) int { return F[n] * invF[k] % mod * invF[n-k] % mod }

	var n, m, l, r, x int
	Fscan(in, &n, &m)
	type pair struct{ l, x int }
	ban := make([][]pair, n+1)
	for ; m > 0; m-- {
		Fscan(in, &l, &r, &x)
		ban[r] = append(ban[r], pair{l, x})
	}

	f := make([][]int, n+2)
	for i := range f {
		f[i] = make([]int, n+2)
	}
	f[n+1][n] = 1
	for l := n; l > 0; l-- {
		b := make([]bool, n+1)
		f[l][l-1] = 1
		for r := l; r <= n; r++ {
			for _, p := range ban[r] {
				if p.l >= l {
					b[p.x] = true
				}
			}
			for x := l; x <= r; x++ {
				if !b[x] {
					f[l][r] = (f[l][r] + C(r-l, x-l)*f[l][x-1]%mod*f[x+1][r]) % mod
				}
			}
		}
	}
	Fprint(out, f[1][n])
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

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
