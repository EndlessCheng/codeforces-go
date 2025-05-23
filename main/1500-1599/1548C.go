package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1548C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 1_000_000_007
	const inv3 = (mod + 1) / 3
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

	const mx int = 3e6 + 3
	var F, invF [mx + 1]int
	F[0] = 1
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF[mx] = pow(F[mx], mod-2)
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	C := func(n, k int) int { return F[n] * invF[k] % mod * invF[n-k] % mod }

	var n, q, x int
	Fscan(in, &n, &q)
	f := make([]int, n*3+2)
	f[1] = n + 1
	for i := 2; i < len(f); i++ {
		f[i] = (C(n*3+3, i) - f[i-1]*3 - f[i-2] + mod*4) * inv3 % mod
	}
	for range q {
		Fscan(in, &x)
		Fprintln(out, f[x+1])
	}
}

//func main() { cf1548C(bufio.NewReader(os.Stdin), os.Stdout) }
